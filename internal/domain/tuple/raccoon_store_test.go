package tuple

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/cosmos/cosmos-sdk/store/mem"
    "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/internal/test_utils"
)


func TestTupleSetAndGet(t *testing.T) {
    kv := mem.NewStore()
    ts := NewRaccoonStore[*test_utils.Appdata](kv, nil)
    builder := TupleBuilder[*test_utils.Appdata]{}

    tuple := builder.Grant("file", "readme", "owner", "bob")
    err := ts.SetTuple(tuple)
    assert.Nil(t, err)

    opt, err := ts.GetTuple("", tuple.Source, tuple.Dest)
    assert.False(t, opt.IsEmpty())
    assert.Nil(t, err)
    got := opt.Value()
    assert.True(t, tuple.Equivalent(&got))
}

func TestTupleSetDelete(t *testing.T) {
    kv := mem.NewStore()
    ts := NewRaccoonStore[*test_utils.Appdata](kv, nil)
    builder := TupleBuilder[*test_utils.Appdata]{}

    tuple := builder.Grant("file", "abc", "owner", "bob")
    err := ts.SetTuple(tuple)
    assert.Nil(t, err)

    err = ts.DeleteTuple("", tuple.Source, tuple.Dest)
    assert.Nil(t, err)

    got, err := ts.GetTuple("", tuple.Source, tuple.Dest)
    assert.True(t, got.IsEmpty())
    assert.Nil(t, err)
}

func TestGetAncestors(t *testing.T) {
    kv := mem.NewStore()
    ts := NewRaccoonStore[*test_utils.Appdata](kv, nil)
    builder := TupleBuilder[*test_utils.Appdata]{}

    t1 := builder.Grant("file", "abc", "owner", "bob")
    err := ts.SetTuple(t1)
    assert.Nil(t, err)

    t2 := builder.Grant("file", "doc", "owner", "bob")
    err = ts.SetTuple(t2)
    assert.Nil(t, err)

    t3 := builder.Grant("file", "doc", "owner", "alice")
    err = ts.SetTuple(t3)
    assert.Nil(t, err)

    ancestors, err := ts.GetAncestors("", t2.Dest)
    assert.Nil(t, err)
    assert.True(t, contains(ancestors, &t1))
    assert.True(t, contains(ancestors, &t2))
}

func TestGetSucessors(t *testing.T) {
    kv := mem.NewStore()
    ts := NewRaccoonStore[*test_utils.Appdata](kv, nil)
    builder := TupleBuilder[*test_utils.Appdata]{}

    t1 := builder.Grant("file", "abc", "owner", "bob")
    err := ts.SetTuple(t1)
    assert.Nil(t, err)

    t2 := builder.Grant("file", "abc", "owner", "alice")
    err = ts.SetTuple(t2)
    assert.Nil(t, err)

    t3 := builder.Grant("file", "doc", "owner", "alice")
    err = ts.SetTuple(t3)
    assert.Nil(t, err)

    sucessors, err := ts.GetSucessors("", t1.Source)
    assert.Nil(t, err)
    assert.True(t, contains(sucessors, &t1))
    assert.True(t, contains(sucessors, &t2))
}

func contains[T proto.Message](ts []Tuple[T], t *Tuple[T]) bool {
    for _, elem := range ts {
        if t.Equivalent(&elem) {
            return true
        }
    }
    return false
}
