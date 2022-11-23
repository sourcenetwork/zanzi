package builder

struct Builder {
    id string
    name string
    attrs map[string]string
    resources []Resource
    actors []Actor
}

func (b *Builder) Name(name string) {
    b.name = name
}

func (b *Builder) Id(id string) {
    b.id = id
}

func (b *Builder) AddAttr(key, value string) {
    if b.attrs == nil {
        b.attrs = make(map[string]string
    }
    b.attrs[key] = value
}

func (b *Builder) Resources(resource ...Resource) {
}

func (b *Builder) Actors(actor ...Actor) {
}


func buildOpNode(op Operation, left, right *RewriteNode) *RewriteNode {
	return &RewriteNode{
		Node: &RewriteNode_Opnode{
			Opnode: &OpNode{
				Op:    op,
				Left:  left,
				Right: right,
			},
		},
	}
}

func Union(left, right *RewriteNode) *RewriteNode {
	return buildOpNode(Operation_UNION, left, right)
}

func Interesection(left, right *RewriteNode) *RewriteNode {
	return buildOpNode(Operation_INTERSECTION, left, right)
}

func Diff(left, right *RewriteNode) *RewriteNode {
	return buildOpNode(Operation_DIFFERENCE, left, right)
}

func buildLeaf(rule *Rule) *RewriteNode {
	return &RewriteNode{
		Node: &RewriteNode_Leaf{
			Leaf: &Leaf{
				Rule: rule,
			},
		},
	}
}

func This() *RewriteNode {
	rule := &Rule{
		Rule: &Rule_This{
			This: &This{},
		},
	}
	return buildLeaf(rule)
}

func CU(relation string) *RewriteNode {
	rule := &Rule{
		Rule: &Rule_ComputedUserset{
			ComputedUserset: &ComputedUserset{
				Relation: relation,
			},
		},
	}
	return buildLeaf(rule)
}

func TTU(tuplesetRelation, computedUsersetRelation string) *RewriteNode {
	rule := &Rule{
		Rule: &Rule_TupleToUserset{
			TupleToUserset: &TupleToUserset{
				TuplesetRelation:        tuplesetRelation,
				ComputedUsersetRelation: computedUsersetRelation,
			},
		},
	}
	return buildLeaf(rule)
}

func ThisRelation(name string) *Relation {
	return Relation(name, This())
}

func Relation(name string, expTree *RewriteNode) *Relation {
	return &Relation{
		Name: name,
		Rewrite: &UsersetRewrite{
			ExpressionTree: expTree,
		},
	}
}

func Namespace(name string, relations ...*Relation) Namespace {
	return Namespace{
		Name:      name,
		Relations: relations,
	}
}
