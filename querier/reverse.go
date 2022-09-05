package querier

import (
	"context"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/utils"
)


func ReverseLookup(ctx context.Context, checker Checker, user model.User) ([]model.Userset, error) {
    uset := user.Userset
    lookuper := newLookuper(uset, checker)
    err := lookuper.reverseLookup(ctx, uset)
    if err != nil {
        // wrap
        return nil, err
    }

    return lookuper.getReachableUsets(), nil
}


type lookuper struct {
    user *model.Userset
    checker Checker
    checkResults map[model.KeyableUset]bool
    parents map[model.KeyableUset]struct{}
}

func newLookuper(user *model.Userset, checker Checker) lookuper {
    return lookuper {
        user: user,
        checker: checker,
        checkResults: make(map[model.KeyableUset]bool),
        parents: make(map[model.KeyableUset]struct{}),
    }
}

// Perform DFS search over nodes which are reachable from `node` 
// for the user specified during lookuper construction
//
// Uses `checker` to validate whether potential neighbors are in fact reachable.
func (l *lookuper) reverseLookup(ctx context.Context, node *model.Userset) (error) {
    select {
    case <-ctx.Done():
        return ctx.Err()
    default:
    }
    

    tentativeReferrers, err := l.produceReferrers(ctx, node)
    if err != nil {
        return err
    }

    referrers, err := l.validateReferrers(ctx, tentativeReferrers)
    if err != nil {
        return err
    }

    for _, ref := range referrers {
        key := ref.ToKey()
        _, ok := l.parents[key]
        if !ok {
            l.parents[key] = struct{}{}
            l.reverseLookup(ctx, &ref)
        }
    }

    return nil
}

func (l *lookuper) validateReferrers(ctx context.Context, referrers []model.Userset) ([]model.Userset, error) {
    validated := make([]model.Userset, 0, len(referrers))
    for _, referrer := range referrers {
        ok, err := l.check(ctx, referrer)
        if err != nil {
            return nil, err
        }
        if ok {
            validated = append(validated, referrer)
        }
    }
    return validated, nil
}

// how to make this explainable and efficient?
// in reality, for efficiency sake i could probably just keep a map of everything
// for the explain stuff i need to do some extra work
// is the explain implementation worth?
// definitely, it makes the lookup intelligeable
// i need to explain the same machinery though
// how to do that

func (l *lookuper) check(ctx context.Context, node model.Userset) (bool, error) {
    key := node.ToKey()

    result, ok := l.checkResults[key]
    if ok {
        return result, nil
    }

    result, err := l.checker.Check(ctx, node, *l.user)
    if err != nil {
        return false, err
    }

    l.checkResults[key] = result
    return result, nil
}

// return all potential referrers for uset
// returned nodes must be checked in order to assure that
// targetted user is allow to interact with it.
func (l *lookuper) produceReferrers(ctx context.Context, uset *model.Userset) ([]model.Userset, error) {
    _ = utils.GetTupleRepo(ctx)
    _ = utils.GetNamespaceRepo(ctx)


    // TODO
    return nil, nil
}


// Compile results from DFS search and return slice of all reachable nodes 
func (l *lookuper) getReachableUsets() []model.Userset {
    usets := make([]model.Userset, 0, len(l.parents))
    for k, _ := range l.parents {
        usets = append(usets, k.ToUset())
    }
    return usets
}

// Apply a CU rule backwards and return the original node
// Return node is not guaranteed to exist
func revertCU(ctx context.Context, uset *model.Userset, referer string) (*model.Userset) {
    return &model.Userset{
        Namespace: uset.Namespace,
        ObjectId: uset.ObjectId,
        Relation: referer,
    }
}

// Apply a TTU rule backwards and return the original node
// Return node is not guaranteed to exist
func revertTTU(ctx context.Context, uset model.Userset, referer string, tuplesetFilter string) (*model.Userset, error) {
    // TODO
    return nil, nil
}
