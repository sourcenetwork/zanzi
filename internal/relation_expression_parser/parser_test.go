package relation_expression_parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sourcenetwork/zanzi/pkg/domain"
)

func TestExpressionParserSimple(t *testing.T) {
	got, err := Parse("a + b")

	var want Term = domain.UnionNode(
		domain.CUNode("a"),
		domain.CUNode("b"),
	)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestExpressionParserNested(t *testing.T) {
	got, err := Parse("a + (c - d) + b")

	var want Term = domain.UnionNode(
		domain.UnionNode(
			domain.CUNode("a"),
			domain.DifferenceNode(
				domain.CUNode("c"),
				domain.CUNode("d"),
			),
		),
		domain.CUNode("b"),
	)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestExpressionParserDoublyNested(t *testing.T) {
	got, err := Parse("a + (_this - (f->g & d)) + b")

	var want Term = domain.UnionNode(
		domain.UnionNode(
			domain.CUNode("a"),
			domain.DifferenceNode(
				domain.ThisNode(),
				domain.IntersectionNode(
					domain.TTUNode("f", "g"),
					domain.CUNode("d"),
				),
			),
		),
		domain.CUNode("b"),
	)

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestRuleParserTTU(t *testing.T) {
	tokens := []token{
		newToken(tokenIdentifier, "from"),
		newToken(tokenArrow, ""),
		newToken(tokenIdentifier, "to"),
		newToken(tokenIdentifier, "tail"),
	}

	tree, err, tail := ruleParser(tokens)

	assert.Nil(t, err)
	assert.Equal(t, len(tail), 1)
	assert.Equal(t, tree, domain.TTUNode("from", "to"))
}

func TestRuleParserCU(t *testing.T) {
	tokens := []token{
		newToken(tokenIdentifier, "relation"),
		newToken(tokenIdentifier, "to"),
	}

	tree, err, tail := ruleParser(tokens)

	assert.Nil(t, err)
	assert.Equal(t, len(tail), 1)
	assert.Equal(t, tree, domain.CUNode("relation"))
}

func TestRuleParserThis(t *testing.T) {
	tokens := []token{
		newToken(tokenThis, ""),
		newToken(tokenIdentifier, "tail"),
	}

	tree, err, tail := ruleParser(tokens)

	assert.Nil(t, err)
	assert.Equal(t, len(tail), 1)
	assert.Equal(t, tree, domain.ThisNode())
}
