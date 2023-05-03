package policy_definition

import (
    "sort"

    "golang.org/x/exp/constraints"

        "github.com/sourcenetwork/source-zanzibar/types"
        "github.com/sourcenetwork/source-zanzibar/pkg/utils"
)

// PolicyFromYaml attempts to unmarshal a PolicyDefinition from definition,
// assumes definition is a YAML document.
// If the PolicyDefinition is sucessfuly Unmarshaled, maps it to Zanzi's Policy type.
func PolicyFromYaml(definition string) (*types.Policy, error) {
    def, err := UnmarshalPolicyDefinition(definition)
    if err != nil {
        return nil, err
    }

    return MapPolicyDefinition(def), nil
}

// maps a PolicyDefinition to zanzi's Policy type.
// Does not perform any validation or input verification over the given PolicyDefinition.
func MapPolicyDefinition(def *PolicyDefinition) *types.Policy {
    resourceDefs := mapItemsToSlice[string, *ResourceDefinition, string](
        def.Resources,
        func(r *ResourceDefinition) string {
            return r.Name
        })

    actorDefs := mapItemsToSlice[string, *ActorDefinition, string](
        def.Actors, 
        func(a *ActorDefinition) string {
            return a.Name
        })

    return &types.Policy {
        Name: def.Name,
        Description: def.Doc,
        Resources: utils.MapSlice(resourceDefs, mapResource),
        Actors: utils.MapSlice(actorDefs, mapActor),
        Attributes: def.Attributes,
    }
}

// map ResourceDefinition to zanzi's Resource
func mapResource(resourceDef *ResourceDefinition) *types.Resource{
    permDefs := mapItemsToSlice(
        resourceDef.Permissions,
        func(p *PermissionDefinition) string{
            return p.Name
        })

    relDefs := mapItemsToSlice(
        resourceDef.Relations,
        func (r *RelationDefinition) string{
            return r.Name
        })

    return &types.Resource {
        Name: resourceDef.Name,
        Description: resourceDef.Doc,
        Permissions: utils.MapSlice(permDefs, mapPermission),
        Relations: utils.MapSlice(relDefs, mapRelation),
    }
}


// map ActorDefinition to zanzi's Actor
func mapActor(actorDef *ActorDefinition) *types.Actor {
    return &types.Actor {
        Name: actorDef.Name,
        Description: actorDef.Doc,
    }
}

// map PermissionDefinition to zanzi's Permission
func mapPermission(permDef *PermissionDefinition) *types.Permission {
    return &types.Permission {
        Name: permDef.Name,
        Description: permDef.Doc,
        PermissionExpr: permDef.Expr,
    }
}

// map RelationDefinition to zanzi's Relation
func mapRelation(relDef *RelationDefinition) *types.Relation {
    return &types.Relation {
        Name: relDef.Name,
        Description: relDef.Doc,
    }
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
// the keyFunc could return T.Name to sort the Persons in the slice
// by their names.
type sortable[T any, K constraints.Ordered] struct {
    items []T
    keyFunc func(T) K
}

func asSortable[T any, K constraints.Ordered](slice []T, keyFunc func(T) K) sortable[T, K] {
    return sortable[T, K] {
        items: slice,
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
