package mappers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sourcenetwork/zanzi/internal/domain/policy"
	parser "github.com/sourcenetwork/zanzi/internal/permission_parser"
)

func TestParseTreeCuNode(t *testing.T) {
	expr := "relation"
	parseTree, _ := parser.Parse(expr)

	got := ToRewriteTree(parseTree)

	want := policy.CU("relation")

	assert.Equal(t, got, want)
}

func TestParseTreeTTUNode(t *testing.T) {
	expr := "relation->subrelation"
	parseTree, _ := parser.Parse(expr)

	got := ToRewriteTree(parseTree)

	want := policy.TTU("relation", "", "subrelation")

	assert.Equal(t, got, want)
}

func TestParseTreeThisNode(t *testing.T) {
	expr := "_this"
	parseTree, _ := parser.Parse(expr)

	got := ToRewriteTree(parseTree)

	want := policy.ThisTree()

	assert.Equal(t, got, want)
}

func TestParseTreeNestedExprssion(t *testing.T) {
	expr := "_this + (relation - from->to) & all"
	parseTree, _ := parser.Parse(expr)

	got := ToRewriteTree(parseTree)

	want := policy.Intersection(
		policy.Union(
			policy.ThisTree(),
			policy.Diff(
				policy.CU("relation"),
				policy.TTU("from", "", "to"),
			),
		),
		policy.CU("all"),
	)

	assert.Equal(t, got, want)
}
