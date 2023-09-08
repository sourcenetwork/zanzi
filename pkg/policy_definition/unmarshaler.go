package policy_definition

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// UnmarshalPolicyDefinition receives a policy as a yaml string and returns the parsed Policy object
func UnmarshalPolicyDefinition(policyYaml string) (*PolicyDefinition, error) {
	reader := strings.NewReader(policyYaml)
	decoder := yaml.NewDecoder(reader)
	decoder.KnownFields(true)

	policyDef := PolicyDefinition{}
	err := decoder.Decode(&policyDef)
	if err != nil {
		return nil, fmt.Errorf("could not parse policy: %w", err)
	}

	setEntityNames(&policyDef)

	return &policyDef, nil
}

func setEntityNames(p *PolicyDefinition) {

	for resName, resource := range p.Resources {
		if resource == nil {
			resDef := ResourceDefinition{}
			p.Resources[resName] = &resDef
			resource = &resDef
		}

		resource.name = resName

		for relationName, relation := range resource.Relations {
			// The policy may define an empty relation placeholder
			// which means we need to create a default relation
			if relation == nil {
				relDef := RelationDefinition{}
				resource.Relations[relationName] = &relDef
				relation = &relDef
			}

			relation.name = relationName
		}

	}
}
