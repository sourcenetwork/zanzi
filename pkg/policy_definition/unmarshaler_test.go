package policy_definition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalPolicyDefinitionFailsIfDefinitionHasExtraFields(t *testing.T) {
	const policyYaml = `
    extra_field: askfjaskdfj

    name: Pastebin Policy

    resources:
      snippet: 
        relations:
          author:
    `

	got, err := UnmarshalPolicyDefinition(policyYaml)

	assert.Nil(t, got)
	assert.NotNil(t, err)
}
