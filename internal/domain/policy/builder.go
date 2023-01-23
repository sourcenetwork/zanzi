package policy

func NewActor(name string, types ...ActorIdType) *Actor {
	return &Actor{
		Name:        name,
		Constraints: types,
	}
}

func BuildResource(name string, rules ...*Rule) *Resource {
	return &Resource{
		Name:  name,
		Rules: rules,
	}
}

func BuildRule(name string, t RuleType, tree *Tree, expr string) *Rule {
	return &Rule{
		Name:           name,
		Type:           t,
		ExpressionTree: tree,
		RewriteExpr:    expr,
	}
}

func BuildPerm(name string, tree *Tree, expr string) *Rule {
	return BuildRule(name, RuleType_PERMISSION, tree, expr)
}

func BuildRelation(name string, tree *Tree) *Rule {
	return BuildRule(name, RuleType_RELATION, tree, "_this")
}

func ThisRelation(name string) *Rule {
	return BuildRelation(name, ThisTree())
}

func BuildOpNode(op Operation, left, right *Tree) *Tree {
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
	return BuildOpNode(Operation_UNION, left, right)
}

func Intersection(left, right *Tree) *Tree {
	return BuildOpNode(Operation_INTERSECTION, left, right)
}

func Diff(left, right *Tree) *Tree {
	return BuildOpNode(Operation_DIFFERENCE, left, right)
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

func ThisTree() *Tree {
	rule := &RewriteRule{
		Rule: &RewriteRule_This{
			This: &This{},
		},
	}
	return buildLeaf(rule)
}

func CU(relation string) *Tree {
	rule := &RewriteRule{
		Rule: &RewriteRule_ComputedUserset{
			ComputedUserset: &ComputedUserset{
				Relation: relation,
			},
		},
	}
	return buildLeaf(rule)
}

func TTU(tuplesetRelation, cuRelNamespace, cuRel string) *Tree {
	rule := &RewriteRule{
		Rule: &RewriteRule_TupleToUserset{
			TupleToUserset: &TupleToUserset{
				TuplesetRelation:    tuplesetRelation,
				CuRelationNamespace: cuRelNamespace,
				CuRelation:          cuRel,
			},
		},
	}
	return buildLeaf(rule)
}
