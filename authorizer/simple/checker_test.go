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

func fixture() (repository.NamespaceRepository, repository.TupleRepository) {
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

func TestCheckOnDirectRelation(t *testing.T) {
	nr, tr := fixture()

	ctx := context.Background()
	obj := builder.Userset("test", "readme", "reader")
	user := builder.User("users", "alice")
	checker := NewChecker(nr, tr)
	got, err := checker.Check(ctx, obj, user)

	require.Nil(t, err)
	require.Equal(t, true, got)
}

func TestCheckOnNestedDirectRelation(t *testing.T) {
	nr, tr := fixture()

	ctx := context.Background()
	obj := builder.Userset("test", "readme", "owner")
	user := builder.User("users", "bob")
	checker := NewChecker(nr, tr)
	got, err := checker.Check(ctx, obj, user)

	require.Nil(t, err)
	require.Equal(t, true, got)
}

func TestCheckComputedUsersetExpansion(t *testing.T) {
	nr, tr := fixture()

	ctx := context.Background()
	obj := builder.Userset("test", "readme", "reader")
	user := builder.User("users", "bob")
	checker := NewChecker(nr, tr)
	got, err := checker.Check(ctx, obj, user)

	require.Nil(t, err)
	require.Equal(t, true, got)
}

func TestCuExpandOnUndeclaredNodeReturnsTrue(t *testing.T) {
	// this is a subtler test because the node
	// ("secret-doc", "reader") was not defined in any tuple.
	// the check call must still succeed due to userset rewerite semantics
	// it's possible that an object might not have any direct readers,
	// only inderect readers, through CU or other rewrite rules
	nr, tr := fixture()

	ctx := context.Background()
	obj := builder.Userset("test", "secret-doc", "reader")
	user := builder.User("users", "bob")
	checker := NewChecker(nr, tr)
	got, err := checker.Check(ctx, obj, user)

	require.Nil(t, err)
	require.Equal(t, true, got)
}

func TestExpandTtuRule(t *testing.T) {
	nr, tr := fixture()

	ctx := context.Background()
	obj := builder.Userset("test", "readme", "reader")
	user := builder.User("users", "charlie")
	checker := NewChecker(nr, tr)
	got, err := checker.Check(ctx, obj, user)

	require.Nil(t, err)
	require.Equal(t, true, got)
}
