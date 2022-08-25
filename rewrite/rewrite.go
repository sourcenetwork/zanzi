// Package rewrite defines types for Userset Rewrite Expression tree handling
package rewrite

import (
    "context"

    "github.com/sourcenetwork/source-zanzibar/model"
    "github.com/sourcenetwork/source-zanzibar/utils"
)

// Build Expression Tree which are used for userset expansion
func BuildExpressionTree(ctx context.Context, namespace string, relName string) (Node, error) {
    repo := utils.GetNamespaceRepo(ctx)

    relation, err := repo.GetRelation(namespace, relName)
    if err != nil {
        // wrap err?
        return nil, err
    }

    tree := convertTree(relation.Rewrite.ExpressionTree, relation.Name)
    return tree, nil
}

func convertTree(root *model.RewriteNode, relation string) Node {
    var node Node
    switch n := root.Node.(type) {
    case *model.RewriteNode_Opnode:
        opnode := n.Opnode
        left := convertTree(opnode.Left, relation)
        right := convertTree(opnode.Right, relation)
        node = &OpNode {
            Left: left,
            Right: right,
            Op: opnode.Op,
        }
    case *model.RewriteNode_Leaf:
        leaf := n.Leaf
        rule := convertLeaf(leaf, relation)
        node = &RuleNode {
            Rule: rule,
        }
    }
    return node
}

// Maps a Userset Rewrite Rule from a Leaf into a rewerite Rule
// which is used in userset rewrite expansion
func convertLeaf(leaf *model.Leaf, relation string) Rule {
    var rule Rule

    switch r := leaf.Rule.GetRule().(type) {

    case *model.Rule_This:
        // This is really just a shorthand for a Computed Userset
        // where the target relation is the one being defined.
        // Since during userset expansion we'd need to propagate
        // the current relation, it's easier to explicitly add 
        // the relation name to the Rule struct
        args := make([]string, 1)
        args[0] = relation
        rule = Rule{
            Type: RuleType_THIS,
            Args: args,
        }

    case *model.Rule_TupleToUserset:
        ttu := r.TupleToUserset
        args := make([]string, 2)
        args[0] = ttu.TuplesetRelation
        args[1] = ttu.ComputedUsersetRelation
        rule = Rule{
            Type: RuleType_TTU,
            Args: args,
        }

    case *model.Rule_ComputedUserset:
        cu := r.ComputedUserset
        args := make([]string, 1)
        args[0] = cu.Relation
        rule = Rule{
            Type: RuleType_CU,
            Args: args,
        }
    }
    return rule
}

// TODO an Eager builder which builds the trees TTU and CU rules
