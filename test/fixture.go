package test

import (
    "github.com/sourcenetwork/source-zanzibar/types"
)

// FilesystemFixture returns a fixture for a sample
// policy structure for a permisioned file system
func FilesystemFixture() (types.Policy, []types.Relationship) {
    rb := types.ResourceBuilder{}

    rb.Name("file")
    rb.Relations("owner", "reader")
    rb.Perm("read", "reader+owner")
    rb.Perm("write", "owner")
    file := rb.Build()

    rb.Name("group")
    rb.Relations("member")
    group := rb.Build()
    
    policy := types.Policy{
        Id: "fs",
        Name: "Filesystem Sample Policy",
        Resources: []types.Resource{
            file,
            group,
        },
        Actors: []types.Actor{
            types.NewActor("user"),
        },
    }
    return policy, filesystemRelationships()
}

func filesystemRelationships(polId string) []types.Relationship {
    b := types.RelationshipBuilder(polId)

    return []types.Relationship{
        b.Grant(b.En("group", "admin"), "member", b.En("user", "alice")),
        b.Grant(b.En("group", "staff"), "member", b.En("user", "bob")),

        b.Delegate(b.En("directory", "project"), "dir_owner", b.En("group", "admin"), "member"),
        b.Delegate(b.En("file", "readme"), "reader", b.En("group", "staff"), "member"),

        b.Attribute(b.En("file", "foo"), "parent", b.En("directory", "project")),
        b.Attribute(b.En("file", "readme"), "parent", b.En("directory", "project")),
    }
}
