// package policy_definition defines a compact policy representation
//
// The definitions within this package are entirely compatable with
// the main Policy definition and are used to conveniently
// express a policy as a yaml document.
package policy_definition

// UniversalSetType is a reserved symbol for a `Relation` type list representing that any type is accepted.
const UniversalSetType string = "*"

// TypeRelationSeparator is a separator used within Relation type
// contraints to indicate an optional relation within a constraint.
// eg. the type constraint `group:member` would mean that only
// members from entity_set composed by the resource "group" and
// relation "member" are allowed.
const TypeRelationSeparator string = ":"

// PolicyDefinition represents a compact definition of a domain.Policy
type PolicyDefinition struct {
	// User defined policy ID
	Id string `yaml:"id"`
	//Version    string                         `yaml:"version"`

	// Display name for Policy
	Name string `yaml:"name"`

	// User defined doc string
	Doc string `yaml:"doc"`

	// Map of resource name to ResourceDefinition
	Resources map[string]*ResourceDefinition `yaml:"resources"`

	// Attributes is a map of user defined attributes
	// for a Policy
	Attributes map[string]string `yaml:"attributes"`
}

// ResourceDefinition represents a compact definition of a domain.Resource
type ResourceDefinition struct {
	name      string
	Doc       string                         `yaml:"doc"`
	Relations map[string]*RelationDefinition `yaml:"relations"`
}

// RelationDefinition represents a compact definition of a domain.Resource
type RelationDefinition struct {
	name string

	// User defined doc string for Relation
	Doc string `yaml:"doc"`

	// Expr defines a relation rewrite expression
	Expr string `yaml:"expr"`

	// Types is a set of type identifiers for allowed
	// Subjects for relationships of the given Relation.
	//
	// eg. if types is ["user"], only subjects whose
	// resource name are "user" will be allowed,
	// meaning that the Relationship:
	// obj=file:readme, relation={current relation}, subj=user:bob
	// - where {current relation} represents a placeholder name for the Relation being evaluated - is allowed because the subject is of type "user".
	// Now consider the Relationship:
	// obj=file:readme, relation={current relation}, subj=admin:alice,
	// in this example, the subject is not allowed because
	// "alice" is an "admin", not an "user".
	//
	// Futhermore, a type constraint can specify an optional Relation alongside a Resource name.
	// eg. the Type constraint "group:member" is equivalent
	// to stating that only Subjects of type ResourceSet
	// whose Resource is "group" and relation "member" are allowed,
	// such as the Relationship:
	// obj=file:readme, relation={current relation}, subject=(obj=group:engineering, relation=member).
	//
	// Lastly the special symbol defined in the UniversalSetType constant,
	// is interpreted to mean that any Subject is allowed
	// for the Relationships.
	Types []string `yaml:"types"`
}
