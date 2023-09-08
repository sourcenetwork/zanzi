package kv_store

import (
	"crypto/sha256"

	rcdb "github.com/sourcenetwork/raccoondb"

	"github.com/sourcenetwork/zanzi/pkg/domain"
)

const (
	sha256bytes int = 256 / 8
)

var _ rcdb.Edge[*RelationNode] = (*Relationship)(nil)
var _ rcdb.NodeKeyer[*RelationNode] = (*relationNodeKeyer)(nil)
var _ rcdb.Ider[*domain.PolicyRecord] = (*policyIDer)(nil)
var _ rcdb.Ider[*RelationshipData] = (*relationshipDataIDer)(nil)


// relationNodeKeyer implements the NodeKeyer interface as defined in raccoondb
// Map an TupleNodeRecord into []byte
// Uses a sha256 to generete the keys
type relationNodeKeyer struct { }

func (t *relationNodeKeyer) Key(node *RelationNode) []byte {
    hasher := sha256.New()
    hasher.Write([]byte(node.Resource))
    hasher.Write([]byte(node.Id))
    hasher.Write([]byte(node.Relation))
    return hasher.Sum(nil)
}

// Min possible key is a slice of sha256bytes min byte val (0)
func (t *relationNodeKeyer) MinKey() []byte {
	k := make([]byte, sha256bytes)
	return k
}

// Min possible key is a slice of sha256bytes max byte val (255)
func (t *relationNodeKeyer) MaxKey() []byte {
	k := make([]byte, sha256bytes)
	for i := range k {
		k[i] = byte(255)
	}
	return k
}

type policyIDer struct{}

func (id *policyIDer) Id(record *domain.PolicyRecord) []byte {
	return []byte(record.Policy.Id)
}

type relationshipDataIDer struct{}

func (id *relationshipDataIDer) Id(record *RelationshipData) []byte {
	return record.RelationshipId
}
