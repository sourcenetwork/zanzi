package types

// RecordFound is used to indicate the result of Set and Delete
// operations in Repositories.
// If RecordFound is true then it means the record was found and deleted
// or that it was found and updated
type RecordFound bool

// type PolicyService api.PolicyServiceServer
