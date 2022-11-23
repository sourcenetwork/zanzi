package policy

import (
    "testing"
)

func TestMapPolicyGraph(t *testing.T) {
    policyGraphTestSuite(NewMapPolicyGraph, t)
}
