// Temp bufferw with prototyped code.
// Just keeping it for reference atm

// Recursively fetches tuples matching namespace, objectId and relation
func ChaseUsersets(repo Repository, namespace, objectId, relation string) ([]model.TupleRecord, error) {
    records, err := repo.GetTuplesFromObjRel(namespace, objectId, relation)
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

// Perform an inverse lookup of all tuples whose user match the given userset
// Walks up the tuple graph, following usersets
func ReverseChaseUserset(repo Repository, namespace, id, relation string) ([]model.TupleRecord, error) {
    records, err := repo.GetTuplesFromUserset(namespace, id, relation)
    if err != nil {
        // TODO wrap err
        return nil, err
    }

    usersets := make((string, string, string), 0, len(records))
    for record := range records {
        tuple := record.Tuple
        append(usersets, (tuple.Namespace, tuple.ObjectId, tuple.Relation))
    }

    for userset := range usersets {
        subRecords, err := ReverseChaseUserset(ctx, userset.Namespace, userset.Identifier, userset.Relation)
        if err != nil {
            // TODO wrap err
            return nil, err
        }
        records = append(records, subRecords)
    }

    return records, nil
}
