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


// UserType represents the variant in how the User type should be interpreted.
// 
// The data model allows for an user to be an intermidiary node or
// a Leaf which represents an object.
// The Leaf Nodes of the Graph should - under normal usage - represent 
// client application specific subjects (ie entities which needs to be authorized).
type NodeType uint8

const (
    // Represents an Application Actor
    NodeType_ACTOR NodeType = iota

    // Represents a system object
    NodeType_OBJECT

    // Represents a node with an object/relation pair,
    // which when followed lead to a set of actors
    NodeType_PROXY
)


type RelationType uint8

const (
    // Represents a relation from an ACTORSET to an OBJECT node.
    // used to express some relation between the source object and target object.
    RelationType_ATTRIBUTE RelationType = iota

    // Represents an authorization grant between an ACTORSET and an ACTOR node
    // Effectively sets a relation between source object and dest actor
    RelationType_GRANT

    // Represents a delegated relation between two ACTORSET nodes.
    // Delegation is used to build indirect relations between users
    RelationType_DELEGATION
)
