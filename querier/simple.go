
package querier

import (
	"context"
        _ "log"

	"github.com/sourcenetwork/source-zanzibar/repository"
	"github.com/sourcenetwork/source-zanzibar/rewrite"
	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/utils"
)

type object struct {
    Namespace string
    ObjectId string
}

// NOTE The Expand assumes as an invariant that the namespace relation definition
// forms a DAG and does not contain back edges
// Namespaces must be validate in the service layer to guarantee relation graphs are well behaved
// Calling Expand with a cyclic directed graph will cause a stack overflow

func Expand(ctx context.Context, userset model.Userset) (rewrite.Node, error) {
    tree, err := rewrite.BuildExpressionTree(ctx, userset.Namespace, userset.Relation)
    if err != nil {
        // wrap
        return nil, err
    }

    obj := object {
        Namespace: userset.Namespace,
        ObjectId: userset.ObjectId,
    }

    return expandTree(ctx, tree, obj)
}

func expandTree(ctx context.Context, root rewrite.Node, obj object) (rewrite.Node, error) {
    // expandTree is the recusion entrypoint for Expand subcalls
    // Therefore it's a centralized point to check whether ctx is still valid before continueing
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
    }

    var result rewrite.Node
    var err error

    switch root.GetNodeType() {
    case rewrite.Node_OpNode:
        opnode := root.(*rewrite.OpNode)
        result, err = expandOpNode(ctx, *opnode, obj)

    case rewrite.Node_Leaf:
        result, err = root, nil

    case rewrite.Node_RuleNode:
        ruleNode := root.(*rewrite.RuleNode)
        result, err = expandRuleNode(ctx, *ruleNode, obj)
    }

    if err != nil {
        // wrap
        return nil, err
    }
    return result, nil
}

func expandOpNode(ctx context.Context, root rewrite.OpNode, obj object) (rewrite.Node, error) {
    left, err := expandTree(ctx, root.Left, obj)
    if err != nil {
        return nil, err
    }

    right, err := expandTree(ctx, root.Right, obj)
    if err != nil {
        return nil, err
    }

    node := &rewrite.OpNode {
        Left: left,
        Right: right,
        Op: root.Op,
    }

    return node, nil
}

func expandRuleNode(ctx context.Context, root rewrite.RuleNode, obj object) (rewrite.Node, error) {
    var child rewrite.Node
    var children []rewrite.Node
    var err error

    switch root.Rule.Type {
    case rewrite.RuleType_THIS:
        child, err = expandThis(ctx, root.Rule.Args[0], obj)
        children = append(children, child)
    case rewrite.RuleType_CU:
        child, err = expandCU(ctx, root.Rule.Args[0], obj)
        children = append(children, child)
    case rewrite.RuleType_TTU:
        children, err =  expandTTU(ctx, root.Rule.Args[0], root.Rule.Args[1], obj)
    }

    if err != nil {
        return nil, err
    }

    root = rewrite.RuleNode {
        Rule: root.Rule,
        Children: children,
    }
    return &root, nil
}

// Receive releation and object, builds userset and performs a bfs search on graph
// for all reachable nodes.
func expandThis(ctx context.Context, relation string, obj object) (rewrite.Node, error) {
        repo := utils.GetTupleRepo(ctx)
        chaser := repository.NewChaser(repo)
        
        uset := model.Userset {
            Namespace: obj.Namespace,
            ObjectId: obj.ObjectId,
            Relation: relation,
        }

        tuples, err := chaser.ChaseUsersets(uset)
        if err != nil {
            // wrap err
            return nil, err
        }

        usets := utils.MapSlice(tuples, tupleToUserset)

        leaf := &rewrite.Leaf {
            Users: usets,
        }

        return leaf, nil
}


// return a joinable node because the result might be a leaf or a opnode
func expandCU(ctx context.Context, relation string, obj object) (rewrite.Node, error) {
        tree, err := rewrite.BuildExpressionTree(ctx, obj.Namespace, relation)
        if err != nil {
            // wrap
            return nil, err
        }

        return expandTree(ctx, tree, obj)
}

func expandTTU(ctx context.Context, tsetRel string, cuRel string, obj object) ([]rewrite.Node, error) {
    repo := utils.GetTupleRepo(ctx)

    tuplesetFilter := model.Userset{
        Namespace: obj.Namespace,
        ObjectId: obj.ObjectId,
        Relation: tsetRel,
    }

    records, err := repo.GetRelatedUsersets(tuplesetFilter)
    if err != nil {
        // wrap
        return nil, err
    }

    tree, err := rewrite.BuildExpressionTree(ctx, obj.Namespace, cuRel)
    if err != nil {
        // wrap
        return nil, err
    }

    results := make([]rewrite.Node, 0, len(records))
    for _, record := range records {
        targetObj := object {
            Namespace: record.Tuple.ObjectRel.Namespace,
            ObjectId: record.Tuple.ObjectRel.Namespace,
        }

        result, err := expandTree(ctx, tree, targetObj)
        if err != nil {
            //wrap
            return nil, err
        }

        results = append(results, result)
    }

    return results, nil
}


func tupleToUserset(tuple model.TupleRecord) model.Userset {
    return model.Userset{
        Namespace: tuple.Tuple.User.Userset.Namespace,
        ObjectId: tuple.Tuple.User.Userset.ObjectId,
        Relation: tuple.Tuple.User.Userset.Relation,
    }
}
