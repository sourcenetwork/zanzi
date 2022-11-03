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
		builder.Namespace("staff"),

		builder.Namespace("department", builder.ThisRelation("head")),

		builder.Namespace("team", 
                    builder.ThisRelation("dept"),
                    builder.ThisRelation("member"),
                    builder.Relation("participant",
                        builder.Union(
                        builder.CU("member"),
                        builder.TTU("dept", "head"),
                    ),
                ),
            ),

            builder.Namespace("channel", 
                builder.ThisRelation("owner_team"),
                builder.ThisRelation("collaborator"),
                builder.ThisRelation("guest"),
                builder.Relation("admin", builder.TTU("owner_team", "participant")),
                builder.Relation("read", builder.Union(builder.CU("write"), builder.CU("guest"))),
                builder.Relation("write", builder.Union(builder.CU("admin"), builder.CU("collaborator"))),
            ),
	)
	return nr
}

func buildSampleTupleRepo() repository.TupleRepository {
	tb := builder.WithUserNamespace("staff")
	tr := maprepo.NewTupleRepo(
		tb.ObjRel("department", "eng", "head").User("John").Build(),

		tb.ObjRel("team", "db", "dept").Userset("department", "eng", model.EMPTY_REL).Build(),
		tb.ObjRel("team", "db", "member").User("Andy").Build(),
		tb.ObjRel("team", "db", "member").User("Fred").Build(),
		tb.ObjRel("team", "db", "member").User("Shahzad").Build(),
		tb.ObjRel("team", "db", "member").User("Orpheus").Build(),

		tb.ObjRel("team", "protocol", "dept").Userset("department", "eng", model.EMPTY_REL).Build(),
		tb.ObjRel("team", "protocol", "member").User("Bruno").Build(),

		tb.ObjRel("channel", "dev-db", "owner_team").Userset("team", "db", model.EMPTY_REL).Build(),
		tb.ObjRel("channel", "dev-db", "collaborator").Userset("team", "protocol", model.EMPTY_REL).Build(),
		tb.ObjRel("channel", "dev-db", "guest").User("Addo").Build(),
		tb.ObjRel("channel", "dev-db", "guest").User("Chris").Build(),
	)
	return tr
}
