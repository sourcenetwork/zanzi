package stores

// PolicyStore abstract interfacing with namespace storage.
//
// Contract: methods should return instance of `EntityNotFound`
// when some operation failed because the lookup resulted in an empty set.
type PolicyStore interface {
	GetNamespace(namespace string) (utils.Option[model.Namespace], error)

	SetNamespace(namespace model.Namespace) (model.Namespace, error)

	RemoveNamespace(namespace string) error

	// Return a Relation definition from a namespace
	GetRelation(namespace, relation string) (model.Relation, error)

	// Return all relations which reference the given `relation`.
	GetReferrers(namespace, relation string) ([]model.Relation, error)
}
