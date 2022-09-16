package builder

import (
	"github.com/sourcenetwork/source-zanzibar/model"
)

func buildOpNode(op model.Operation, left, right *model.RewriteNode) *model.RewriteNode {
	return &model.RewriteNode{
		Node: &model.RewriteNode_Opnode{
			Opnode: &model.OpNode{
				Op:    op,
				Left:  left,
				Right: right,
			},
		},
	}
}

func Union(left, right *model.RewriteNode) *model.RewriteNode {
	return buildOpNode(model.Operation_UNION, left, right)
}

func Interesection(left, right *model.RewriteNode) *model.RewriteNode {
	return buildOpNode(model.Operation_INTERSECTION, left, right)
}

func Diff(left, right *model.RewriteNode) *model.RewriteNode {
	return buildOpNode(model.Operation_DIFFERENCE, left, right)
}

func buildLeaf(rule *model.Rule) *model.RewriteNode {
	return &model.RewriteNode{
		Node: &model.RewriteNode_Leaf{
			Leaf: &model.Leaf{
				Rule: rule,
			},
		},
	}
}

func This() *model.RewriteNode {
	rule := &model.Rule{
		Rule: &model.Rule_This{
			This: &model.This{},
		},
	}
	return buildLeaf(rule)
}

func CU(relation string) *model.RewriteNode {
	rule := &model.Rule{
		Rule: &model.Rule_ComputedUserset{
			ComputedUserset: &model.ComputedUserset{
				Relation: relation,
			},
		},
	}
	return buildLeaf(rule)
}

func TTU(tuplesetRelation, computedUsersetRelation string) *model.RewriteNode {
	rule := &model.Rule{
		Rule: &model.Rule_TupleToUserset{
			TupleToUserset: &model.TupleToUserset{
				TuplesetRelation:        tuplesetRelation,
				ComputedUsersetRelation: computedUsersetRelation,
			},
		},
	}
	return buildLeaf(rule)
}

func ThisRelation(name string) *model.Relation {
    return Relation(name, This())
}

func Relation(name string, expTree *model.RewriteNode) *model.Relation {
    return &model.Relation {
        Name: name,
        Rewrite: &model.UsersetRewrite{
            ExpressionTree: expTree,
        },
    }
}


func Namespace(name string, relations ...*model.Relation) model.Namespace {
	return model.Namespace{
		Name:      name,
                Relations: relations,
	}
}
