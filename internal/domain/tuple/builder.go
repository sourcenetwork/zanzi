package tuple

import (
    "google.golang.org/protobuf/proto"
)

type TupleBuilder[T proto.Message] struct {
	ActorNamespace string
}

func (b *TupleBuilder[T]) Delegate(srcNamespace, srcObjId, relation, dstNamespace, dstObj, dstRelation string) Tuple[T] {
    return Tuple[T] {
        ObjectRel: OR(srcNamespace, srcObjId, relation),
        Actor: OR(dstNamespace, dstObj, dstRelation),
        Type: RelType_DELEGATION,
    }
}

func (b *TupleBuilder[T]) Grant(namespace, objId, relation, actorId string) Tuple[T] {
    return Tuple[T] {
        ObjectRel: OR(namespace, objId, relation),
        Actor: OR(b.ActorNamespace, actorId, ""),
        Type: RelType_GRANT,
    }
}

func (b *TupleBuilder[T]) Attribute(namespace, objId, relation, dstNamespace, dstObj string) Tuple[T] {
    return Tuple[T] {
        ObjectRel: OR(namespace, objId, relation),
        Actor: OR(dstNamespace, dstObj, ""),
        Type: RelType_OBJECT_REL,
    }
}

func OR(ns, obj, rel string) ObjRel {
	return ObjRel{
		Namespace: ns,
		Id:  obj,
		Relation:  rel,
        }
}
