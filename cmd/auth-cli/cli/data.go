package cli

import (
	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/model/builder"
	"github.com/sourcenetwork/source-zanzibar/repository"
	"github.com/sourcenetwork/source-zanzibar/repository/maprepo"
)

func buildSampleNamespaceRepo() repository.NamespaceRepository {
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
	return nr
}

func buildSampleTupleRepo() repository.TupleRepository {
	// Alice is a reader of readme
	// bob is an owner of readme
	// charlie is an owner of readme
	tb := builder.WithUserNamespace("users")
	tr := maprepo.NewTupleRepo(
		tb.ObjRel("test", "readme", "reader").User("alice").Build(),
		tb.ObjRel("test", "readme", "owner").Userset("test", "group", "member").Build(),
		tb.ObjRel("test", "secret-doc", "owner").Userset("test", "group", "member").Build(),
		tb.ObjRel("test", "group", "member").User("bob").Build(),
		tb.ObjRel("test", "readme", "parent").Userset("test", "directory", model.EMPTY_REL).Build(),
		tb.ObjRel("test", "directory", "owner").User("charlie").Build(),
	)
	return tr
}
