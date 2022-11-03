package model

// AuthNode defines a Node in the relation graph
// It's analogous to Zanzibar's Userset definition.
// Namespace should be globally unique in the system
// Relation must be defined within the given Namespace
type AuthNode struct {
    Type NodeType
    Namespace string
    ObjectId string
    Relation string
}

// Tuples represents an entry in the Relation graph.
// A tuple can be thought of as an Edge which both defines and relates usersets (nodes).
// TupleRecord is an abstract representation of a tuple.
// Applications may implement custom records with their own data
type TupleRecord interface {
    GetObject() AuthNode
    GetSubject() AuthNode
}
