package policy

import (
	parser "github.com/sourcenetwork/zanzi/internal/relation_expression_parser"
	"github.com/sourcenetwork/zanzi/pkg/api"
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/errors"
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
			err = errors.Wrap("parsing expression tree", err)
		}
	case *domain.RelationExpression_Tree:
		tree = t.Tree
	default:
		err = errors.Wrap("expression tree", errors.ErrInvalidVariant,
			errors.Pair("relation", r.Name),
		)
	}

	return tree, err
}

func GetPolicyFromDefinition(d *api.PolicyDefinition) (*domain.Policy, error) {
	var policy *domain.Policy
	var err error

	switch definition := d.Definition.(type) {
	case *api.PolicyDefinition_Policy:
		policy = definition.Policy
	case *api.PolicyDefinition_PolicyYaml:
		policy, err = policy_definition.PolicyFromYaml(definition.PolicyYaml)
	default:
		err = errors.Wrap("policy definition", errors.ErrInvalidVariant)
	}

	if err != nil {
		return nil, err
	}
	return policy, err
}
