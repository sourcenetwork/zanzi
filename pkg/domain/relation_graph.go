package domain

import (
	"fmt"

	"crypto/sha256"
	"encoding/base64"
)

func (n *RelationNode) GetObject() *Entity {
	switch node := n.Node.(type) {
	case *RelationNode_EntitySet:
		return node.EntitySet.Object
	case *RelationNode_Entity:
		return node.Entity.Object
	}
	return nil
}

func (n *RelationNode) IsEntity() bool {
	entity := n.GetEntity()
	return entity != nil
}

// IsTerminalNode returns true if node is a leaf / terminal node in the graph.
// ie. Wilcard or Entity node
func (n *RelationNode) IsTerminalNode() bool {
	switch n.Node.(type) {
	case *RelationNode_EntitySet:
		return false
	}
	return true
}

// Hash produces an Id of a RelationNode which uniquely identifies it.
func (n *RelationNode) Id() string {
	hasher := sha256.New()
	separator := []byte{byte('/')}

	// since Ids and Relations are required, it's impossible
	// for users to created an entity node which would cause a clash
	switch node := n.Node.(type) {
	case *RelationNode_EntitySet:
		hasher.Write([]byte(node.EntitySet.Object.Resource))
		hasher.Write(separator)
		hasher.Write([]byte(node.EntitySet.Object.Id))
		hasher.Write(separator)
		hasher.Write([]byte(node.EntitySet.Relation))
	case *RelationNode_Entity:
		hasher.Write([]byte(node.Entity.Object.Resource))
		hasher.Write(separator)
		hasher.Write([]byte(node.Entity.Object.Id))
		hasher.Write(separator)
	case *RelationNode_Wildcard:
		hasher.Write([]byte(node.Wildcard.Resource))
		hasher.Write(separator)
		hasher.Write(separator)
	default:
	}

	hash := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash)
}

func (n *RelationNode) PrettyString() string {
	var resource, id, relation string
	switch node := n.Node.(type) {
	case *RelationNode_EntitySet:
		resource = node.EntitySet.Object.Resource
		id = node.EntitySet.Object.Id
		relation = node.EntitySet.Relation
	case *RelationNode_Entity:
		resource = node.Entity.Object.Resource
		id = node.Entity.Object.Id
	case *RelationNode_Wildcard:
		resource = node.Wildcard.Resource
		id = "*"
	default:
	}
	str := fmt.Sprintf("(%v:%v", resource, id)
	if relation != "" {
		str += ", " + relation
	}
	str += ")"
	return str
}
