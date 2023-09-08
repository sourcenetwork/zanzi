package kv_store

import (
	"context"

	"github.com/sourcenetwork/zanzi/internal/relation_graph"
	"github.com/sourcenetwork/zanzi/internal/utils"
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/types"
)

var _ relation_graph.NodeRepository = (*nodeRepository)(nil)

func newNodeRepository(store *KVStore) nodeRepository {
	return nodeRepository{
		kvStore: store,
		mapper:  relationshipMapper{},
	}
}

type nodeRepository struct {
	kvStore *KVStore
	mapper  relationshipMapper
}

func (r *nodeRepository) GetSucessors(ctx context.Context, policyId string, node *domain.EntitySetNode) ([]*domain.RelationNode, error) {
	store := r.kvStore.getRelationshipStore(policyId)
	relNode := &domain.RelationNode{
		Node: &domain.RelationNode_EntitySet{
			EntitySet: node,
		},
	}
	internalNode := r.mapper.FromPublicRelationNode(relNode)

	sucessors, err := store.GetSucessors(&internalNode)
	if err != nil {
		return nil, err
	}

	return utils.MapSlice(sucessors, func(relationship *Relationship) *domain.RelationNode {
		node := r.mapper.ToPublicRelationNode(relationship.Dest)
		return &node
	}), nil
}

func (r *nodeRepository) ListEdges(ctx context.Context, policyId string) ([]types.Pair[*domain.RelationNode, *domain.RelationNode], error) {
	store := r.kvStore.getRelationshipStore(policyId)

	relationships, err := store.List()
	if err != nil {
		return nil, err
	}

	pairs := utils.MapSlice(relationships, func(relationship *Relationship) types.Pair[*domain.RelationNode, *domain.RelationNode] {
		source := r.mapper.ToPublicRelationNode(relationship.Source)
		dest := r.mapper.ToPublicRelationNode(relationship.Dest)
		return types.NewPair(&source, &dest)
	})
	return pairs, nil
}
