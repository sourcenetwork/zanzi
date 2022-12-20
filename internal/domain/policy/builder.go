package policy

func BuildActor(name string, types ...ActorIdType) *Actor {
	return &Actor{
		Name: name,
	}
}

func BuildResource(name string, rules ...*Rule) *Resource {
	return &Resource{
		Name:  name,
		Rules: rules,
	}
}

func BuildRule(name string, t RuleType, tree *Tree) *Rule {
	return &Rule{
		Name:           name,
		Type:           t,
		ExpressionTree: tree,
	}
}

func BuildPerm(name string, tree *Tree) *Rule {
	return BuildRule(name, RuleType_PERMISSION, tree)
}

func BuildRelation(name string, tree *Tree) *Rule {
	return BuildRule(name, RuleType_RELATION, tree)
}

func ThisRelation(name string) *Rule {
	return BuildRelation(name, _This())
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

func _This() *Tree {
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
