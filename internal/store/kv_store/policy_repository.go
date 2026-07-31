package kv_store

import (
	"context"

	rcdb "github.com/sourcenetwork/raccoondb"

	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/internal/utils"
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/errors"
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
		return false, errors.NewWithCause("setting policy", errors.Internal, err)
	}

	err = store.SetObject(record)
	if err != nil {
		return false, errors.NewWithCause("setting policy", errors.Internal, err)
	}

	return types.RecordFound(has), nil
}

func (r *policyRepository) GetPolicy(ctx context.Context, id string) (*domain.PolicyRecord, error) {
	store := r.kvStore.getPolicyStore()
	opt, err := store.GetObject([]byte(id))
	if err != nil {
		return nil, errors.NewWithCause("setting policy", errors.Internal, err)
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
		return false, errors.NewWithCause("deleting policy", errors.Internal, err)
	}
	if !found {
		return false, nil
	}

	err = store.DeleteById([]byte(id))
	if err != nil {
		return false, errors.NewWithCause("deleting policy", errors.Internal, err)
	}
	return true, nil
}

func (r *policyRepository) ListPolicyIds(context.Context) ([]string, error) {
	store := r.kvStore.getPolicyStore()
	ids, err := store.ListIds()
	if err != nil {
		return nil, errors.NewWithCause("listing policy ids", errors.Internal, err)
	}

	return utils.MapSlice(ids, func(id []byte) string {
		return string(id)
	}), nil
}

func (r *policyRepository) SetRelationship(ctx context.Context, record *domain.RelationshipRecord) (bool, error) {
	relationshipStore := r.kvStore.getRelationshipStore(record.PolicyId)
	relationshipDataStore := r.kvStore.getRelationshipDataStore(record.PolicyId)

	// FIXME wrap in tx
	relationship := r.mapper.ToInternal(record.Relationship)

	opt, err := relationshipStore.Get(relationship.GetSource(), relationship.GetDest())
	updated := !opt.IsEmpty()
	if err != nil {
		return false, errors.NewWithCause("setting relationship", errors.Internal, err)
	}

	err = relationshipStore.Set(&relationship)
	if err != nil {
		return false, errors.NewWithCause("setting relationship", errors.Internal, err)
	}

	data := RelationshipData{
		RelationshipId: r.relationshipIDer.Id(&relationship),
		AppData:        record.AppData,
	}
	err = relationshipDataStore.SetObject(&data)
	if err != nil {
		return false, errors.NewWithCause("setting relationship", errors.Internal, err)
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
		return false, errors.NewWithCause("deleting relationship", errors.Internal, err)
	}

	err = relationshipStore.Delete(&relationship)
	if err != nil {
		return false, errors.NewWithCause("deleting relationship", errors.Internal, err)
	}

	err = relationshipDataStore.DeleteById(id)
	if err != nil {
		return false, errors.NewWithCause("deleting relationship", errors.Internal, err)
	}

	return types.RecordFound(found), nil
}

func (r *policyRepository) FindRelationships(ctx context.Context, policyId string, selector *domain.RelationshipSelector) ([]*domain.Relationship, error) {
	relationshipStore := r.kvStore.getRelationshipStore(policyId)
	fetcher := newRelationshipFetcher(relationshipStore, r.mapper)
	result, err := fetcher.Fetch(ctx, selector)
	if err != nil {
		return nil, errors.NewWithCause("looking up relationships", errors.Internal, err)
	}
	return result, nil
}

func (r *policyRepository) GetRelationship(ctx context.Context, policyId string, relationship *domain.Relationship) (*domain.RelationshipRecord, error) {
	relationshipDataStore := r.kvStore.getRelationshipDataStore(policyId)

	internalRelationship := r.mapper.ToInternal(relationship)
	id := r.relationshipIDer.Id(&internalRelationship)
	dataOpt, err := relationshipDataStore.GetObject(id)
	if err != nil {
		return nil, errors.NewWithCause("getting relationship", errors.Internal, err)
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
		return nil, errors.NewWithCause("looking up relationships", errors.Internal, err)
	}

	records, err := utils.MapSliceErr(relationships, func(relationship *domain.Relationship) (*domain.RelationshipRecord, error) {
		return r.GetRelationship(ctx, policyId, relationship)
	})
	if err != nil {
		return nil, errors.NewWithCause("looking up relationships", errors.Internal, err)
	}

	return records, nil
}

func (r *policyRepository) DeleteRelationships(ctx context.Context, policyId string, selector *domain.RelationshipSelector) (uint64, error) {
	relationships, err := r.FindRelationships(ctx, policyId, selector)
	if err != nil {
		return 0, errors.NewWithCause("deleting relationships", errors.Internal, err)
	}

	founds, err := utils.MapSliceErr(relationships, func(relationship *domain.Relationship) (types.RecordFound, error) {
		return r.DeleteRelationship(ctx, policyId, relationship)
	})
	if err != nil {
		return 0, errors.NewWithCause("deleting relationships", errors.Internal, err)
	}

	removed := len(utils.MapSlice(founds, utils.Identity[types.RecordFound]))
	return uint64(removed), nil
}

func (r *policyRepository) ListPolicies(ctx context.Context) ([]*domain.PolicyRecord, error) {
	records, err := r.kvStore.policyStore.List()
	if err != nil {
		return nil, errors.NewWithCause("listing policies", errors.Internal, err)
	}
	return records, nil
}
