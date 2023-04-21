package types

// EntityRelPair is a pair composed of an entity and a relation.
// Conceptually it represents a set of actors which are reachable
// by expanding the pair
type EntityRelPair struct {
	Entity       Entity
	Relationship string
}

// Record represents a Relationship Record stored in the
type Record[T any] struct {
	Relationship Relationship
	Data         T
}
