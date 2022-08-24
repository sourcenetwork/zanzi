
package querier

/*
import (
	"context"
	"github.com/sourcenetwork/source-zanzibar/repository"
	"github.com/sourcenetwork/source-zanzibar/rewrite"
)

// Userset Rewrite rules specifiy a function which controls how a relation will return a userset during a query.
// Quoting the paper:
// "Each rule specifies a function that takes an object ID as input and outputs a userset expression tree."
//
// For efficiency reason, a rule will have a different algorithm based on the desired result.

type object {
    Namespace string
    ObjectId string
}

struct Simple {
}

func (s *Simple) Expand(ctx context.Context, userset model.Userset) (rewrite.Node, error) {
    // computer userset should build an rewrite rule tree for cu.Relation
    // and perform an expand call on it
    //return chaseUsersets(ctx, namespace, objectId, t.Relation)

    // get relation tree
    // build rewrite tree
    // peform BFS expand calls for nodes in tree
    // each node type should have a func
    // return final tree pog
    // 
    // sub expand calls should take in a tree
}

func (s *Simple) expandTree(root *rewrite.Node, obj object) (*rewrite.Node, error) {
    switch root.GetNodeType() {

    case Node_OpNode:
        opnode := root.(rewrite.OpNode)
        return s.expandOpNode(opnode, object)

    case Node_Leaf:
        return root, nil

    case Node_RuleNode:
        // call dispatchrulenode
    }
}

func (s *Simple) expandOpNode(root *rewrite.OpNode, obj object) (*rewrite.Node, error) {
    left, err := expandTree(root.Left, obj)
    if err != nil {
        return nil, err
    }

    right, err := expandTree(root.Right, obj)
    if err != nil {
        return nil, err
    }

    node := &rewrite.OpNode {
        Left: left,
        Right: right,
        Op: root.Op
    }

    return node, nil
}

func (s *Simple) expandRuleNode(root *rewrite.RuleJoinNode) (*rewrite.Node, error) {
    switch root.Rule.RuleType {
    case rewrite.RuleType_This:
        return s.expandThis()
    case model.RuleType_TTU:
        return s.expandTTU()
    case model.RuleType_CU:
        return s.expandCU()
    }
}

func (s *Simple) expandThis(node *rewrite.RuleJoinNode, string relation, obj object) (*rewrite.RuleJoinNode, error) { 
	// get repository from ctx
	var repo repository.TupleRepository
        // call method to recursively fetch tuples
        // map tuples to list of usersets
        // make leaf node
        // set leaf node in  rule join node child
        // return node
}

func (s *Simple) expandCU(node rewrite.RuleJoinNode, namespace, objectId string) (rewrite.Node, error) {
	// get repository from ctx
	var repo repository.TupleRepository
        // build tree for the target relation
        // return call to expand tree
}
func (s *Simple) expandTTU() (rewrite.RuleJoinNode, error) {
    // get repository from ctx
    var repo repository.Repository

    // fetch relation tree
    // do ttu processing
    // get tuple nodes
    // call expand tree
    // create new nodes to merge trees
    // set merged nodes in rule node child
    // return rule node

    users, err := repo.GetUsersets(namespace, objectId, ttu.TuplesetRelation)
    if err != nil {
        // TODO wrap err
        return err
    }

    objs := make([]string, len(users))
    for user := range users {
        if user.Type == model.UserType_USER_SET {
            objs = append(objs, user.Identifier)
        }
    }

    // perform
    for obj := range objs {
        // now we gotta perform a This call for each resulting object?
        // no, I reckon we get the ComputedUsersetRelation tree
        // and Eval it passing these objects
        // collect the results and return
        //
    }
    return nil
}
*/
