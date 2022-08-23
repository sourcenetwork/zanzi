package querier

import (
	"context"
	"github.com/sourcenetwork/source-zanzibar/repository"
)

// Userset Rewrite rules specifiy a function which controls how a relation will return a userset during a query.
// Quoting the paper:
// "Each rule specifies a function that takes an object ID as input and outputs a userset expression tree."
//
// For efficiency reason, a rule will have a different algorithm based on the desired result.

// Recursively fetches tuples matching namespace, objectId and relation
func chaseUsersets(ctx context.Context, namespace, objectId, relation string) ([]model.TupleRecord, error) {
	// get repository from ctx
	var repo repository.Repository

	records, err := repo.GetTuplesFromUserset(namespace, objectId, relation)
	if err != nil {
		// TODO wrap err
		return nil, err
	}

	usersets := make(model.User, 0, len(records))
	for record := range records {
		user := record.Tuple.User
		if user.Type == model.UserType_USER_SET {
			append(usersets, user)
		}
	}

	for userset := range usersets {
		subRecords, err := chaseUsersets(ctx, userset.Namespace, userset.Identifier, userset.Relation)
		if err != nil {
			// TODO wrap err
			return nil, err
		}
		records = append(records, subRecords)
	}

	return records, nil
}

func (cu *ComputerUserset) Expand(ctx context.Context, namespace, objectId string) ([]model.User, error) {
	// computer userset should build an rewrite rule tree for cu.Relation
	// and perform an expand call on it
	//return chaseUsersets(ctx, namespace, objectId, t.Relation)
}

func (ttu *TupleToUserset) Expand(ctx context.Context, namespace, objectId string) ([]model.User, error) {
	// get repository from ctx
	var repo repository.Repository

	users, err := repo.GetUsersets(namespace, objectId, ttu.TuplesetRelation)
	if err != nil {
		// TODO wrap err
		return err
	}

	objs := make([]string, len(users))
	for user := range users {
		if user.Type == model.UserType_USER_SET {
			objs = append(objs, user.Identifier)
		}
	}

	// perform
	for obj := range objs {
		// now we gotta perform a This call for each resulting object?
		// no, I reckon we get the ComputedUsersetRelation tree
		// and Eval it passing these objects
		// collect the results and return
		//
	}
	return nil
}
