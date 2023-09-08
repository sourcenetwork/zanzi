package store

import (
	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/internal/relation_graph"
)

type Store interface {
    GetPolicyRepository() policy.Repository
    GetRelationNodeRepository() relation_graph.NodeRepository
}
