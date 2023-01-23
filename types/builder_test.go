package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolicyBuilder(t *testing.T) {
	rb := ResourceBuilder{}
	rb.Name("file")
	rb.Relations("owner", "reader")
	rb.Perm("read", "reader+owner")
	rb.Perm("write", "owner")
	file := rb.Build()

	rb.Name("group")
	rb.Relations("member")
	group := rb.Build()

	pb := PolicyBuilder()
	pb.IdName("1", "test policy")
	pb.Actors(NewActor("staff"), NewActor("guest"))
	pb.Resources(file, group)
	pb.Attr("description", "test policy using builder")

	got := pb.Build()

	want := Policy{
		Id:   "1",
		Name: "test policy",
		Resources: []Resource{
			Resource{
				Name: "file",
				Relations: []Relation{
					Relation{
						Name: "owner",
					},
					Relation{
						Name: "reader",
					},
				},
				Permissions: []Permission{
					Permission{
						Name:       "read",
						Expression: "reader+owner",
					},
					Permission{
						Name:       "write",
						Expression: "owner",
					},
				},
			},
			Resource{
				Name: "group",
				Relations: []Relation{
					Relation{
						Name: "member",
					},
				},
			},
		},
		Actors: []Actor{
			Actor{
				Name: "staff",
			},
			Actor{
				Name: "guest",
			},
		},
		Attributes: map[string]string{
			"description": "test policy using builder",
		},
	}

	assert.Equal(t, got, want)
}
