package policy_definition

import (
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"

	"github.com/sourcenetwork/zanzi/internal/utils"
	"github.com/sourcenetwork/zanzi/pkg/domain"
)

// PolicyFromYaml attempts to unmarshal a PolicyDefinition from definition,
// assumes definition is a YAML document.
// If the PolicyDefinition is sucessfuly Unmarshaled, maps it to Zanzi's Policy type.
func PolicyFromYaml(definition string) (*domain.Policy, error) {
	def, err := UnmarshalPolicyDefinition(definition)
	if err != nil {
		return nil, err
	}

	return MapPolicyDefinition(def), nil
}

// maps a PolicyDefinition to zanzi's Policy type.
// Does not perform any validation or input verification over the given PolicyDefinition.
func MapPolicyDefinition(def *PolicyDefinition) *domain.Policy {
	resourceDefs := mapItemsToSlice(
		def.Resources,
		func(r *ResourceDefinition) string {
			return r.name
		})

	return &domain.Policy{
		Id:          def.Id,
		Name:        def.Name,
		Description: def.Doc,
		Resources:   utils.MapSlice(resourceDefs, mapResource),
		Attributes:  def.Attributes,
	}
}

// map ResourceDefinition to zanzi's Resource
func mapResource(resourceDef *ResourceDefinition) *domain.Resource {
	relDefs := mapItemsToSlice(
		resourceDef.Relations,
		func(r *RelationDefinition) string {
			return r.name
		})

	return &domain.Resource{
		Name:        resourceDef.name,
		Description: resourceDef.Doc,
		Relations:   utils.MapSlice(relDefs, mapRelation),
	}
}

// map RelationDefinition to zanzi's Relation
func mapRelation(relDef *RelationDefinition) *domain.Relation {
	return &domain.Relation{
		Name:        relDef.name,
		Description: relDef.Doc,
		RelationExpression: &domain.RelationExpression{
			Expression: &domain.RelationExpression_Expr{
				Expr: relDef.Expr,
			},
		},
		SubjectRestriction: mapSubjectRestriction(relDef.Types),
	}
}

func mapSubjectRestriction(types []string) *domain.SubjectRestriction {
	if slices.Contains(types, UniversalSetType) {
		return &domain.SubjectRestriction{
			SubjectRestriction: &domain.SubjectRestriction_UniversalSet{
				UniversalSet: &domain.UniversalSet{},
			},
		}
	}

	return &domain.SubjectRestriction{
		SubjectRestriction: &domain.SubjectRestriction_RestrictionSet{
			RestrictionSet: &domain.SubjectRestrictionSet{
				Restrictions: utils.MapSlice(types, mapSubjectRestrictionElem),
			},
		},
	}
}

func mapSubjectRestrictionElem(elem string) *domain.SubjectRestrictionSet_Restriction {
	restriction := &domain.SubjectRestrictionSet_Restriction{}

	resource, relation, found := strings.Cut(elem, TypeRelationSeparator)

	if !found {
		restriction.Entry = &domain.SubjectRestrictionSet_Restriction_Entity{
			Entity: &domain.EntityRestriction{
				ResourceName: resource,
			},
		}
	} else {
		restriction.Entry = &domain.SubjectRestrictionSet_Restriction_EntitySet{
			EntitySet: &domain.EntitySetRestriction{
				ResourceName: resource,
				RelationName: relation,
			},
		}
	}

	return restriction
}

// mapItemsToSlice extract all items in a Map, compiles them into a Slice
// and sorts the Slice using the keyFunc
//
// keyFunc is used to extract a field from the items and use it as a sortable key.
func mapItemsToSlice[Key comparable, Val any, Comp constraints.Ordered](m map[Key]Val, keyFunc func(Val) Comp) []Val {
	items := make([]Val, 0, len(m))
	for _, item := range m {
		items = append(items, item)
	}

	sort.Sort(asSortable(items, keyFunc))

	return items
}

// sortable wraps a slice of Ts and uses a KeyFunction
// which extracts an Orderable field of T.
// Implements the "sort" package sorting interface
//
// eg. Suppose T is a Person and each Person has a name,
// the keyFunc could return T.name to sort the Persons in the slice
// by their names.
type sortable[T any, K constraints.Ordered] struct {
	items   []T
	keyFunc func(T) K
}

func asSortable[T any, K constraints.Ordered](slice []T, keyFunc func(T) K) sortable[T, K] {
	return sortable[T, K]{
		items:   slice,
		keyFunc: keyFunc,
	}
}

func (s sortable[T, K]) Len() int {
	return len(s.items)
}

func (s sortable[T, K]) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func (s sortable[T, K]) Less(i, j int) bool {
	return s.keyFunc(s.items[i]) < s.keyFunc(s.items[j])
}
