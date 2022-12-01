package tuple

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


// Tuple represent a tuple to be serialized with a type parameter.
// The type parameter allows users to embed custom application data
type Tuple[T proto.Message] struct {
    ObjectRel ObjRel
    Actor ObjRel
    data T
    any *anypb.Any 
    Type RelType
}

func (t *Tuple[T]) GetData() T {
    // FIXME
    if false {
        var data T
        err := t.any.UnmarshalTo(data)
        if err != nil {
            panic(err)
        }
        t.data = data
    }
    return t.data
}

func (t *Tuple[T]) SetData(data T) {
    t.data = data
}

func (t *Tuple[T]) ToRec() TupleRecord {
    data, err := anypb.New(t.data)
    if err != nil {
        panic(err)
    }

    objRel := t.ObjectRel.ToRec()
    actor := t.Actor.ToRec()
    return TupleRecord {
        ObjectRel: &objRel,
        Actor: &actor,
        Type: t.Type,
        ClientData: data,
    }
}

func (t *Tuple[T]) Equivalent(other *Tuple[T]) bool {
    return t.ObjectRel == other.ObjectRel && t.Actor == other.Actor && t.Type == other.Type
}

func toTuple[T proto.Message](rec *TupleRecord) Tuple[T] {
    return Tuple[T] {
        ObjectRel: rec.ObjectRel.ToRel(),
        Actor: rec.Actor.ToRel(),
        Type: rec.Type,
        any: rec.ClientData,
    }
}
