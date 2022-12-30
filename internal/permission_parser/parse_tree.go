package permission_parser

type SetOperation int
const (
    Union SetOperation = iota
    Intersection
    Difference
)

type Error struct { }

// ThisNode represents a _this rule
type ThisNode struct {}

func (n *ThisNode) IsRule(){}
func (n *ThisNode) IsTerm(){}
func (n *ThisNode) IsPermExpr(){}


// CUNode represents a computed userset rule
type CUNode struct {
    Name string
}

func (n *CUNode) IsRule(){}
func (n *CUNode) IsTerm(){}
func (n *CUNode) IsPermExpr(){}


// TTUNode represents a tuple to userset rule
type TTUNode struct {
    Source string
    Target string
}

func (n *TTUNode) IsRule(){}
func (n *TTUNode) IsTerm(){}
func (n *TTUNode) IsPermExpr(){}


type Expression struct {
    Op SetOperation 
    Left Term
    Right Term
}

func (n *Expression) IsTerm(){}
func (n *Expression) IsPermExpr(){}


// rule represents a userset rewrite rule in the parse tree
// eg: a Computed Userset, Tuple To Userset or This rule.
type Rule interface {
    Term
    IsRule()
}


// Term represents a Rule or an Expression
type Term interface {
    PermExpr
    IsTerm()
}

// Term represents a permission expression
type PermExpr interface {
    IsPermExpr()
}
