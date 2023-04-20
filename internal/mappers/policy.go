package mappers

import (
	"fmt"

	"github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
	parser "github.com/sourcenetwork/source-zanzibar/internal/permission_parser"
	"github.com/sourcenetwork/source-zanzibar/pkg/utils"
	"github.com/sourcenetwork/source-zanzibar/types"
)

type PolicyMapper struct{}

func (m *PolicyMapper) ToInternal(p types.Policy) (policy.Policy, error) {
	var resources []*policy.Resource = make([]*policy.Resource, 0, len(p.Resources))

	for _, resource := range p.Resources {
		mapped, err := toInternalResource(resource)
		if err != nil {
			return policy.Policy{}, fmt.Errorf("failed mapping policy %v: %w", p.Id, err)
		}
		resources = append(resources, mapped)
	}

	return policy.Policy{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Resources:   resources,
		Actors:      utils.MapSlice(p.Actors, toInternalActor),
		Attributes:  p.Attributes,
                //FIXME CreatedAt:   p.CreatedAt,
	}, nil
}

func relationToInternal(rel *types.Relation) *policy.Relation {
	return &policy.Relation{
		Name:           rel.Name,
                Description: rel.Description,
		RewriteExpr:    "_this",
		ExpressionTree: policy.ThisTree(),
	}
}



func permissionToRelation(perm *types.Permission) (*policy.Relation, error) {
	parseTree, err := parser.Parse(perm.PermissionExpr)
	if err != nil {
		return nil, fmt.Errorf("failed mapping permission %v: %w", perm.Name, err)
	}
	tree := ToRewriteTree(parseTree)

	return &policy.Relation{
		Name:           perm.Name,
                Description: perm.Description,
		RewriteExpr:    perm.PermissionExpr,
		ExpressionTree: tree,
	}, nil
}

func toInternalActor(actor *types.Actor) *policy.Actor {
	return &policy.Actor{
		Name:        actor.Name,
                Description: actor.Description,
	}
}

func fromInternalActor(actor *policy.Actor) *types.Actor {
	return &types.Actor{
		Name:       actor.Name,
                Description: actor.Description,
	}
}

func toInternalResource(res *types.Resource) (*policy.Resource, error) {
	var relations []*policy.Relation

	for _, rel := range res.Relations {
		relation := relationToInternal(rel)
		relations = append(relations, relation)
	}

	for _, perm := range res.Permissions {
		relation, err := permissionToRelation(perm)
		if err != nil {
			return nil, fmt.Errorf("failed mapping relation %v: %w", res.Name, err)
		}
		relations = append(relations, relation)
	}

	return &policy.Resource{
		Name:      res.Name,
		Relations: relations,
	}, nil
}

func (m *PolicyMapper) FromInternal(p *policy.Policy) types.Policy {
	return types.Policy{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		// FIXME CreatedAt:   p.CreatedAt,
		Resources:   utils.MapSlice(p.Resources, fromInternalResource),
		Actors:      utils.MapSlice(p.Actors, fromInternalActor),
		Attributes:  p.Attributes,
	}
}

func fromInternalResource(res *policy.Resource) *types.Resource {
	permissions, relations := fromInternalRelations(res.Relations)
	return &types.Resource{
		Name:        res.Name,
                Description: res.Description,
		Relations:   relations,
		Permissions: permissions,
	}
}

func fromInternalRelations(relations []*policy.Relation) ([]*types.Permission, []*types.Relation) {
	var rels []*types.Relation
	var permissions []*types.Permission
	for _, rel := range relations {
		if isThisTree(rel.ExpressionTree) {
			rel := &types.Relation{
				Name: rel.Name,
                                Description: rel.Description,
			}
			rels = append(rels, rel)
		} else {
			permission := &types.Permission{
				Name:       rel.Name,
                                Description: rel.Description,
				PermissionExpr: rel.RewriteExpr,
			}
			permissions = append(permissions, permission)
		}
	}
	return permissions, rels
}

func isThisTree(tree *policy.Tree) bool {
	switch node := tree.Node.(type) {
	case *policy.Tree_Leaf:
		switch node.Leaf.Rule.RewriteRule.(type) {
		case *policy.RewriteRule_This:
			return true
		default:
			return false
		}
	default:
		return false
	}
}
