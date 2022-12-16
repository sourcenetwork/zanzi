package source_zanzibar

import (
    "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    rg "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph"
    "github.com/sourcenetwork/source-zanzibar/types"
    "github.com/sourcenetwork/source-zanzibar/pkg/utils"
)

type policyMapper struct {}

func (m *policyMapper) ToInternal(p types.Policy) policy.Policy {
    resMapper := func(res types.Resource) *policy.Resource {return m.toInternalResource(res)}
    actorMapper := func(a types.Actor) *policy.Actor {return m.toInternalActor(a)}
    return policy.Policy{
        Id: p.Id,
        Name: p.Name,
        Resources: utils.MapSlice(p.Resources, resMapper),
        Actors: utils.MapSlice(p.Actors, actorMapper),
        Attributes: p.Attributes,
    }
}

func (m *policyMapper) FromInternal(p *policy.Policy) types.Policy {
    return types.Policy {
        // TODO
    }
}


func (m *policyMapper) mapInternalRule(rule *policy.Rule) any {
    // TODO
    return nil
}

func (m *policyMapper) relationToRule(rel types.Relation) *policy.Rule {
    return &policy.Rule {
        Type: policy.RuleType_RELATION,
        Name: rel.Name,
        // TODO map constraint
    }
}

func (m *policyMapper) permissionToRule(perm types.Permission) *policy.Rule {
    return &policy.Rule {
        Type: policy.RuleType_PERMISSION,
        Name: perm.Name,
        RewriteExpr: perm.Expression,
        // TODO call parser in order to build tree
    }
}

func (m *policyMapper) toInternalActor(actor types.Actor) *policy.Actor {
    return &policy.Actor {
        Name: actor.Name,
        // TODO constraints
    }
}

func (m *policyMapper) fromInternalActor(actor *policy.Actor) types.Actor {
    return types.Actor {
        Name: actor.Name,
        // TODO map kinds
    }
}

func (m *policyMapper) addRules() {

}

func (m *policyMapper) toInternalResource(res types.Resource) *policy.Resource {
    var rules []*policy.Rule

    for _, rel := range res.Relations {
        rule := m.relationToRule(rel)
        rules = append(rules, rule)
    }

    for _, perm := range res.Permissions {
        rule := m.permissionToRule(perm)
        rules = append(rules, rule)
    }

    return &policy.Resource {
        Name: res.Name,
        Rules: rules,
    }
}

func (m *policyMapper) fromInternalResource(res *policy.Resource) types.Resource {
    return types.Resource {

    }
}


type tupleMapper[T proto.Message] struct { }

func (m *tupleMapper[T]) FromRecord(rec types.Record[T]) tuple.Tuple[T] {
    tuple := m.FromRelationship(rec.Relationship)
    tuple.SetData(rec.Data)
    tuple.CreatedAt = rec.CreatedAt
    return tuple
}

func (m *tupleMapper[T]) FromRelationship(rel types.Relationship) tuple.Tuple[T] {
    src := tuple.TupleNode {
        Namespace: rel.Object.Namespace,
        Id: rel.Object.Id,
        Relation: rel.Relation,
        Type: tuple.NodeType_RELATION_SOURCE,
    }

    dst := tuple.TupleNode {
        Namespace: rel.Subject.Namespace,
        Id: rel.Subject.Id,
        Relation: rel.SubjectRelation,
    }

    switch rel.Type {
    case types.RelationshipType_ATTRIBUTE:
        dst.Type = tuple.NodeType_OBJECT
    case types.RelationshipType_GRANT:
        dst.Type = tuple.NodeType_ACTOR
    case types.RelationshipType_DELEGATE:
        dst.Type = tuple.NodeType_RELATION_SOURCE
    }

    return tuple.Tuple[T]{
        Partition: rel.PolicyId,
        Source: src,
        Dest: dst,
    }
}


func (m *tupleMapper[T]) toRelationship(t tuple.Tuple[T]) types.Relationship {
    var relType types.RelationshipType
    switch t.Dest.Type {
    case tuple.NodeType_ACTOR:
        relType = types.RelationshipType_GRANT
    case tuple.NodeType_OBJECT:
        relType = types.RelationshipType_ATTRIBUTE
    case tuple.NodeType_RELATION_SOURCE:
        relType = types.RelationshipType_DELEGATE
    }

    return types.Relationship{
        PolicyId: t.Partition,
        Type: relType,
        Object: types.Entity{
            Namespace: t.Source.Namespace,
            Id: t.Source.Id,
        },
        Relation: t.Source.Relation,
        Subject: types.Entity{
            Namespace: t.Dest.Namespace,
            Id: t.Dest.Id,
        },
        SubjectRelation: t.Dest.Relation,
    }
}

func (m *tupleMapper[T]) ToRecord(t tuple.Tuple[T]) types.Record[T] {
    relationship := m.toRelationship(t)
    return types.Record[T] {
        CreatedAt: t.CreatedAt,
        Relationship: relationship,
        Data: t.GetData(),
    }
}

type treeMapper struct { }

func (m *treeMapper) ToExpandTree(tree *rg.RelationNode) types.ExpandTree {
    return types.ExpandTree {
        Entity: types.Entity {
            Namespace: tree.ObjRel.Namespace,
            Id: tree.ObjRel.Id,
        },
        RelOrPerm: tree.ObjRel.Relation,
        RelationExpression: m.toExpressionTree(tree.Child),
    }
}

func (m *treeMapper) mapOp(op policy.Operation) types.Operator {
    var operator types.Operator
    switch op {
    case policy.Operation_UNION:
        operator = types.UNION
    case policy.Operation_DIFFERENCE:
        operator = types.DIFFERENCE
    case policy.Operation_INTERSECTION:
        operator = types.INTERSECTION
    }
    return operator
}

func (m *treeMapper) toFactor(node *rg.RuleNode) *types.Factor {
    f := func(relNode *rg.RelationNode) types.ExpandTree {return m.ToExpandTree(relNode)}
    return &types.Factor{
        Operand: m.ruleToOperand(&node.RuleData),
        Result: utils.MapSlice(node.Children, f),
    }
}

func (m *treeMapper) toExpression(node *rg.OpNode) *types.Expression {
    return &types.Expression {
        Operator: m.mapOp(node.JoinOp),
        Left: m.toExpressionTree(node.Left),
        Right: m.toExpressionTree(node.Right),
    }
}

func (m *treeMapper) toExpressionTree(node rg.RewriteNode) types.ExpressionTree {
    switch n := node.(type) {
    case *rg.OpNode:
        return m.toExpression(n)
    case *rg.RuleNode:
        return m.toFactor(n)
    default:
        panic("unknown RewriteNode")
    }
}

func (m *treeMapper) ruleToOperand(rule *rg.RuleData) string {
    switch rule.Type {
    case rg.RuleType_THIS:
        return "this"
    case rg.RuleType_CU:
        operand := rule.Args["Relation"]
        return operand
    case rg.RuleType_TTU:
        operand := rule.Args["TuplesetRelation"] + "->" + rule.Args["ComputedUsersetRelation"]
        return operand
    default:
        panic("unknown rule type")
    }
}
