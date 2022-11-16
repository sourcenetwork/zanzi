package relation_graph

type RelationGraph interface {
    Walk(source AuthNode) []RelationEdge

    // path
    GetPath(source AuthNode, dest AuthNode) []RelationEdge

    GetSucessors(source AuthNode) []RelationEdge

    GetAncestors(source AuthNode) []RelationEdge
}


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

// Relation graph edge definition
// Edges in the relation graph express a relation between two nodes.
type RelationEdge struct {
    Source AuthNode
    Dest AuthNode
}
