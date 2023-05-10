package integration

import (
	"log"
	"testing"

	rcdb "github.com/sourcenetwork/raccoondb"
	"github.com/stretchr/testify/assert"

	zanzi "github.com/sourcenetwork/zanzi"
	"github.com/sourcenetwork/zanzi/test"
	"github.com/sourcenetwork/zanzi/types"
)

func setup() types.SimpleClient {
	kv := rcdb.NewMemKV()

	client := zanzi.NewSimpleFromKV(kv)

	policy, relationships := test.FilesystemFixture()

	relServ := client.GetRelationshipService()
	for _, relationship := range relationships {
		err := relServ.Set(relationship)
		if err != nil {
			log.Panicf("failed setting relationship: %v", err)
		}
	}

	polServ := client.GetPolicyService()
	err := polServ.Set(policy)
	if err != nil {
		log.Panicf("failed setting policy: %v", err)
	}
	return client
}

func TestActorsDirectlyRelatedToObjectCanOperateOverResource(t *testing.T) {
	authorizer := setup().GetAuthorizer()

	ok, err := authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/src/foo.c"),
		"owner",
		types.NewEntity(test.ActorNamespace, "bob"),
	)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestActorsNotDirectlyRelatedToObjectCannotOperateOverResource(t *testing.T) {
	authorizer := setup().GetAuthorizer()

	ok, err := authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/trent"),
		"write",
		types.NewEntity(test.ActorNamespace, "alice"),
	)
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/trent"),
		"read",
		types.NewEntity(test.ActorNamespace, "alice"),
	)
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestActorReachedThroughPointerChasingCanOperateOverResource(t *testing.T) {
	authorizer := setup().GetAuthorizer()

	ok, err := authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/src"),
		"owner",
		types.NewEntity(test.ActorNamespace, "alice"),
	)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestUnknownActorCannotOperateOverResource(t *testing.T) {
	authorizer := setup().GetAuthorizer()

	ok, err := authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/"),
		"owner",
		types.NewEntity(test.ActorNamespace, "unknown"),
	)
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestPermissionExpandsRelation(t *testing.T) {
	authorizer := setup().GetAuthorizer()

	ok, err := authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/src/foo.c"),
		"write",
		types.NewEntity(test.ActorNamespace, "bob"),
	)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/src/foo.c"),
		"read",
		types.NewEntity(test.ActorNamespace, "bob"),
	)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestPermissionExpandsTupleToUserset(t *testing.T) {
	authorizer := setup().GetAuthorizer()

	ok, err := authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/src/foo.c"),
		"write",
		types.NewEntity(test.ActorNamespace, "alice"),
	)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/src/foo.c"),
		"read",
		types.NewEntity(test.ActorNamespace, "alice"),
	)
	assert.Nil(t, err)
	assert.True(t, ok)

	// TTU expands files in directories
	// to be writable by the parent, not owned.
	// therefore, since /src/foo.c is not directly owned
	// by engineering or alice, she does not own it
	ok, err = authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/src/foo.c"),
		"owner",
		types.NewEntity(test.ActorNamespace, "alice"),
	)
	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestPermissionsUnion(t *testing.T) {
	authorizer := setup().GetAuthorizer()

	// alice and trent should be able to read file.txt
	// since read is defined as reader + write
	// given that trent is owner and alice is reader,
	// the result set must contain both
	ok, err := authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/trent/file.txt"),
		"read",
		types.NewEntity(test.ActorNamespace, "alice"),
	)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = authorizer.Check(test.FsPolicyId,
		types.NewEntity("file", "/trent/file.txt"),
		"read",
		types.NewEntity(test.ActorNamespace, "trent"),
	)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestPermissionsIntersection(t *testing.T) {
	// TODO
}

func TestPermissionsSubtraction(t *testing.T) {
	// TODO
}
