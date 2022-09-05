package model

// KeyableUset represents a stripped down version of Userset,
// such that it can be used as a map key
type KeyableUset struct {
	namespace string
	objectId  string
	relation  string
}

// Map Userset into a keyable type
func (u *Userset) ToKey() KeyableUset {
	return KeyableUset{
		namespace: u.Namespace,
		objectId:  u.ObjectId,
		relation:  u.Relation,
	}
}

// Restore Key to Userset
func (k *KeyableUset) ToUset() Userset {
	return Userset{
		Namespace: k.namespace,
		ObjectId:  k.objectId,
		Relation:  k.relation,
	}
}
