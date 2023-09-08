package relation_graph

import (
    "fmt"
    "context"

    "github.com/sourcenetwork/zanzi/pkg/domain"
    "github.com/sourcenetwork/zanzi/internal/utils"
    "github.com/sourcenetwork/zanzi/pkg/types"
)

func newEvaluator(repository NodeRepository, logger types.Logger) evaluator{
    return evaluator{
        repository: repository,
        logger: logger,
    }
}

// evaluator evaluates rewrite rules
type evaluator struct {
    repository NodeRepository
    logger types.Logger
}

func (e *evaluator) getRepository() NodeRepository {
    return e.repository
}

func (e *evaluator) Evaluate(ctx context.Context, policyId string, rule *domain.Rule, node *domain.EntitySetNode) ([]*domain.RelationNode, error) {
    var nodes []*domain.RelationNode
    var err error

    switch ruleType := rule.Rule.(type) {
    case *domain.Rule_This:
        nodes, err = e.evaluateThis(ctx, policyId, node)
    case *domain.Rule_Ttu:
        nodes, err = e.evaluateTupleToUserset(ctx, policyId, ruleType.Ttu, node)
    case *domain.Rule_Cu:
        nodes = e.evaluateComputedUserset(ruleType.Cu, node)
    default:
        err = fmt.Errorf("rule %v: %w", rule, domain.ErrInvalidVariant)
    }

    if err != nil {
        return nil, fmt.Errorf("rule evaluation: node %v rule %v: %w", node, rule, err)
    }

    e.logger.Debugf("evalutator: rule %v, nodes %v", rule, nodes)
    return nodes, nil
}

// evaluateThis evaluates a This rule by returning the direct sucessors of the given node
func (f *evaluator) evaluateThis(ctx context.Context, policyId string, node *domain.EntitySetNode) ([]*domain.RelationNode, error) {
    repository := f.getRepository()

    sucessors, err := repository.GetSucessors(ctx, policyId, node)
    if err != nil {
        return nil, fmt.Errorf("this rule: %w", err)
    }

    return sucessors, nil
}

// evaluateComputedUserset evaluates a ComputedUserset rule by constructing a RelationNode
// and returning it
func (f *evaluator) evaluateComputedUserset(rule *domain.ComputedUserset, node *domain.EntitySetNode) []*domain.RelationNode {
    return []*domain.RelationNode{
        &domain.RelationNode{
            Node: &domain.RelationNode_EntitySet{
                EntitySet: &domain.EntitySetNode{
                    Object: node.Object,
                    Relation: rule.TargetRelation,
                },
            },
        },
    }
}

// evaluateTupleToUserset evaluates a Tuple to Userset (TTU) rule
//
// The steps to evaluate the TTU rule are:
// 1. Use node to build a new node with its relation replaced - call it filter
// 2. Fetch filter's sucessors from tuple store
// 3. Map fetched sucessors into a new node by replacing its relation with the cu_relation given in
// the TupleToUserset rule
// The mapped sucessors are the resulting nodes from evaluating TTU
func (f *evaluator) evaluateTupleToUserset(ctx context.Context, policyId string, rule *domain.TupleToUserset, node *domain.EntitySetNode) ([]*domain.RelationNode, error) {
    repository := f.getRepository()

    filter :=  &domain.EntitySetNode{
        Object: node.Object,
        Relation: rule.TuplesetRelation,
    }
    sucessors, err := repository.GetSucessors(ctx, policyId, filter)
    if err != nil {
        return nil, fmt.Errorf("tuple to userset rule: %w", err)
    }
    f.logger.Debugf("ttu evaluation found tuplesets: %v", sucessors)

    mapper := func (node *domain.RelationNode) (*domain.RelationNode, error) {
        return buildTupleToUsersetNode(node, rule.ComputedUsersetRelation)
    }
    nodes, err := utils.MapSliceErr(sucessors, mapper)

    return nodes, err
}

func buildTupleToUsersetNode(node *domain.RelationNode, computedRelation string) (*domain.RelationNode, error) {
    entitySetNode := &domain.EntitySetNode{
        Object: nil,
        Relation: computedRelation,
    }

    switch nodeType := node.Node.(type) {
    case *domain.RelationNode_EntitySet:
        entitySetNode.Object = nodeType.EntitySet.Object
    case *domain.RelationNode_Entity:
        entitySetNode.Object = nodeType.Entity.Object
    case *domain.RelationNode_Wildcard:
        //TODO not sure how I should go about this scenario yet...
    default:
        return nil, fmt.Errorf("building tuple to userset node: node %v: %w", node, domain.ErrInvalidVariant)
    }

    return &domain.RelationNode{
        Node: &domain.RelationNode_EntitySet{
            EntitySet: entitySetNode,
        },
    }, nil
}
