package policy

func BuildActor(name string, types ...ActorIdType) Actor {
    return Acotr
}

func BuildResource(name string, rules ...*Rule) {
    // todo
}

func BuildRule(name string, t RuleType, tree ExpressionTree) Rule {
    return Rule{
        Name: name,
        Type: t,
        ExpressionTree: tree,
    }
}

func BuilPerm(name string, tree ExpressionTree) Rule {
    return BuildRule(name, RuleType_PERMISSION, tree)
}

func BuildRelation(name string, tree ExpressionTree) Rule {
    return BuildRule(name, RuleType_RELATION, tree)
}

func ThisRelation(name string) Rule {
	return BuildRelation(name, This())
}


func buildOpNode(op Operation, left, right *Tree) *Tree {
	return &Tree{
		Node: &Tree_Opnode{
			Opnode: &OpNode{
				Op:    op,
				Left:  left,
				Right: right,
			},
		},
	}
}

func Union(left, right *Tree) *Tree {
	return buildOpNode(Operation_UNION, left, right)
}

func Interesection(left, right *Tree) *Tree {
	return buildOpNode(Operation_INTERSECTION, left, right)
}

func Diff(left, right *Tree) *Tree {
	return buildOpNode(Operation_DIFFERENCE, left, right)
}

func buildLeaf(rule *RewriteRule) *Tree {
	return &Tree{
		Node: &Tree_Leaf{
			Leaf: &Leaf{
				Rule: rule,
			},
		},
	}
}

func This() *Tree {
	rule := &RewriteRule{
		Rule: &RuleRewriteRule_This{
			This: &This{},
		},
	}
	return buildLeaf(rule)
}

func CU(relation string) *Tree {
	rule := &RewriteRule{
		Rule: &RuleRewriteRule_ComputedUserset{
			ComputedUserset: &ComputedUserset{
				Relation: relation,
			},
		},
	}
	return buildLeaf(rule)
}

func TTU(tuplesetRelation, computedUsersetRelation string) *Tree {
	rule := &RewriteRule{
		Rule: &RuleRewriteRule_TupleToUserset{
			TupleToUserset: &TupleToUserset{
				TuplesetRelation:        tuplesetRelation,
				ComputedUsersetRelation: computedUsersetRelation,
			},
		},
	}
	return buildLeaf(rule)
}

