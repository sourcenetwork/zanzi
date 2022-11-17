package types

import (
    "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/lib/types/tree"
    "github.com/sourcenetwork/source-zanzibar/lib/types"
)


type Authorizer interface {
    Check(id types.Id, relation string, actor types.Id) bool
    Reverse(actor types.Id) [](Id, string)
    Expand(id types.Id, relation string) tree.UsersetNode
    // additional useful methods like verbose reverse
}

type RelationService[T proto.Message] interface {
    Set(rel EntityRelation, data T) error
    Delete(rel Relation) error

    Grant(rel GrantRelation, data T) error
    Ungrant(rel GrantRelation) error

    Delegate(rel DelegationRelation, data T) error
    Undelegate(rel DelegationRelation) error

    GetRelations(obj ObjectId, relationName string) []Relation
}

type PolicyService interface {
    Set(policy Policy) error
    Get(policyId string) (types.Policy, error)
    Delete(policy Policy) error
}
