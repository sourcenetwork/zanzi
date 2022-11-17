package tuples

import (
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/types/known/anypb"
)


// ObjRel represents a Zanzibar userset.
// In the stored graph, ObjRel represents a Node
type ObjRel struct {
	Namespace string 
	Id        string 
	Relation  string
}


func (o *ObjRel) ToRec() ObjRelRecord {
    return ObjRelRecord {
        Namespace: o.Namespace,
        Id: o.Id,
        Relation: o.Relation,
    }
}

func (o *ObjRelRecord) ToRel() ObjRel {
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
    Type TupleType
}

func (t *Tuple[T]) ToRec() TupleRecord {
    data, err := anypb.New(t.Data)
    if err != nil {
        panic(err)
    }

    return TupleRecord {
        ObjectRel: &t.ObjectRel.ToObjRelRec(),
        Actor: &t.Actor.ToObjRelRec(),
        Type: t.Type,
        Data: data,
    }
}

func (t *TupleRecord) ToTuple[T proto.Message]() Tuple[T] {
    var data T
    err := any.UnmarshalTo(&data)
    if err != nil {
        panic(err)
    }

    return Tuple[T] {
        ObjectRel: t.ObjectRel.ToObjRel(),
        Actor: t.Actor.ToObjRel(),
        Type: t.Type,
        Data: data,
    }
}
