package kv_store

import (
	"context"
	"fmt"

	rcdb "github.com/sourcenetwork/raccoondb"

	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/internal/utils"
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/types"
)

var _ policy.Repository = (*policyRepository)(nil)

func newPolicyRepository(store *KVStore) policyRepository {
	nodeKeyer := relationNodeKeyer{}
	ider := rcdb.IderFromNodeKeyer[*Relationship, *RelationNode](&nodeKeyer)
	return policyRepository{
		kvStore:          store,
		mapper:           relationshipMapper{},
		relationshipIDer: ider,
	}
}

type policyRepository struct {
	kvStore          *KVStore
	mapper           relationshipMapper
	relationshipIDer rcdb.Ider[*Relationship]
}

func (r *policyRepository) SetPolicy(ctx context.Context, record *domain.PolicyRecord) (types.RecordFound, error) {
	store := r.kvStore.getPolicyStore()

	has, err := store.HasById([]byte(record.Policy.Id))
	if err != nil {
		return false, err
	}

	err = store.SetObject(record)
	if err != nil {
		return false, err
	}

	return types.RecordFound(has), nil
}

func (r *policyRepository) GetPolicy(ctx context.Context, id string) (*domain.PolicyRecord, error) {
	store := r.kvStore.getPolicyStore()
	opt, err := store.GetObject([]byte(id))
	if err != nil {
		return nil, err
	}
	if opt.IsEmpty() {
		return nil, nil
	} else {
		return opt.Value(), nil
	}
}

func (r *policyRepository) DeletePolicy(ctx context.Context, id string) (types.RecordFound, error) {
	store := r.kvStore.getPolicyStore()

	found, err := store.HasById([]byte(id))
	if err != nil {
		return false, err
	}
	if !found {
		return false, nil
	}

	err = store.DeleteById([]byte(id))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *policyRepository) ListPolicyIds(context.Context) ([]string, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *policyRepository) SetRelationship(ctx context.Context, record *domain.RelationshipRecord) (bool, error) {
	relationshipStore := r.kvStore.getRelationshipStore(record.PolicyId)
	relationshipDataStore := r.kvStore.getRelationshipDataStore(record.PolicyId)

	// FIXME wrap in tx
	relationship := r.mapper.ToInternal(record.Relationship)

	opt, err := relationshipStore.Get(relationship.GetSource(), relationship.GetDest())
	updated := !opt.IsEmpty()
	if err != nil {
		return false, fmt.Errorf("kvstore: has relationship: %w", err)
	}

	err = relationshipStore.Set(&relationship)
	if err != nil {
		return false, fmt.Errorf("kvstore: set relationship: %w", err)
	}

	data := RelationshipData{
		RelationshipId: r.relationshipIDer.Id(&relationship),
		AppData:        record.AppData,
		CreatedAt:      record.CreatedAt,
	}
	err = relationshipDataStore.SetObject(&data)
	if err != nil {
		return false, fmt.Errorf("kvstore: set relationship appdata: %w", err)
	}

	return updated, nil
}

func (r *policyRepository) DeleteRelationship(ctx context.Context, policyId string, spec *domain.Relationship) (types.RecordFound, error) {
	relationshipStore := r.kvStore.getRelationshipStore(policyId)
	relationshipDataStore := r.kvStore.getRelationshipDataStore(policyId)

	// FIXME wrap in tx
	relationship := r.mapper.ToInternal(spec)
	id := r.relationshipIDer.Id(&relationship)

	opt, err := relationshipStore.Get(relationship.GetSource(), relationship.GetDest())
	found := !opt.IsEmpty()
	if err != nil {
		return false, err
	}

	err = relationshipStore.Delete(&relationship)
	if err != nil {
		return false, err
	}

	err = relationshipDataStore.DeleteById(id)
	if err != nil {
		return false, err
	}

	return types.RecordFound(found), nil
}

func (r *policyRepository) FindRelationships(ctx context.Context, policyId string, selector *domain.RelationshipSelector) ([]*domain.Relationship, error) {
	relationshipStore := r.kvStore.getRelationshipStore(policyId)

	rels, err := relationshipStore.List()
	if err != nil {
		return nil, err
	}

	spec, err := policy.NewSelectorSpec(selector)
	if err != nil {
		return nil, err
	}

	mapped := utils.MapSlice(rels, func(relationship *Relationship) *domain.Relationship {
		mapped := r.mapper.FromInternal(relationship)
		return &mapped
	})
	filtered := utils.Filter(mapped, func(relationship *domain.Relationship) bool { return spec.Satisfies(relationship) })
	return filtered, nil
}

func (r *policyRepository) GetRelationship(ctx context.Context, policyId string, relationship *domain.Relationship) (*domain.RelationshipRecord, error) {
	relationshipDataStore := r.kvStore.getRelationshipDataStore(policyId)

	internalRelationship := r.mapper.ToInternal(relationship)
	id := r.relationshipIDer.Id(&internalRelationship)
	dataOpt, err := relationshipDataStore.GetObject(id)
	if err != nil {
		return nil, err
	}
	if dataOpt.IsEmpty() {
		return nil, nil
	}
	record := r.mapper.FromInternalRecord(policyId, &internalRelationship, dataOpt.Value())
	return &record, nil
}

func (r *policyRepository) FindRelationshipRecords(ctx context.Context, policyId string, selector *domain.RelationshipSelector) ([]*domain.RelationshipRecord, error) {
	relationships, err := r.FindRelationships(ctx, policyId, selector)
	if err != nil {
		return nil, err
	}

	records, err := utils.MapSliceErr(relationships, func(relationship *domain.Relationship) (*domain.RelationshipRecord, error) {
		return r.GetRelationship(ctx, policyId, relationship)
	})
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (r *policyRepository) DeleteRelationships(ctx context.Context, policyId string, selector *domain.RelationshipSelector) (uint64, error) {
	relationships, err := r.FindRelationships(ctx, policyId, selector)
	if err != nil {
		return 0, err
	}

	founds, err := utils.MapSliceErr(relationships, func(relationship *domain.Relationship) (types.RecordFound, error) {
		return r.DeleteRelationship(ctx, policyId, relationship)
	})
	if err != nil {
		return 0, err
	}

	removed := len(utils.MapSlice(founds, utils.Identity[types.RecordFound]))
	return uint64(removed), nil
}
