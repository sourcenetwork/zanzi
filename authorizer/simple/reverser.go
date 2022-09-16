package simple

import (
	"context"

	"github.com/sourcenetwork/source-zanzibar/authorizer"
	"github.com/sourcenetwork/source-zanzibar/graph"
	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/repository"
)

var _ authorizer.Reverser = (*reverser)(nil)

type reverser struct {
	checker      authorizer.Checker
	root         model.Userset
	checkResults map[model.KeyableUset]bool
	visitedNodes map[model.KeyableUset]struct{}
	nsRepo       repository.NamespaceRepository
	tupleRepo    repository.TupleRepository
}

func (l *reverser) ReverseLookup(ctx context.Context, user model.User) ([]model.Userset, error) {
	l.root = *user.Userset

	err := l.reverseLookup(ctx, *user.Userset)
	if err != nil {
		// wrap
		return nil, err
	}

	usets := make([]model.Userset, 0, len(l.visitedNodes))
	for key, _ := range l.visitedNodes {
		usets = append(usets, key.ToUset())
	}

	return usets, nil
}

// Perform DFS search over nodes which are reachable from `node`
// for the user specified during reverser construction
//
// Uses `checker` to validate whether potential neighbors are in fact reachable.
func (l *reverser) reverseLookup(ctx context.Context, uset model.Userset) error {
	l.visitedNodes[uset.ToKey()] = struct{}{}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	fetcher := graph.NewFetcher(l.nsRepo, l.tupleRepo)
	ancestors, err := fetcher.FetchAll(ctx, uset)
	if err != nil {
		return err
	}

	for _, uset := range ancestors {
		_, err := l.check(ctx, uset)
		if err != nil {
			return err
		}
	}

	for _, ancestor := range ancestors {
		key := ancestor.ToKey()
		if _, seen := l.visitedNodes[key]; seen {
			continue
		}

		if ok, _ := l.check(ctx, ancestor); !ok {
			continue
		}

		err := l.reverseLookup(ctx, ancestor)
		if err != nil {
			return err
		}
	}

	return nil
}

// check whether a node should be included in the final result set
func (l *reverser) check(ctx context.Context, node model.Userset) (bool, error) {
	key := node.ToKey()

	result, ok := l.checkResults[key]
	if ok {
		return result, nil
	}

	result, err := l.checker.Check(ctx, node, l.root)
	if err != nil {
		return false, err
	}

	l.checkResults[key] = result
	return result, nil
}

func ReverserFromChecker(nsRepo repository.NamespaceRepository, tupleRepo repository.TupleRepository, checker authorizer.Checker) authorizer.Reverser {
	return &reverser{
		checker:      checker,
		checkResults: make(map[model.KeyableUset]bool),
		visitedNodes: make(map[model.KeyableUset]struct{}),
		nsRepo:       nsRepo,
		tupleRepo:    tupleRepo,
	}
}
