package querier

import (
	"context"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/tree"
)

func ReverseLookup(ctx context.Context, user model.User) (tree.UsersetNode, error) {
    // get userset from user
    // start

    // receive uset node
    // lookup direct neighbors
    // revert rewerite rules and check whether there are implicit neighbors
    // for each implicit neighbor, check whether the original user has access to that object [1], if it does keep it
    // for the final neighbors, call reverse lookup on them
    
    // is there a better way? this seems incredibly bad

    // [1] - at first i thought about checking the current node but that might not work?
    // my question is, is there a scenario where the subnode would be accessible but not the user?
    // suppose an extreme case like: we start of from a node whose relation is "excluded"
    // now we go up the tree to a dirrect neighbor with the relation "foo"
    // from "foo" we go to "bar" and bar has the rewrite rule (this Subtract CU "excluded")
    // well it depends because we gotta remember that what is being tested is the object, whatever pair
    // but for the sake of the argument assume it's ok.
    // then we might have a situation where uset (obj, "excluded") is not allowed to see obj with relation (obj, "foo")
    // due to the subtract business
    // whhich means, (obj, excluded) might be in THIS but then gets removed due to the Subtract op
    // meaning that it's not enough to call Check on the current userset
    // this is terrible, feels like we are essentially going to call expand on every reachable node in the graph
    // seems like such an expensive operation
    // could it be modeled some other way?

    // i'm toying with this idea, the idea of a reverse user expression tree
    return tree.UsersetNode{}, nil
}
