package tuple

import (
    "google.golang.org/protobuf/proto"
)

type TupleBuilder[T proto.Message] struct {
	ActorNamespace string
        Partition string
}

// Return Tuple which delegates a relation to another relation source (ie userset)
func (b *TupleBuilder[T]) Delegate(srcNamespace, srcObjId, relation, dstNamespace, dstObj, dstRelation string) Tuple[T] {
    return Tuple[T] {
        Partition: b.Partition,
        Source: OR(srcNamespace, srcObjId, relation, NodeType_RELATION_SOURCE),
        Dest: OR(dstNamespace, dstObj, dstRelation, NodeType_RELATION_SOURCE),
    }
}

// Return Tuple which grants actorId relation to objId
func (b *TupleBuilder[T]) Grant(namespace, objId, relation, actorId string) Tuple[T] {
    return Tuple[T] {
        Partition: b.Partition,
        Source: OR(namespace, objId, relation, NodeType_RELATION_SOURCE),
        Dest: OR(b.ActorNamespace, actorId, "", NodeType_ACTOR),
    }
}

// Return Tuple which grants ALL actors (in actor namespace) relation to objId
func (b *TupleBuilder[T]) GrantAll(namespace, objId, relation string) Tuple[T] {
    return Tuple[T] {
        Partition: b.Partition,
        Source: OR(namespace, objId, relation, NodeType_RELATION_SOURCE),
        Dest: OR(b.ActorNamespace, "", "", NodeType_ALL_ACTORS),
    }
}

// Return Tuple which relates two system objects
func (b *TupleBuilder[T]) Attribute(namespace, objId, relation, dstNamespace, dstObj string) Tuple[T] {
    return Tuple[T] {
        Partition: b.Partition,
        Source: OR(namespace, objId, relation, NodeType_RELATION_SOURCE),
        Dest: OR(dstNamespace, dstObj, "", NodeType_OBJECT),
    }
}

func OR(ns, obj, rel string, t NodeType) TupleNode {
	return TupleNode{
		Namespace: ns,
		Id:  obj,
		Relation:  rel,
                Type: t,
        }
}
