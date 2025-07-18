package relation_graph

import (
	"context"

	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/types"
)

type SearchResult struct {
	Authorized bool
	Completed  bool
	Explored   bool
}

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
