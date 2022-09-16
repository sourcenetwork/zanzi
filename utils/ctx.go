package utils

import (
	_ "context"
)

type ctxKey uint8

const (
	tupleRepoKey ctxKey = iota
	namespaceRepoKey
)
