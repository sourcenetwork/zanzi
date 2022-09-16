package builder

import (
	"github.com/sourcenetwork/source-zanzibar/model"
)

type TupleBuilder struct {
	objRel        *model.Userset
	user          *model.User
	userNamespace string
}

func WithUserNamespace(ns string) TupleBuilder {
	b := TupleBuilder{}
	b.userNamespace = ns
	return b
}

func (b *TupleBuilder) ObjRel(namespace, objectId, relation string) *TupleBuilder {
	b.objRel = &model.Userset{
		Namespace: namespace,
		ObjectId:  objectId,
		Relation:  relation,
	}
	return b
}

func (b *TupleBuilder) User(userId string) *TupleBuilder {
	uset := &model.Userset{
		Namespace: model.USERS_NAMESPACE,
		ObjectId:  userId,
		Relation:  model.EMPTY_REL,
	}
	b.user = &model.User{
		Type:    model.UserType_USER,
		Userset: uset,
	}
	return b
}

func (b *TupleBuilder) Userset(namespace, objectId, relation string) *TupleBuilder {
	uset := &model.Userset{
		Namespace: namespace,
		ObjectId:  objectId,
		Relation:  relation,
	}
	b.user = &model.User{
		Type:    model.UserType_USER_SET,
		Userset: uset,
	}
	return b
}

func (b *TupleBuilder) Build() model.Tuple {
	return model.Tuple{
		ObjectRel: b.objRel,
		User:      b.user,
	}
}

func Userset(ns, obj, rel string) model.Userset {
	return model.Userset{
		Namespace: ns,
		ObjectId:  obj,
		Relation:  rel,
	}
}

func User(ns, obj string) model.Userset {
	return model.Userset{
		Namespace: ns,
		ObjectId:  obj,
		Relation:  model.EMPTY_REL,
	}
}
