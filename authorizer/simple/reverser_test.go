package simple

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/model/builder"
	"github.com/sourcenetwork/source-zanzibar/repository"
	"github.com/sourcenetwork/source-zanzibar/repository/maprepo"
)

func revFixture() (repository.NamespaceRepository, repository.TupleRepository) {
	// build test namespace with:
	// - member, owner, parent, reader relations
	// - owner implies reader
	// - owning a parent of an object implies reading the obj
	nr := maprepo.NewNamespaceRepo(
		builder.Namespace("test",
			builder.ThisRelation("member"),
			builder.ThisRelation("owner"),
			builder.ThisRelation("parent"),

			builder.Relation("reader",
				builder.Union(
					builder.This(),
					builder.Union(
						builder.CU("owner"),
						builder.TTU("parent", "owner"),
					),
				),
			),
		),
		builder.Namespace("users"),
	)

	// Alice is a reader of readme
	// bob is an owner of readme
	// charlie is an owner of readme
	tb := builder.WithActorNamespace("users")
	tr := maprepo.NewTupleRepo(
		tb.ObjRel("test", "readme", "reader").User("alice").Build(),
		tb.ObjRel("test", "readme", "owner").Userset("test", "group", "member").Build(),
		tb.ObjRel("test", "secret-doc", "owner").Userset("test", "group", "member").Build(),
		tb.ObjRel("test", "group", "member").User("bob").Build(),
		tb.ObjRel("test", "readme", "parent").Userset("test", "directory", model.EMPTY_REL).Build(),
		tb.ObjRel("test", "directory", "owner").User("charlie").Build(),
	)
	return nr, tr

}

func TestReverserOnDirectEdges(t *testing.T) {
	nr, tr := revFixture()

	user := model.User{
		Userset: &model.AuthNode{
			Namespace: "users",
			ObjectId:  "alice",
			Relation:  model.EMPTY_REL,
		},
		Type: model.UserType_USER,
	}

	ctx := context.Background()
	rev := NewReverser(nr, tr)
	got, err := rev.ReverseLookup(ctx, user)

	require.Nil(t, err)
	want := []model.AuthNode{
		model.AuthNode{
			Namespace: "test",
			ObjectId:  "readme",
			Relation:  "reader",
		},
	}

	usetEquals(t, want, got)
}

func TestReverserWithComputedUsersetRule(t *testing.T) {
	nr, tr := revFixture()

	user := model.User{
		Userset: &model.AuthNode{
			Namespace: "users",
			ObjectId:  "bob",
			Relation:  model.EMPTY_REL,
		},
		Type: model.UserType_USER,
	}

	ctx := context.Background()
	rev := NewReverser(nr, tr)
	got, err := rev.ReverseLookup(ctx, user)

	require.Nil(t, err)
	want := []model.AuthNode{
		model.AuthNode{
			Namespace: "test",
			ObjectId:  "readme",
			Relation:  "owner",
		},

		{
			Namespace: "test",
			ObjectId:  "readme",
			Relation:  "reader",
		},

		model.AuthNode{
			Namespace: "test",
			ObjectId:  "group",
			Relation:  "member",
		},

		model.AuthNode{
			Namespace: "test",
			ObjectId:  "secret-doc",
			Relation:  "reader",
		},

		model.AuthNode{
			Namespace: "test",
			ObjectId:  "secret-doc",
			Relation:  "owner",
		},
	}

	usetEquals(t, want, got)
}

func TestReverserWithTTURule(t *testing.T) {
	nr, tr := revFixture()

	user := model.User{
		Userset: &model.AuthNode{
			Namespace: "users",
			ObjectId:  "charlie",
			Relation:  model.EMPTY_REL,
		},
		Type: model.UserType_USER,
	}

	ctx := context.Background()
	rev := NewReverser(nr, tr)
	got, err := rev.ReverseLookup(ctx, user)

	require.Nil(t, err)
	want := []model.AuthNode{
		{
			Namespace: "test",
			ObjectId:  "readme",
			Relation:  "reader",
		},

		model.AuthNode{
			Namespace: "test",
			ObjectId:  "directory",
			Relation:  "owner",
		},

		model.AuthNode{
			Namespace: "test",
			ObjectId:  "directory",
			Relation:  "reader",
		},
	}

	usetEquals(t, want, got)
}

func usetEquals(t *testing.T, want, got []model.AuthNode) {

	wantMap := make(map[model.KeyableUset]struct{})
	gotMap := make(map[model.KeyableUset]struct{})

	for _, uset := range want {
		wantMap[uset.ToKey()] = struct{}{}
	}

	for _, uset := range got {
		gotMap[uset.ToKey()] = struct{}{}
	}

	require.Equal(t, wantMap, gotMap)
}
