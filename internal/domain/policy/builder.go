package policy

func NewActor(name string, types ...ActorIdType) *Actor {
	return &Actor{
		Name:        name,
		Constraints: types,
	}
}

func NewResource(name string, relations ...*Relation) *Resource {
	return &Resource{
		Name:      name,
		Relations: relations,
	}
}

func NewRelation(name string, tree *Tree, expr string) *Relation {
	return &Relation{
		Name:           name,
		ExpressionTree: tree,
		RewriteExpr:    expr,
	}
}

func ThisRelation(name string) *Relation {
	return NewRelation(name, ThisTree(), "_this")
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
		RewriteRule: &RewriteRule_This{
			This: &This{},
		},
	}
	return buildLeaf(rule)
}

func CU(relation string) *Tree {
	rule := &RewriteRule{
		RewriteRule: &RewriteRule_ComputedUserset{
			ComputedUserset: &ComputedUserset{
				Relation: relation,
			},
		},
	}
	return buildLeaf(rule)
}

func TTU(tuplesetRelation, cuRelNamespace, cuRel string) *Tree {
	rule := &RewriteRule{
		RewriteRule: &RewriteRule_TupleToUserset{
			TupleToUserset: &TupleToUserset{
				TuplesetRelation:    tuplesetRelation,
				CuRelationNamespace: cuRelNamespace,
				CuRelation:          cuRel,
			},
		},
	}
	return buildLeaf(rule)
}
