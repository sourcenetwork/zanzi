package types

import (
    "time"
)

// Relationship is a container type for any relation.
// Embeds client application data.
type Relationship struct {
    createdAt time.Time
    lastModified time.Time
    PolicyId string
    Type RelationshipType
    Object Entity
    Relation string
    Subject Entity
    SubjectRelation string // SubjectRelation may be empty depending on the relation type
}

func (r *Relationship) setCreationDate(timestamp time.Time) {
    r.createdAt = timestamp
}

func (r *Relationship) setLastModified(timestamp time.Time) {
    r.lastModified = timestamp
}

func (r *Relationship) GetCreationDate() time.Time {
    return r.createdAt
}

func (r *Relationship) GetLastModified() time.Time {
    return r.lastModified
}

// Identifies a system entity
type Entity struct {
    Namespace string
    Id string
}

// EntityRelPair is a pair composed of an entity and a relation.
// Conceptually it represents a set of actors which are reachable
// by expanding the pair
type EntityRelPair struct {
    Entity Entity
    Relationship string
}

// RelationshipType enumerates the different possible relation types
type RelationshipType int 

const (
    // Represents a relation from an ACTORSET to an OBJECT node.
    // used to express some relation between the source object and target object.
    RelationshipType_ATTRIBUTE RelationshipType = iota

    // Represents a relation grant between an ACTORSET and an ACTOR node
    // Effectively sets a relation between source object and dest actor
    RelationshipType_GRANT

    // Represents a delegated relation between two ACTORSET nodes.
    // Delegation is used to build indirect relations between users
    RelationshipType_DELEGATE
)

// Record represents a Relationship Record stored in the 
type Record[T any] struct {
    Relationship Relationship
    Data T
}
