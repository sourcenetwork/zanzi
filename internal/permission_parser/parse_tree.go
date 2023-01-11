package permission_parser

type SetOperation int

const (
	Union SetOperation = iota
	Intersection
	Difference
)

type Error struct{}

// ThisNode represents a _this rule
type ThisNode struct{}

func (n *ThisNode) IsRule()     {}
func (n *ThisNode) IsTerm()     {}
func (n *ThisNode) IsPermExpr() {}

// CUNode represents a computed userset rule
type CUNode struct {
	Name string
}

func (n *CUNode) IsRule() {}
func (n *CUNode) IsTerm() {}

// TTUNode represents a tuple to userset rule
type TTUNode struct {
	Source string
	Target string
}

func (n *TTUNode) IsRule() {}
func (n *TTUNode) IsTerm() {}

type Expression struct {
	Left  Term
	Op    SetOperation
	Right Term
}

func (n *Expression) IsTerm() {}

// rule represents a userset rewrite rule in the parse tree
// eg: a Computed Userset, Tuple To Userset or This rule.
type Rule interface {
	Term
	IsRule()
}

// Term represents a Rule or an Expression
type Term interface {
	IsTerm()
}
