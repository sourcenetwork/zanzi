package mappers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
	"github.com/sourcenetwork/source-zanzibar/types"
)

func TestPublicPolicyMapping(t *testing.T) {
	time := time.Date(2023, time.January, 01, 0, 0, 0, 0, time.UTC)
	timePb := timestamppb.New(time)

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
	publicPolicy.CreatedAt = time

	mapper := PolicyMapper{}
	got, err := mapper.ToInternal(publicPolicy)

	assert.Nil(t, err)

	want := policy.Policy{
		Id:          "1",
		Name:        "test policy",
		Description: "test policy description",
		CreatedAt:   timePb,
		Resources: []*policy.Resource{
			policy.NewResource(
				"file",
				policy.ThisRelation("owner"),
				policy.ThisRelation("reader"),
				policy.NewRelation("read", policy.Union(policy.CU("reader"), policy.CU("owner")), "reader+owner"),
				policy.NewRelation("write", policy.CU("owner"), "owner"),
			),
			policy.NewResource(
				"group",
				policy.ThisRelation("member"),
			),
		},
		Actors: []*policy.Actor{
			policy.NewActor("staff", policy.ActorIdType_STRING),
			policy.NewActor("guest", policy.ActorIdType_NUMBER),
		},
		Attributes: map[string]string{
			"foo": "bar",
		},
	}
	assert.Equal(t, got, want)
}

func TestInternalPolicyReverseMapping(t *testing.T) {
	time := time.Date(2023, time.January, 01, 0, 0, 0, 0, time.UTC)
	timePb := timestamppb.New(time)

	policy := policy.Policy{
		Id:          "1",
		Name:        "Test Policy",
		Description: "test policy description",
		CreatedAt:   timePb,
		Resources: []*policy.Resource{
			&policy.Resource{
				Name: "file",
				Relations: []*policy.Relation{
					&policy.Relation{
						Name:           "owner",
						RewriteExpr:    "_this",
						ExpressionTree: policy.ThisTree(),
					},
					&policy.Relation{
						Name:           "reader",
						RewriteExpr:    "_this",
						ExpressionTree: policy.ThisTree(),
					},
					&policy.Relation{
						Name:           "read",
						RewriteExpr:    "owner + reader",
						ExpressionTree: policy.Union(policy.CU("owner"), policy.CU("reader")),
					},
				},
			},
			policy.NewResource("group",
				policy.ThisRelation("member"),
			),
		},
		Actors: []*policy.Actor{
			policy.NewActor("user"),
		},
		Attributes: map[string]string{
			"foo": "bar",
		},
	}

	rb := types.ResourceBuilder{}
	rb.Name("file")
	rb.Relations("owner", "reader")
	rb.Perm("read", "owner + reader")
	file := rb.Build()

	rb.Name("group")
	rb.Relations("member")
	group := rb.Build()

	pb := types.PolicyBuilder()
	pb.IdNameDescription("1", "Test Policy", "test policy description")
	pb.Actors(types.NewActor("user"))
	pb.Resources(file, group)
	pb.Attr("foo", "bar")
	want := pb.Build()
	want.CreatedAt = time

	mapper := PolicyMapper{}
	got := mapper.FromInternal(&policy)

	assert.Equal(t, want, got)
}
