package policy_definition

import (
        "fmt"
        "strings"

        "gopkg.in/yaml.v3"
)

/*
parser should not contain policy validation, that should be in the polciy package

parser is actually just an yaml unmarshaller

add policy validation functions, no need for complex logic in unmarshaler
*/


// Parse receives a policy as a yaml string and returns the parsed Policy object
func UnmarshalPolicyDefinition(policyYaml string) (*PolicyDefinition, error) {
    reader := strings.NewReader(policyYaml)
    decoder := yaml.NewDecoder(reader)
    decoder.KnownFields(true)

    policyDef := PolicyDefinition{}
    err := decoder.Decode(&policyDef)
    if err != nil {
        return nil, fmt.Errorf("could not parse policy: %w", err)
    }

    if !isVersionSupported(policyDef.Version) {
        return nil, fmt.Errorf("invalid policy version %v", policyDef.Version)
    }

    setEntityNames(&policyDef)

    return &policyDef, nil
}

func setEntityNames(p *PolicyDefinition) {
    
    for resName, resource := range p.Resources {
        resource.Name = resName

        for permissionName, permission := range resource.Permissions {
            permission.Name = permissionName
        }

        for relationName, relation := range resource.Relations {
            // The policy may define an empty relation placeholder
            // which means we need to create a default relation
            if relation == nil {
                relDef := RelationDefinition { }
                resource.Relations[relationName] = &relDef
                relation = &relDef
            }

            relation.Name = relationName
        }

    }

    for actorName, actor := range p.Actors {
        // The `actors` section accepts an actor name key
        // as a placeholder definition for an actor.
        if actor == nil {
            actorDef := ActorDefinition { }
            p.Actors[actorName] = &actorDef
            actor = &actorDef
        }
        actor.Name = actorName
    }
}

func isVersionSupported(version string) bool {
    switch version {
    case "0.1":
        return true
    default:
        return false
    }
}
