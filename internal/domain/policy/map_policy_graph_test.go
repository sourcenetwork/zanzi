package policy

import (
	"testing"
)

func TestMapPolicyGraph(t *testing.T) {
	s := buildTestSuite(NewMapPolicyGraph)
	s.Run(t)
}
