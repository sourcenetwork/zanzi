// Package rewrite defines types for Userset Rewrite Expression tree handling
package rewrite

import (
    "context"

    "github.com/sourcenetwork/source-zanzibar/repository"
    "github.com/sourcenetwork/source-zanzibar/model"
)

// Build Expression Tree which are used for userset expansion
func BuildExpressionTree(ctx context.Context, namespace string, relName string) (*Node, error) {
    // NOTE should I use a ctx or pass the Repository instance directly here?
    // Not sure which would be a better approach
    var repo repository.NamespaceRepository

    relation, err := repo.GetRelation(namespace, relName)
    if err != nil {
        // wrap err?
        return nil, err
    }

    tree := convertTree(relation.Rewrite.ExpressionTree, relation.Name)
    return &tree, nil
}

func convertTree(root *model.RewriteNode, relation string) Node {
    var node Node
    switch n := root.Node.(type) {
    case model.RewriteNode_Opnode:
        opnode := n.Opnode
        left := convertTree(n.Left, relation)
        right := convertTree(n.Left, relation)
        node = OpNode {
            Left: left,
            Right: right,
            Op: n.Op,
        }
    case model.Leaf:
        rule := convertLeaf(n, relation)
        node = RuleNode {
            Rule: rule,
        }
    }
    return node
}

// Maps a Userset Rewrite Rule from a Leaf into a rewerite Rule
// which is used in userset rewrite expansion
func convertLeaf(leaf model.Leaf, relation string) Rule {
    var rule Rule

    switch rule := leaf.Rule.GetRule().(type) {

    case model.This:
        // This is really just a shorthand for a Computed Userset
        // where the target relation is the one being defined.
        // Since during userset expansion we'd need to propagate
        // the current relation, it's easier to explicitly add 
        // the relation name to the Rule struct
        args := make([]string, 1)
        args[0] = relation
        rule = Rule{
            Type: RuleType_This,
            Args: args,
        }

    case model.TupleToUserset:
        args := make([]string, 2)
        args[0] = rule.TuplesetRelation
        args[1] = rule.ComputedUsersetRelation
        rule = Rule{
            Type: RuleType_TTU,
            Args: args,
        }

    case model.ComputedUserset:
        args := make([]string, 1)
        args[0] = rule.Relation
        rule = Rule{
            Type: RuleType_CU,
            Args: args,
        }
    }
    return rule
}

// TODO an Eager builder which builds the trees TTU and CU rules
