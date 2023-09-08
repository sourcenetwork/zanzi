package policy

import (
        "fmt"

        "github.com/sourcenetwork/zanzi/pkg/domain"
        "github.com/sourcenetwork/zanzi/pkg/api"
        parser "github.com/sourcenetwork/zanzi/internal/relation_expression_parser"
        "github.com/sourcenetwork/zanzi/pkg/policy_definition"
)


func allRelationshipsSelector() domain.RelationshipSelector {
    builder := domain.SelectorBuilder{}
    builder.AnyObject()
    builder.AnyRelation()
    builder.AnySubject()
    return builder.Build()
}

func GetExpressionTree(r *domain.Relation) (*domain.RelationExpressionTree, error) {
    var tree *domain.RelationExpressionTree
    var err error

    switch t := r.RelationExpression.Expression.(type) {
    case *domain.RelationExpression_Expr:
        tree, err = parser.Parse(t.Expr)
        if err != nil {
            err = fmt.Errorf("expression tree: relation %v: %w", r.Name, err)
        }
    case *domain.RelationExpression_Tree:
        tree = t.Tree
    default:
        err = fmt.Errorf("expression tree: relation %v: obj %v: %w", r.Name, t, domain.ErrInvalidVariant)
    }

    return tree, err
}

func GetPolicyFromDefinition(d *api.PolicyDefinition) (*domain.Policy, error) {
    var policy *domain.Policy
    var err error

    switch definition := d.Definition.(type){
    case *api.PolicyDefinition_Policy:
        policy = definition.Policy
    case *api.PolicyDefinition_PolicyYaml:
        policy, err = policy_definition.PolicyFromYaml(definition.PolicyYaml)
    default:
        err = fmt.Errorf("PolicyDefinition %v: %w", definition, domain.ErrInvalidVariant)
    }

    if err != nil {
        return nil, err
    }

    return policy, err
}
