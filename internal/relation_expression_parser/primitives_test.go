package relation_expression_parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newToken(t tokenType, lexeme string) token {
	return token{
		Type:   t,
		Lexeme: lexeme,
	}
}

func TestUnitParserConsumesOneToken(t *testing.T) {
	tokens := []token{
		newToken(tokenIdentifier, "first"),
		newToken(tokenIdentifier, "second"),
	}

	token, err, tail := unit(tokens)

	assert.Nil(t, err)
	assert.Equal(t, token, newToken(tokenIdentifier, "first"))
	assert.Equal(t, len(tail), 1)
	assert.Equal(t, tail[0], newToken(tokenIdentifier, "second"))
}

func TestUnitParserFailsWithEmptyTokenSlice(t *testing.T) {
	var tokens []token

	_, err, tail := unit(tokens)

	assert.NotNil(t, err)
	assert.Equal(t, len(tail), 0)
}

func TestManyGreedilyConsumesTokens(t *testing.T) {
	// parser to consume entire stream
	all := many(unit)

	tokens := []token{
		newToken(tokenIdentifier, "first"),
		newToken(tokenIdentifier, "second"),
		newToken(tokenThis, ""),
	}

	parsed, err, tail := all(tokens)

	assert.Nil(t, err)
	assert.Equal(t, parsed, tokens)
	assert.Equal(t, len(tail), 0)
}

func TestManyReturnsNothingOnEmptyTokenStream(t *testing.T) {
	// parser to consume entire stream
	all := many(unit)

	var tokens []token

	parsed, err, tail := all(tokens)

	assert.Nil(t, err)
	assert.Equal(t, len(parsed), 0)
	assert.Equal(t, len(tail), 0)
}

func TestParseIfAnyTypeReturnsErrorOnNoMatch(t *testing.T) {
	tokens := []token{
		newToken(tokenIdentifier, "first"),
	}

	parser := parseIfAnyType(tokenGroupBegin, tokenGroupEnd, tokenArrow)
	_, err, tail := parser(tokens)

	assert.NotNil(t, err)
	assert.Equal(t, tail, tokens)
}

func TestParseIfAnyTypeExitsOnFirstMatch(t *testing.T) {
	tokens := []token{
		newToken(tokenIdentifier, "first"),
	}

	parser := parseIfAnyType(tokenGroupBegin, tokenIdentifier, tokenArrow)
	id, err, tail := parser(tokens)

	assert.Nil(t, err)
	assert.Equal(t, len(tail), 0)
	assert.Equal(t, id, tokens[0])
}
