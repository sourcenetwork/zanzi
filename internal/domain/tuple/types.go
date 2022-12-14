package tuple

import (
    "time"

    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/types/known/anypb"

    opt "github.com/sourcenetwork/source-zanzibar/pkg/option"
)


// TupleNode represents a Zanzibar userset.
type TupleNode struct {
	Namespace string 
	Id        string 
	Relation  string
        Type NodeType
}


func (o *TupleNode) ToRec() TupleNodeRecord {
    return TupleNodeRecord {
        Namespace: o.Namespace,
        Id: o.Id,
        Relation: o.Relation,
        Type: o.Type,
    }
}

func (o *TupleNodeRecord) ToNode() TupleNode {
    return TupleNode {
        Namespace: o.Namespace,
        Id: o.Id,
        Relation: o.Relation,
        Type: o.Type,
    }
}


// Tuple represent a tuple to be serialized with a type parameter.
// The type parameter allows users to embed custom application data
type Tuple[T proto.Message] struct {
    Partition string
    CreatedAt time.Time
    Source TupleNode
    Dest TupleNode
    data T
    any *anypb.Any 
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

// Map tuple to a TupleRecord, serializes client data to protobuf's any
func (t *Tuple[T]) ToRec() TupleRecord {
    data, err := anypb.New(t.data)
    if err != nil {
        panic(err)
    }

    src := t.Source.ToRec()
    dst := t.Dest.ToRec()
    return TupleRecord {
        PartitionKey: t.Partition,
        Source: &src,
        Dest: &dst,
        ClientData: data,
    }
}

// Compares two tuples, verifies that the source, dest and partition are the same.
// Ignores client data in comparasion.
func (t *Tuple[T]) Equivalent(other *Tuple[T]) bool {
    return t.Partition == other.Partition && t.Source == other.Source && t.Dest == other.Dest
}


// Convert a TupleRecord to a Tuple
func toTuple[T proto.Message](rec *TupleRecord) Tuple[T] {
    return Tuple[T] {
        Partition: rec.PartitionKey,
        Source: rec.Source.ToNode(),
        Dest: rec.Dest.ToNode(),
        any: rec.ClientData,
    }
}


// TupleStore abstracts a backend storage service for tuples
type TupleStore[T proto.Message] interface {
    // Store a new tuple
    SetTuple(tuple Tuple[T]) error

    // Looks up a relation tuple from the backend storage
    GetTuple(partition string, source TupleNode, dest TupleNode) (opt.Option[Tuple[T]], error)

    // Purge tuple from storage.
    DeleteTuple(partition string, source TupleNode, dest TupleNode) error

    // Return all tuples whose source is node
    GetSucessors(partition string, node TupleNode) ([]Tuple[T], error)

    // Return all tuples whose dest is node
    GetAncestors(partition string, node TupleNode) ([]Tuple[T], error)
    
    // Return all tuples that grant `relation` to `objectId` contained in `objNamespace`
    // That is:
    // tuple.Source.Relation == relation
    // tuple.Dest.Id == objectId
    // tuple.Dest.Namespace == objNamespace
    //
    // (This peculiar query is used during reverse lookup)
    GetGrantingTuples(partition string, relation string, objNamespace string, objectId string) ([]Tuple[T], error)
}
