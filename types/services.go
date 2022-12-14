package types

import (
    "google.golang.org/protobuf/proto"
)


// Authorizer models Zanzibar's authorization verification calls
// over the relation graph.
type Authorizer interface {

    // Checks whether actor has relation with obj, return true if it does.
    // This is equivalent to asking whether the actor is reachable
    // from the (obj, relation) node.
    Check(obj Id, relation string, actor Id) (bool, error)


    // Reverse traverses up through the relation graph and return
    // all object which are somehow related to the given actor.
    Reverse(actor Id) ([]EntityRelPair, error) // TODO would be wise to have a threshold on the number of objects

    // Expand returns a Tree representing all relations of obj with other entities
    Expand(obj Id, relation string) (ExpandTree, error)

    // TODO additional useful methods like verbose reverse
}

// RelationshipService exposes operations to manipulate system Relationships
type RelationshipService[T proto.Message] interface {
    Set(rel Relationship, data T) error
    Delete(rel Relationship) error
    Get(rel Relationship) (o.Option[Record[T]], error)
    GetRelationships(entity Entity, relation string) ([]Relationship, error)
}

// PolicyService exposes operations to manipulate system policies
type PolicyService interface {
    Set(policy Policy) error
    Get(policyId string) (o.Option[Policy], error)
    Delete(policy Policy) error
}
