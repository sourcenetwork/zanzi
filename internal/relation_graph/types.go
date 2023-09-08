package relation_graph

import (
	"context"
	"encoding/json"

	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/types"
)

type unit struct{}

// NodeRepository
type NodeRepository interface {
	GetSucessors(ctx context.Context, policyId string, node *domain.EntitySetNode) ([]*domain.RelationNode, error)
	ListEdges(ctx context.Context, policyId string) ([]types.Pair[*domain.RelationNode, *domain.RelationNode], error)
}

type GoalTree interface {
	GetResult() SearchResult
	SetParent(GoalTree)
}

// ORNode is a branching node within the GoalTree
// The ORNode is satisfied if the Goal is found in *any*
// Paths
type ORNode struct {
	Parent GoalTree
	Result SearchResult
	Paths  []GoalTree
}

func (n *ORNode) GetResult() SearchResult { return n.Result }
func (n *ORNode) SetParent(p GoalTree)    { n.Parent = p }

// ANDNode is a branching node within the GoalTree
// The ANDNode is satisfied if the Goal is found in *all*
// Paths
type ANDNode struct {
	Parent GoalTree
	Result SearchResult
	Paths  []GoalTree
}

func (n *ANDNode) GetResult() SearchResult { return n.Result }
func (n *ANDNode) SetParent(p GoalTree)    { n.Parent = p }

// DifferenceNode is a branching node within the GoalTree
// A DifferenceNode is satisfied if the Goal is found
// on the Left GoalTree but not on the Right.
type DifferenceNode struct {
	Parent GoalTree
	Result SearchResult
	Left   GoalTree
	Right  GoalTree
}

func (n *DifferenceNode) GetResult() SearchResult { return n.Result }
func (n *DifferenceNode) SetParent(p GoalTree)    { n.Parent = p }

type SearchResult int

const (
	// Path hasn't been fully explored
	// should be the default value
	SearchResult_UNKNOWN SearchResult = iota

	// Path leads to Goal
	SearchResult_SUCCESS

	// Path leads to dead end
	SearchResult_FAILURE
)

func (r SearchResult) String() string {
	switch r {
	case SearchResult_UNKNOWN:
		return "UNKNOWN"
	case SearchResult_SUCCESS:
		return "SUCCESS"
	case SearchResult_FAILURE:
		return "FAILURE"
	default:
		return ""
	}
	return ""
}

func (r SearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// PathNode represents a possible search path in the GoalTree
type PathNode struct {
	Parent       GoalTree
	Result       SearchResult
	RelationNode *domain.RelationNode
	Path         GoalTree
	Reason       string
}

func (n *PathNode) GetResult() SearchResult { return n.Result }
func (n *PathNode) SetParent(p GoalTree)    { n.Parent = p }

// Goal represents the concrete target of a Goal Tree,
// that is the concrete node which is being searched for using the Goal Tree
type Goal struct {
	Target *domain.RelationNode
}

type GoalTreeSerializer interface {
	Serialize(GoalTree) (string, error)
}
