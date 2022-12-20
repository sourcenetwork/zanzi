package tuple

type TupleBuilder struct {
	ActorNamespace string
	Partition      string
}

// Return Tuple which delegates a relation to another relation source (ie userset)
func (b *TupleBuilder) Delegate(srcNamespace, srcObjId, relation, dstNamespace, dstObj, dstRelation string) Tuple {
	return Tuple{
		Partition: b.Partition,
		Source:    OR(srcNamespace, srcObjId, relation, NodeType_RELATION_SOURCE),
		Dest:      OR(dstNamespace, dstObj, dstRelation, NodeType_RELATION_SOURCE),
	}
}

// Return Tuple which grants actorId relation to objId
func (b *TupleBuilder) Grant(namespace, objId, relation, actorId string) Tuple {
	return Tuple{
		Partition: b.Partition,
		Source:    OR(namespace, objId, relation, NodeType_RELATION_SOURCE),
		Dest:      OR(b.ActorNamespace, actorId, "", NodeType_ACTOR),
	}
}

// Return Tuple which grants ALL actors (in actor namespace) relation to objId
func (b *TupleBuilder) GrantAll(namespace, objId, relation string) Tuple {
	return Tuple{
		Partition: b.Partition,
		Source:    OR(namespace, objId, relation, NodeType_RELATION_SOURCE),
		Dest:      OR(b.ActorNamespace, "", "", NodeType_ALL_ACTORS),
	}
}

// Return Tuple which relates two system objects
func (b *TupleBuilder) Attribute(namespace, objId, relation, dstNamespace, dstObj string) Tuple {
	return Tuple{
		Partition: b.Partition,
		Source:    OR(namespace, objId, relation, NodeType_RELATION_SOURCE),
		Dest:      OR(dstNamespace, dstObj, "", NodeType_OBJECT),
	}
}

func (b *TupleBuilder) Actor(id string) TupleNode {
	return OR(b.ActorNamespace, id, "", NodeType_ACTOR)
}

func (b *TupleBuilder) ActorWithNamespace(namespace, id string) TupleNode {
	return OR(namespace, id, "", NodeType_ACTOR)
}

func (b *TupleBuilder) RelSource(namespace, id, relation string) TupleNode {
	return OR(namespace, id, relation, NodeType_RELATION_SOURCE)
}

func (b *TupleBuilder) Object(namespace, id string) TupleNode {
	return OR(namespace, id, "", NodeType_OBJECT)
}

func OR(ns, obj, rel string, t NodeType) TupleNode {
	return TupleNode{
		Namespace: ns,
		Id:        obj,
		Relation:  rel,
		Type:      t,
	}
}
