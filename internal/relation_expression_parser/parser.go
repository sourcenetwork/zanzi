// Parser recognizes permission expression language.
// The accept language is:
// expr = term | term (op, term)+
// op = union | diff | intersection
// term = rule | subexpr
// rule = cu | ttu
// cu = identifier
// ttu = identifier + arrow + identifier
// subexpr = groupBegin + expr + groupEnd
package relation_expression_parser

import (
	"fmt"

	"github.com/sourcenetwork/zanzi/internal/utils"
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/types"
)

type Term = *domain.RelationExpressionTree

func Parse(expression string) (Term, error) {
	lexer, tokenChan := NewLexer(expression)
	go lexer.Lex()
	tokens := utils.ConsumeChan(tokenChan)

	if last := tokens[len(tokens)-1]; last.Type == tokenError {
		return nil, fmt.Errorf("error parsing expression %v: %v", expression, last.Lexeme)
	}

	expr, err, tail := expressionParser(tokens)
	if err != nil {
		return nil, err
	}
	_, _, tail = skipEof(tail)
	if len(tail) != 0 {
		// if the language is correct, a valid expression would consume all tokens
		// so maybe this is worthy of a panic
		return nil, fmt.Errorf("error parsing expression %v: leftover tokens, possibly misformed: tokens=%v", expression, tail)
	}

	return expr, nil
}

var identifierTokenParser parserFn[token] = parseIfType(tokenIdentifier)
var thisTokenParser parserFn[token] = parseIfType(tokenThis)
var arrowTokenParser parserFn[token] = parseIfType(tokenArrow)
var groupBeginParser parserFn[token] = parseIfType(tokenGroupBegin)
var groupEndParser parserFn[token] = parseIfType(tokenGroupEnd)
var setOpTokenParser parserFn[token] = parseIfAnyType(tokenUnion, tokenDifference, tokenIntersection)
var skipEof parserFn[[]token] = many(parseIfType(tokenEof))

// parser for Rule node
var ruleParser parserFn[Term] = first[Term]("Rule needs: This | ComputedUserset | TupleToUserset", thisParser, tupleToUsersetParser, computedUsersetParser)

// parser for Term node
var termParser parserFn[Term]

// parser for CUNode
func computedUsersetParser(tokens []token) (Term, error, []token) {
	id, err, tail := identifierTokenParser(tokens)
	if err != nil {
		return nil, err, tokens
	}

	return domain.CUNode(id.Lexeme), nil, tail
}

// parser for ThisNode
func thisParser(tokens []token) (Term, error, []token) {
	_, err, tail := thisTokenParser(tokens)
	if err != nil {
		return nil, err, tokens
	}

	return domain.ThisNode(), nil, tail
}

// parser for TTUNode
func tupleToUsersetParser(tokens []token) (Term, error, []token) {
	first, err, tail := identifierTokenParser(tokens)
	if err != nil {
		return nil, buildErr(tokens, "TupleToUserset needs: identifier arrow identifier"), tokens
	}

	_, err, tail = arrowTokenParser(tail)
	if err != nil {
		return nil, buildErr(tokens, "TupleToUserset needs: identifier arrow identifier"), tokens
	}

	second, err, tail := identifierTokenParser(tail)
	if err != nil {
		return nil, buildErr(tokens, "TupleToUserset needs: identifier arrow identifier"), tokens
	}

	return domain.TTUNode(first.Lexeme, second.Lexeme), nil, tail
}

// parses: tokenUnion | tokenIntersection | tokenDifference
func setOpParser(tokens []token) (domain.Operator, error, []token) {
	var op domain.Operator

	opToken, err, tokens := setOpTokenParser(tokens)
	if err != nil {
		return op, err, tokens
	}

	switch opToken.Type {
	case tokenUnion:
		op = domain.Operator_UNION
	case tokenIntersection:
		op = domain.Operator_INTERSECTION
	case tokenDifference:
		op = domain.Operator_DIFFERENCE
	default:
		panic("invalid token type")
	}
	return op, nil, tokens
}

// parses: setOperation + term
func opTermParser(tokens []token) (types.Pair[domain.Operator, Term], error, []token) {
	op, opErr, tail := setOpParser(tokens)
	if opErr != nil {
		return types.Pair[domain.Operator, Term]{}, buildErr(tokens, "opTerm needs: op + term"), tokens
	}

	term, termErr, tail := termParser(tail)
	if termErr != nil {
		return types.Pair[domain.Operator, Term]{}, buildErr(tokens, "opTerm needs: op + term"), tokens
	}

	pair := types.NewPair(op, term)
	return pair, nil, tail
}

var opTermsParser parserFn[[]types.Pair[domain.Operator, Term]] = many(opTermParser)

// parses: term + many(op, term)
func expressionParser(tokens []token) (Term, error, []token) {
	head, err, tail := termParser(tokens)
	if err != nil {
		return nil, buildErr(tokens, "expression needs: term"), tokens
	}

	pairs, err, tail := opTermsParser(tail)
	if len(pairs) == 0 {
		return head, nil, tail
	}

	var acc Term = head
	for _, pair := range pairs {
		op := pair.First()
		term := pair.Second()
		acc = domain.NewOpNode(acc, op, term)
	}
	return acc, nil, tail
}

// parses: groupBegin + expr + groupEnd
func subexprParser(tokens []token) (Term, error, []token) {
	_, openErr, tail := groupBeginParser(tokens)
	if openErr != nil {
		return nil, buildErr(tokens, "SubExpr needs: GroupBegin Expr GroupEnd"), tokens
	}

	expr, exprErr, tail := expressionParser(tail)
	if exprErr != nil {
		return nil, buildErr(tokens, "SubExpr needs: GroupBegin Expr GroupEnd"), tokens
	}

	_, endErr, tail := groupEndParser(tail)
	if endErr != nil {
		return nil, buildErr(tokens, "SubExpr needs: GroupBegin Expr GroupEnd"), tokens
	}

	return expr, nil, tail
}

func init() {
	// fixes initialization loop
	termParser = first[Term]("Term needs: Rule | SubExpr", ruleParser, subexprParser)
}
