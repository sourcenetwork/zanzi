package types

import (
    "google.golang.org/protobuf/proto"
)


type Authorizer interface {
    Check()
    Reverse()
    Expand()
    // additional useful methods like verbose reverse
}

type RelationService[T proto.Message] interface {
    Set(rel EntityRelation, data T)
    Delete(rel Relation)

    Grant(rel GrantRelation, data T)
    Ungrant(rel GrantRelation)

    Delegate(rel DelegationRelation, data T)
    Undelegate(rel DelegationRelation)

    GetRelations(obj ObjectId, relationName string) []Relation
}

type PolicyService interface {
    Set(policy Policy)
    Get(policyId string)
    Delete(policy Policy)
}
