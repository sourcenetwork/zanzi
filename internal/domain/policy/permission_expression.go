package policy

type predicate func(string) bool

type lexerState int

/*
// permission expression language:
// expr = term op expr | term
// op = union | diff | intersection
// term = subexpr | rule
// rule = cu | ttu | this
// this = "_this"
// cu = identifier
// ttu = identifier arrow identifier
// subexpr = groupBegin expr groupEnd
// identifier = alphanum*


type tokenType int
const (
    eof tokenType = iota
    identifier
    union
    difference
    intersection
    groupBegin
    groupEnd
    arrow
    this
    lexerError
)

type lexer struct {
    txt string
    start int
    pos int
    state lexerState
    q chan item
}

func (l *lexer) lexIdentifier() { }
func (l *lexer) lexUnion() { }
func (l *lexer) lexDifference() { }
func (l *lexer) lexIntersection() { }
func (l *lexer) lexGroupBegin() { }
func (l *lexer) lexGroupEnd() { }
func (l *lexer) lexArrow() { }
func (l *lexer) lexThis() { }

func (l *lexer) skipWS() { }

func (l *lexer) accept(pr predicate) { }
func (l *lexer) stepIf(pr predicate) {
}
func (l *lexer) peek() rune { }
func (l *lexer) nextRune() {

}
func (l *lexer) backstep() { }
func (l *lexer) ignore() { }

// emit token in q channel
func (l *lexer) emit(t tokenType) {
    data := l.txt[l.start:l.pos]
    i := item {
        Type: t,
        Data: data,
        StartPos: l.start,
        EndPos: l.pos,
    }
    l.q <- i
}


// steps through window

type item struct {
    Type tokenType
    Data string
    StartPos int
    EndPos int
}



type parser struct {

}

*/

// parser combinator
// expr = term + many(op expr) - fold many result into one thing
// op = try union, diff, intersection
// term = try rule, subexpr
// rule = try cu, ttu
// cu = identifier
// ttu = identifier + arrow + identifier
// subexpr = groupBegin + expr + groupEnd
// identifier = id token
