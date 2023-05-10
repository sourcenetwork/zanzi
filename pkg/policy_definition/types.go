// package policy_definition defines a policy definition types
// machinery to unmarshal it from a yaml file
// and a mapper to transform from a definition to a permission
package policy_definition

type PolicyDefinition struct {
	Version    string                         `yaml:"version"`
	Name       string                         `yaml:"name"`
	Doc        string                         `yaml:"doc"`
	Resources  map[string]*ResourceDefinition `yaml:"resources"`
	Actors     map[string]*ActorDefinition    `yaml:"actors"`
	Attributes map[string]string              `yaml:"attributes"`
}

type ResourceDefinition struct {
	Name        string
	Doc         string                           `yaml:"doc"`
	Permissions map[string]*PermissionDefinition `yaml:"permissions"`
	Relations   map[string]*RelationDefinition   `yaml:"relations"`
}

type PermissionDefinition struct {
	Name string
	Doc  string `yaml:"doc"`
	Expr string `yaml:"expr"`
}

type RelationDefinition struct {
	Name string
	Doc  string `yaml:"doc"`
	//Types []string `yaml:""`
}

type ActorDefinition struct {
	Name string
	Doc  string `yaml:"doc"`
}
