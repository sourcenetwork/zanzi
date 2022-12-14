package source_zanzibar

import (
    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    "github.com/sourcenetwork/source-zanzibar/types"
)

type policyMapper struct {}

func (m *policyMapper) ToInternal(p types.Policy) policy.Policy {
    var resources []*policy.Resources
    resMapper := func(res types.Resource) {return m.toInternalResource(res)}
    actorMapper := func(a types.Actor) {return m.toInternalActor(a)}
    return policy.Policy{
        Id: p.Id,
        Name: p.Name,
        Resources: utils.MapSlice(p.Resources, resMapper),
        Actors: utils.MapSlice(p.Actors, actorMapper),
        Attributes: p.Attributes,
    }
}

func (m *policyMapper) FromInternal(p policy.Policy) types.Policy {

}


func (m *policyMapper) mapInternalRule(rule *policy.Rule) any {
}

func (m *policyMapper) relToRule(rel types.Relation) *policy.Rule {
    return &policy.Rule {
        Type: policy.RuleType_RELATION,
        Name: perm.Name,
        // TODO map constraint
    }
}

func (m *policyMapper) permToRule(perm types.Permission) *policy.Rule {
    return &policy.Rule {
        Type: policy.RuleType_PERMISSION,
        Name: perm.Name,
        RewriteExpr: perm.RelationExpression,
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
        rule := m.relToRule(rel)
        rules = append(rules, rule)
    }

    for _, perm := range res.Permissions {
        rule := m.permToRule(perm)
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

func (m *tupleMapper[T]) RecordToTuple(rec types.Record[T]) tuple.Tuple[T] {
    tuple := m.relToTuple(rec.Relationship)
    tuple.SetData(rec.Data)
    tuple.CreatedAt = rec.CreatedAt
    return tuple
}

func (m *tupleMapper[T]) relToTuple(rel types.Relationship) tuple.Tuple[T] {
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
    case RelationshipType_ATTRIBUTE:
        dst.Type = tuple.NodeType_OBJECT
    case RelationshipType_GRANT:
        dst.Type = tuple.NodeType_ACTOR
    case RelationshipType_DELEGATE:
        dst.Type = tuple.NodeType_RELATION_SOURCE
    }

    return tuple.Tuple{
        Partition: rel.PolicyId,
        Source: src,
        Dest: dst,
    }
}


func (m *tupleMapper[T]) toRelationship(t tuple.Tuple) types.Relationship {
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
        Object: Entity{
            Namespace: t.Source.Namespace,
            Id: t.Source.Id,
        },
        Relation: t.Source.Relation,
        Subject: Entity{
            Namespace: t.Dest.Namespace,
            Id: t.Dest.Id,
        },
        SubjectRelation: t.Dest.Relation
    }
}

func (m *tupleMapper[T]) ToRecord(t tuple.Tuple) types.Record[T] {
    relationship := m.toRelationship(t)
    return Record[T] {
        CreatedAt: t.CreatedAt,
        Relationship: relationship,
        Data: t.GetData(),
    }
}
