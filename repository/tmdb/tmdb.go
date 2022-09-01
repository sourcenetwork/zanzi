// package tmdb implements repository interfaces using Tendermint's DB abstraction
package tmdb

import (
    _ "github.com/tendermint/tm-db"
)

// tentative paths to segregate keys
const (
    tuplePath      = "tuples/"
    namespacePath  = "namespaces/"
    usersetIdxPath = "usersets/"
)

// NOTE: Potential integration entypoint method in [github.com/cosmos/cosmos-sdk/store/types.CommitMultiStore.MountStoreWithDB]

// possible indexes:
//
// tuples:
// tuples/{namespace}/{obj_id}/{relation}/{namespace}/{id}/{relation} -> TupleRecord
// tuples/reversed/{user_namespace}/{id}/{relation}/{namespace}/{object_id}/{relation} -> TupleRecord  (reverse lookup)
//
// namespaces:
// namespaces/{namespace-name} -> Namespace
// namespaces/{namespace-name}/{relation-name} -> Relation
