package types

wconst (
    "google.golang.org/protobuf/proto"
)

// Relation is a container type for any relation.
// Embeds client application data.
type Relation[T proto.Message] struct {
    Data T
    Type RelationType
    Object Id
    Relation string
    Subject Id
    SubjectRelation string
}
