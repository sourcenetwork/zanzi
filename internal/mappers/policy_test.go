package mappers

import (
    "testing"

	"github.com/stretchr/testify/assert"

    "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    "github.com/sourcenetwork/source-zanzibar/types"
)

func TestPublicPolicyMapping(t *testing.T) {
    rb := types.ResourceBuilder{}
    rb.Name("file")
    rb.Relations("owner", "reader")
    rb.Perm("read", "reader+owner")
    rb.Perm("write", "owner")
    file := rb.Build()

    rb.Name("group")
    rb.Relations("member")
    group := rb.Build()

    pb := types.PolicyBuilder()
    pb.IdNameDescription("1", "test policy", "test policy description")
    pb.Actors(types.NewActor("staff", types.Validator_STRING), types.NewActor("guest", types.Validator_NUMBER))
    pb.Resources(file, group)
    pb.Attr("foo", "bar")
    publicPolicy := pb.Build()
   
    mapper := PolicyMapper{}
    got, err := mapper.ToInternal(publicPolicy)

    assert.Nil(t, err)

    want := policy.Policy{
        Id: "1",
        Name: "test policy",
        Description: "test policy description",
        Resources: []*policy.Resource{
            policy.BuildResource(
                "file",
                policy.ThisRelation("owner"),
                policy.ThisRelation("reader"),
                policy.BuildPerm("read", policy.Union(policy.CU("reader"), policy.CU("owner")), "reader+owner"),
                policy.BuildPerm("write", policy.CU("owner"), "owner"),
            ),
            policy.BuildResource(
                "group",
                policy.ThisRelation("member"),
            ),
        },
        Actors: []*policy.Actor{
            policy.NewActor("staff", policy.ActorIdType_STRING),
            policy.NewActor("guest", policy.ActorIdType_NUMBER),
        },
        Attributes: map[string]string {
            "foo": "bar",
        },
    }
    assert.Equal(t, got, want)
}

func TestInternalPolicyReverseMapping(t *testing.T) {
}
