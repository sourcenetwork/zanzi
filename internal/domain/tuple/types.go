package tuple

import (
    "time"

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
type Tuple struct {
    Partition string
    CreatedAt time.Time
    Source TupleNode
    Dest TupleNode
}

// Map tuple to a TupleRecord, serializes client data to protobuf's any
func (t *Tuple) ToRec() TupleRecord {
    src := t.Source.ToRec()
    dst := t.Dest.ToRec()
    return TupleRecord {
        PartitionKey: t.Partition,
        Source: &src,
        Dest: &dst,
    }
}

// Compares two tuples, verifies that the source, dest and partition are the same.
// Ignores client data in comparasion.
func (t *Tuple) Equivalent(other *Tuple) bool {
    return t.Partition == other.Partition && t.Source == other.Source && t.Dest == other.Dest
}


// Convert a TupleRecord to a Tuple
func toTuple(rec *TupleRecord) Tuple {
    return Tuple {
        Partition: rec.PartitionKey,
        Source: rec.Source.ToNode(),
        Dest: rec.Dest.ToNode(),
    }
}


// TupleStore abstracts a backend storage service for tuples
type TupleStore interface {
    // Store a new tuple
    SetTuple(tuple Tuple) error

    // Looks up a relation tuple from the backend storage
    GetTuple(partition string, source TupleNode, dest TupleNode) (opt.Option[Tuple], error)

    // Purge tuple from storage.
    DeleteTuple(partition string, source TupleNode, dest TupleNode) error

    // Return all tuples whose source is node
    GetSucessors(partition string, node TupleNode) ([]Tuple, error)

    // Return all tuples whose dest is node
    GetAncestors(partition string, node TupleNode) ([]Tuple, error)
    
    // Return all tuples that grant `relation` to `objectId` contained in `objNamespace`
    // That is:
    // tuple.Source.Relation == relation
    // tuple.Dest.Id == objectId
    // tuple.Dest.Namespace == objNamespace
    //
    // (This peculiar query is used during reverse lookup)
    GetGrantingTuples(partition string, relation string, objNamespace string, objectId string) ([]Tuple, error)
}
