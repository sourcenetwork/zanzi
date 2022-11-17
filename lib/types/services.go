package types

import (
    "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/lib/types/tree"
    "github.com/sourcenetwork/source-zanzibar/lib/types"
)


// Authorizer models Zanzibar's authorization verification calls
// over the relation graph.
type Authorizer interface {

    // Checks whether actor has relation with obj, return true if it does.
    // This is equivalent to asking whether the actor is reachable
    // from the (obj, relation) node.
    Check(obj types.Id, relation string, actor types.Id) (bool, error)


    // Reverse traverses up through the relation graph and return
    // all object which are somehow related to the given actor.
    Reverse(actor types.Id) ([](Id, string), error) // TODO would be wise to have a threshold on the number of objects

    // Expand returns a Tree representing all relations of obj with other entities
    Expand(obj types.Id, relation string) (tree.UsersetNode, error)

    // TODO additional useful methods like verbose reverse
}

// RelationService exposes operations to manipulate system Relations
type RelationService[T proto.Message] interface {
    Set(rel EntityRelation, data T) error
    Delete(rel Relation) error

    Grant(rel GrantRelation, data T) error
    Ungrant(rel GrantRelation) error

    Delegate(rel DelegationRelation, data T) error
    Undelegate(rel DelegationRelation) error

    GetRelations(obj ObjectId, relationName string) ([]Relation, error)
}

// PolicyService exposes operations to manipulate system policies
type PolicyService interface {
    Set(policy Policy) error
    Get(policyId string) (types.Option[types.Policy], error)
    Delete(policy Policy) error
}
