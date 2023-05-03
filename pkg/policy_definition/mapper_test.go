package policy_definition

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/sourcenetwork/source-zanzibar/types"
)

func TestPolicyFromYamlReturnsCompleteMappedPolicy(t *testing.T) {
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
    got, err := PolicyFromYaml(policyYaml)

    want := &types.Policy {
        Name: "Pastebin Policy",
        Description: "Policy Description",
        Resources: []*types.Resource {
            &types.Resource {
                Name: "comment",
                Description:"A comment in a snippet",
                Relations: []*types.Relation {
                    &types.Relation {
                        Name: "author",
                    },
                },
                Permissions: []*types.Permission {
                    &types.Permission {
                        Name: "delete",
                        PermissionExpr: "(author)",
                    },
                },
            },

            &types.Resource {
                Name: "snippet",
                Description:"A text snippet",
                Relations: []*types.Relation {
                    &types.Relation {
                        Name: "author",
                    },
                    &types.Relation {
                        Name: "reader",
                        Description: "Reader",
                    },
                },
                Permissions: []*types.Permission {
                    &types.Permission {
                        Name: "can_comment",
                        PermissionExpr: "(read)",
                    },
                    &types.Permission {
                        Name: "read",
                        Description: "Reads a snippet",
                        PermissionExpr: "(reader + author)",
                    },
                },
            },
        },
        Actors: []*types.Actor {
            &types.Actor {
                Name: "user",
                Description: "App user",
            },
        },
        Attributes: map[string]string {
            "test": "attr",
        },
    }

    assert.Nil(t, err)
    assert.Equal(t, want, got)
}
