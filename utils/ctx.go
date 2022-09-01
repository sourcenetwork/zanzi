package utils

import (
	"context"

	"github.com/sourcenetwork/source-zanzibar/repository"
)

type ctxKey uint8

const (
	tupleRepoKey ctxKey = iota
	namespaceRepoKey
)

func WithTupleRepository(ctx context.Context, repo repository.TupleRepository) context.Context {
	return context.WithValue(ctx, tupleRepoKey, repo)
}

func GetTupleRepo(ctx context.Context) repository.TupleRepository {
	repo := ctx.Value(tupleRepoKey).(repository.TupleRepository)
	return repo
}

func WithNamespaceRepository(ctx context.Context, repo repository.NamespaceRepository) context.Context {
	return context.WithValue(ctx, namespaceRepoKey, repo)
}

func GetNamespaceRepo(ctx context.Context) repository.NamespaceRepository {
	repo := ctx.Value(namespaceRepoKey).(repository.NamespaceRepository)
	return repo
}
