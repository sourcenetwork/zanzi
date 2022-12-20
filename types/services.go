package types

import (
	"google.golang.org/protobuf/proto"

	o "github.com/sourcenetwork/source-zanzibar/pkg/option"
)

// Authorizer models Zanzibar's authorization verification calls
// over the relation graph.
type Authorizer interface {

	// Checks whether actor has relation with obj, return true if it does.
	// This is equivalent to asking whether the actor is reachable
	// from the (obj, relation) node.
	Check(policyId string, obj Entity, relation string, actor Entity) (bool, error)

	// Reverse traverses up through the relation graph and return
	// all object which are somehow related to the given actor.
	Reverse(policyId string, actor Entity) ([]EntityRelPair, error) // TODO would be wise to have a threshold on the number of objects

	// Expand returns a Tree representing all relations of obj with other entities
	Expand(policyId string, obj Entity, relation string) (ExpandTree, error)

	// TODO additional useful methods like verbose reverse
}

type ProtoConstraint[T any] interface {
	proto.Message
	*T
}

// RelationshipService exposes operations to manipulate system Relationships
type RelationshipService interface {
	Set(rel Relationship) error
	Delete(rel Relationship) error
	Has(rel Relationship) (bool, error)
	Get(rel Relationship) (o.Option[Relationship], error)
}

// RecordService exposes operations to manipulate records
type RecordService[T any, PT ProtoConstraint[T]] interface {
	Set(rel Relationship, data T) error
	Delete(rel Relationship) error
	Get(rel Relationship) (o.Option[Record[PT]], error)
	Has(rel Relationship) (bool, error)
}

// PolicyService exposes operations to manipulate system policies
type PolicyService interface {
	Set(policy Policy) error
	Get(id string) (o.Option[Policy], error)
	Delete(id string) error
}
