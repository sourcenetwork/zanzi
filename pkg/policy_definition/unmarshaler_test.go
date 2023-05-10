package policy_definition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalPolicyDefinitionReturnsPolicyDefinition(t *testing.T) {
	const policyYaml = `
    version: '0.1'

    name: Pastebin Policy

    doc: Policy Description

    resources:

      snippet: 
        doc: A text snippet
        relations:
          author:
          reader:
            doc: Reader

        permissions:
          read:
            doc: Reads a snippet
            expr: (reader + author)

          # Comment
          can_comment: 
            expr: (read)

      comment:
        doc: A comment in a snippet
        relations:
          author:
        permissions:
          delete: 
            expr: (author)

    actors:
      user:
        doc: App user

    attributes:
      test: attr
    `

	got, err := UnmarshalPolicyDefinition(policyYaml)

	want := PolicyDefinition{
		Version: "0.1",
		Name:    "Pastebin Policy",
		Doc:     "Policy Description",

		Resources: map[string]*ResourceDefinition{
			"snippet": &ResourceDefinition{
				Name: "snippet",
				Doc:  "A text snippet",
				Relations: map[string]*RelationDefinition{
					"author": &RelationDefinition{
						Name: "author",
						Doc:  "",
					},
					"reader": &RelationDefinition{
						Name: "reader",
						Doc:  "Reader",
					},
				},
				Permissions: map[string]*PermissionDefinition{
					"read": &PermissionDefinition{
						Name: "read",
						Doc:  "Reads a snippet",
						Expr: "(reader + author)",
					},
					"can_comment": &PermissionDefinition{
						Name: "can_comment",
						Doc:  "",
						Expr: "(read)",
					},
				},
			},
			"comment": &ResourceDefinition{
				Name: "comment",
				Doc:  "A comment in a snippet",
				Relations: map[string]*RelationDefinition{
					"author": &RelationDefinition{
						Name: "author",
						Doc:  "",
					},
				},
				Permissions: map[string]*PermissionDefinition{
					"delete": &PermissionDefinition{
						Name: "delete",
						Doc:  "",
						Expr: "(author)",
					},
				},
			},
		},

		Actors: map[string]*ActorDefinition{
			"user": &ActorDefinition{
				Name: "user",
				Doc:  "App user",
			},
		},
		Attributes: map[string]string{
			"test": "attr",
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, &want, got)
}

func TestUnmarshalPolicyDefinitionWithInvalidVersionIdReturnsError(t *testing.T) {
	const policyYaml = `
    version: '0.2'

    name: Pastebin Policy

    resources:
      snippet: 
        relations:
          author:

    actors:
      user:
    `

	got, err := UnmarshalPolicyDefinition(policyYaml)

	assert.Nil(t, got)
	assert.NotNil(t, err)
}

func TestUnmarshalPolicyDefinitionFailsIfDefinitionHasExtraFields(t *testing.T) {
	const policyYaml = `
    version: '0.1'

    extra_field: askfjaskdfj

    name: Pastebin Policy

    resources:
      snippet: 
        relations:
          author:

    actors:
      user:
    `

	got, err := UnmarshalPolicyDefinition(policyYaml)

	assert.Nil(t, got)
	assert.NotNil(t, err)
}
