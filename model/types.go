package model


// KeyableUset represents a stripped down version of Userset,
// such that it can be used as a map key
type KeyableUset struct {
	Namespace string
	ObjectId  string
	Relation  string
}

func ToKey(userset Userset) KeyableUset {
	return KeyableUset{
		Namespace: userset.Namespace,
		ObjectId:  userset.ObjectId,
		Relation:  userset.Relation,
	}
}

func ToUset(key KeyableUset) Userset {
	return Userset{
		Namespace: key.Namespace,
		ObjectId:  key.ObjectId,
		Relation:  key.Relation,
	}
}
