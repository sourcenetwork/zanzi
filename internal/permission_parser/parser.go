package permission_parser


// Parser recognizes permission expression language.
// The accept language is:
// expr = term | term (op, term)+
// op = union | diff | intersection
// term = rule | subexpr
// rule = cu | ttu
// cu = identifier
// ttu = identifier + arrow + identifier
// subexpr = groupBegin + expr + groupEnd
type Parser struct { }

// Consume token slice and return parsed result
func (p *parser) Parse(tokens []token) (Term, error) {
    return expression(tokens)
}

// parser function type
// receives slice of tokens, return some type and another token slice
type parserFn[T any] func([]token) (T, []token)


// primitive parsers

// try all parsers until it receives a result
func tryParsers(tokens []token, parsers ...parserFn) (T, []token) {
    var result T
    for _, fn := range parsers {
        result, newTokens := fns(tokens)
        if result != nil {
            return result, newTokens
        }
    }
    return nil, tokens
}

// return next token if it is of given type
func nextIfType(tokens []token, ttype tokenType) (*token, []token) {
    if len(tokens) == 0 || tokens[0].Type != ttype {
        return nil, tokens
    } 

    return &tokens[0], tokens[1:]
}

// return next token if its type is in the given type list
func nextIfAnyType(tokens []token, types ...tokenType) (*token, []token) {
    if len(tokens) == 0 || {
        return nil, tokens
    } 

    for _, tokenType := range types {
        if token[0].Type == tokenType {
            return &tokens[0], tokens[1:]
        }
    }
    return nil, tokens
}

func many[T](tokens []token, parseT parserFn[T]) ([]T, []token) {
    var sliced []token = tokens
    var ts []T
    for {
        t, sliced := parseT(sliced)
        if t == nil {
            break
        }
        ts = append(ts, t)
    }
    return ts, sliced
}

func anyNil(elems ...any) bool {
    for _, elem := range elems {
        if elem == nil {
            return true
        }
    }
    return false
}


// combined parsers

// parses: identifier
func  computedUserset(tokens []token) (*CUNode, []token) {
    id, tokens := nextIfType(tokens, tokenIdentifier)
    if id == nil {
        return nil, tokens
    }

    return &RelationName{ id.Lexeme }, tokens
}

// parses: this
func  this(tokens []token) (*ThisNode, []tokens) {
    this, tokens := nextIfType(tokens, tokenThis)
    if this == nil {
        return nil, tokens
    }
    return &ThisNode{}, tokens
}

// parses: identifier + arrow + identifier
func  tupleToUserset(tokens []token) (*TTUNode, []tokens) {
    var sliced []token = tokens
    first, sliced := nextIfType(tokens, tokenIdentifier)
    arrow, sliced := nextIfType(tokens, tokenArrow)
    second, sliced := nextIfType(tokens, tokenIdentifier)

    if anyNil(first, arrow, second) {
        return nil, tokens
    }

    return &TraverseNode{ first, second }, sliced
}

// parses: this | cu | ttu
func  rule(tokens []token) (Rule, []tokens) {
    return tryParsers[Rule](tokens, this, computedUserset, tupleToUserset)
}

// parses: tokenUnion | tokenIntersection | tokenDifference
func  setOp(tokens []token) (*SetOperation, []tokens) {
    token, tokens := nextIfAnyType(tokens, tokenUnion, tokenDifference, tokenIntersection)
    if token == nil {
        return nil, tokens
    }

    var op *Operation
    switch token.Type {
    case tokenUnion:
        op = &Union
    case tokenIntersection:
        op = &Intersection
    case tokenDifference:
        op = &Difference
    }
    return op, tokens
}

// parses: setOperation + term
func opTerm(tokens []token) (*Pair[*SetOperation, Term], []token) {
    op, sliced := setOp(tokens)
    term, sliced := term(sliced)
    if anyNil(op, factor) {
        return nil, tokens
    }

    return NewPair(op, factor), sliced
}

// parses: term + many(op, term)
func  expression(tokens []token) (PermExpr, []tokens, error) {
    head, sliced := term(tokens)
    if head == nil {
        return nil, sliced
    }

    pairs, sliced := many(opTerm, sliced)
    if pairs == nil {
        return factor, sliced
    }

    var acc Term = head
    for _, pair := range pairs {
        op := pair.First()
        term := pair.Second()
        acc = &Expression {
            Left: acc,
            Op: op,
            Right: term,
        }
    }
    return acc, sliced
}

// parses: rule | subexpr
func term(tokens []token) (Term, []tokens) {
    return tryParsers(tokens, rule, subexpr)
}

// parses: groupBegin + expr + groupEnd
func subexpr(tokens []token) (PermExpr, []tokens) {
    var sliced []token = tokens
    begin, sliced := nextIfType(sliced, tokenGroupBegin)
    expr, sliced := expression(sliced)
    end, sliced := nextIfType(sliced, tokenGroupEnd)
    if anyNil(begin, end, expr) {
        return nil, tokens
    }
    return expr, sliced
}
