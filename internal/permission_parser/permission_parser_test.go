package permission_parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpressionParserSimple(t *testing.T) {
	got, err := Parse("a + b")

	var want Term = &Expression{
		&CUNode{"a"},
		Union,
		&CUNode{"b"},
	}

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestExpressionParserNested(t *testing.T) {
	got, err := Parse("a + (c - d) + b")

	var want Term = &Expression{
		&Expression{
			&CUNode{"a"},
			Union,
			&Expression{
				&CUNode{"c"},
				Difference,
				&CUNode{"d"},
			},
		},
		Union,
		&CUNode{"b"},
	}

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestExpressionParserDoublyNested(t *testing.T) {
	got, err := Parse("a + (_this - (f->g & d)) + b")

	var want Term = &Expression{
		&Expression{
			&CUNode{"a"},
			Union,
			&Expression{
				&ThisNode{},
				Difference,
				&Expression{
					&TTUNode{"f", "g"},
					Intersection,
					&CUNode{"d"},
				},
			},
		},
		Union,
		&CUNode{"b"},
	}

	assert.Nil(t, err)
	assert.Equal(t, got, want)
}

func TestRuleParserTTU(t *testing.T) {
	tokens := []token{
		newToken(tokenIdentifier, "from"),
		newToken(tokenArrow, ""),
		newToken(tokenIdentifier, "to"),
		newToken(tokenThis, ""),
	}

	rule, err, tail := ruleParser(tokens)
	ttu := rule.(*TTUNode)

	assert.Nil(t, err)
	assert.Equal(t, len(tail), 1)
	assert.Equal(t, ttu.Source, "from")
	assert.Equal(t, ttu.Target, "to")
}

func TestRuleParserCU(t *testing.T) {
	tokens := []token{
		newToken(tokenIdentifier, "relation"),
		newToken(tokenIdentifier, "to"),
	}

	rule, err, tail := ruleParser(tokens)
	cu := rule.(*CUNode)

	assert.Nil(t, err)
	assert.Equal(t, len(tail), 1)
	assert.Equal(t, cu.Name, "relation")
}

func TestRuleParserThis(t *testing.T) {
	tokens := []token{
		newToken(tokenThis, ""),
		newToken(tokenIdentifier, "to"),
	}

	rule, err, tail := ruleParser(tokens)
	_ = rule.(*ThisNode)

	assert.Nil(t, err)
	assert.Equal(t, len(tail), 1)
}
