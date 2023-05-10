package cli

import (
	"github.com/sourcenetwork/zanzi/types"
)

const defaultPolicyId string = "1"
const defaultActorNamespace = "user"

// Fixture models a static data set loaded in Zanzibar's store for auth-cli
type Fixture struct {
	Policies      []types.Policy       `json:"policies"`
	Relationships []types.Relationship `json:"relationships"`
}

func buildDefaultFixture() Fixture {
	return Fixture{
		Policies:      buildPolicy(),
		Relationships: buildRelationships(),
	}
}

// build single default policy for default fixture
func buildPolicy() []types.Policy {
	rb := types.ResourceBuilder{}

	rb.Name("file")
	rb.Relations("owner", "reader", "parent")
	rb.Perm("read", "write + reader")
	rb.Perm("write", "owner + parent->dir_owner")
	file := rb.Build()

	rb.Name("group")
	rb.Relations("member")
	group := rb.Build()

	rb.Name("directory")
	rb.Relations("dir_owner")
	directory := rb.Build()

	pb := types.PolicyBuilder()
	pb.IdNameDescription(defaultPolicyId, "Filesystem Policy", "Auth CLI example filesystem policy")
	pb.Actors(types.NewActor(defaultActorNamespace))
	pb.Resources(file, group, directory)
	pb.Attr("author", "Source Network")
	pol := pb.Build()

	return []types.Policy{
		pol,
	}
}

// build set of relationships for default fixture
func buildRelationships() []types.Relationship {
	b := types.RelationshipBuilder(defaultPolicyId)

	return []types.Relationship{
		b.Grant("group", "admin", "member", defaultActorNamespace, "alice"),
		b.Grant("group", "staff", "member", defaultActorNamespace, "bob"),

		b.Delegate("directory", "project", "dir_owner", "group", "admin", "member"),
		b.Delegate("file", "readme", "reader", "group", "staff", "member"),

		b.Attribute("file", "foo", "parent", "directory", "project"),
		b.Attribute("file", "readme", "parent", "directory", "project"),
	}
}
