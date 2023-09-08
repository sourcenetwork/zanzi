package relation_expression_parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sourcenetwork/zanzi/internal/utils"
)

func TestTokens(t *testing.T) {
	l, q := NewLexer("read + write - admin & (res->root)")

	go l.Lex()

	tokens := utils.ConsumeChan(q)
	tokenEq(t, tokens[0], token{tokenIdentifier, "read", 0, 0})
	tokenEq(t, tokens[1], token{tokenUnion, "+", 0, 0})
	tokenEq(t, tokens[2], token{tokenIdentifier, "write", 0, 0})
	tokenEq(t, tokens[3], token{tokenDifference, "-", 0, 0})
	tokenEq(t, tokens[4], token{tokenIdentifier, "admin", 0, 0})
	tokenEq(t, tokens[5], token{tokenIntersection, "&", 0, 0})
	tokenEq(t, tokens[6], token{tokenGroupBegin, "(", 0, 0})
	tokenEq(t, tokens[7], token{tokenIdentifier, "res", 0, 0})
	tokenEq(t, tokens[8], token{tokenArrow, "->", 0, 0})
	tokenEq(t, tokens[9], token{tokenIdentifier, "root", 0, 0})
	tokenEq(t, tokens[10], token{tokenGroupEnd, ")", 0, 0})
	tokenEq(t, tokens[11], token{tokenEof, "", 0, 0})
	assert.Equal(t, len(tokens), 12)
}

func tokenEq(t *testing.T, got, expected token) {
	assert.Equal(t, got.Type, expected.Type)
	assert.Equal(t, got.Lexeme, expected.Lexeme)
}
