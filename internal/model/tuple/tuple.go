package tuples

import (
    "google.golang.org/protobuf/proto"
)


// ObjRel represents a Zanzibar userset.
// In the stored graph, ObjRel represents a Node
type ObjRel struct {
	Namespace string 
	Id        string 
	Relation  string
}


func (o *ObjRel) ToObjRelRec() ObjRelRecord {
    return ObjRelRecord {
        Namespace: o.Namespace,
        Id: o.Id,
        Relation: o.Relation,
    }
}

func (o *ObjRelRecord) ToObjRel() ObjRel {
    return ObjRel {
        Namespace: o.Namespace,
        Id: o.Id,
        Relation: o.Relation,
    }
}


// GenericTuple represent a tuple to be serialized with a type parameter.
// The type parameter allows users to embed custom application data
type Tuple[T proto.Message] struct {
    ObjectRel ObjRel
    Actor ObjRel
    Data T
}

func (t *Tuple[T]) ToRecord() TupleRecord {
    return TupleRecord {
        ObjectRel: t.ObjectRel.ToObjRelRec(),
        Actor: t.Actor.ToObjRelRec(),
        Data: T{} // TODO
    }
}

func (t *TupleRecord) ToRecord[T proto.Message]() Tuple {
    return Tuple[T] {
        ObjectRel: t.ObjectRel.ToObjRel(),
        Actor: t.Actor.ToObjRel(),
        Data: T{} // TODO
    }
}
