package test

import (
	"github.com/sourcenetwork/source-zanzibar/types"
)

const (
	ActorNamespace string = "user"
	FsPolicyId            = "1"
)

// FilesystemFixture returns a fixture for a sample
// policy structure for a permisioned file system
func FilesystemFixture() (types.Policy, []types.Relationship) {
	rb := types.ResourceBuilder{}

	rb.Name("file")
	rb.Relations("owner", "reader", "contains")
	rb.Perm("read", "(write + reader)")
	rb.Perm("write", "(owner + contains->write)")
	file := rb.Build()

	rb.Name("group")
	rb.Relations("member")
	group := rb.Build()

	policy := types.Policy{
		Id:   FsPolicyId,
		Name: "Filesystem Sample Policy",
		Resources: []types.Resource{
			file,
			group,
		},
		Actors: []types.Actor{
			types.NewActor(ActorNamespace),
		},
	}
	return policy, filesystemRelationships(FsPolicyId)
}

// filesystemRelationships models the following file structure:
// /
// |-- trent
//
//	|-- file.txt
//
// |-- src
//
//	|-- foo.c
//	|-- bar.c
//
// the actors are:
// alice, bob and charlie - members of engineering
// root - member of admin
// trent - simple actor
func filesystemRelationships(polId string) []types.Relationship {
	b := types.RelationshipBuilder(polId)

	return []types.Relationship{
		b.Attribute("file", "/src", "contains", "file", "/"),
		b.Attribute("file", "/trent", "contains", "file", "/"),
		b.Attribute("file", "/trent/file.txt", "contains", "file", "/trent"),
		b.Attribute("file", "/src/foo.c", "contains", "file", "/src"),
		b.Attribute("file", "/src/bar.c", "contains", "file", "/src"),

		b.Delegate("file", "/src", "owner", "group", "engineering", "member"),
		b.Delegate("file", "/", "owner", "group", "admin", "member"),
		b.Grant("file", "/trent", "owner", ActorNamespace, "trent"),

		b.Grant("group", "engineering", "member", ActorNamespace, "alice"),
		b.Grant("group", "engineering", "member", ActorNamespace, "bob"),
		b.Grant("group", "engineering", "member", ActorNamespace, "charlie"),
		b.Grant("group", "admin", "member", ActorNamespace, "root"),

		b.Grant("file", "/src/foo.c", "owner", ActorNamespace, "bob"),
		b.Grant("file", "/src/bar.c", "owner", ActorNamespace, "alice"),
		b.Grant("file", "/trent/file.txt", "owner", ActorNamespace, "trent"),
		b.Grant("file", "/trent/file.txt", "reader", ActorNamespace, "alice"),
	}
}
