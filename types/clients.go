package types

// SimpleClient provides a public client that deals purely with relationships
type SimpleClient interface {
	GetAuthorizer() Authorizer
	GetRelationshipService() RelationshipService
	GetPolicyService() PolicyService
}

// RecordClient provides a public interface for records
type RecordClient[T any, PT ProtoConstraint[T]] interface {
	GetAuthorizer() Authorizer
	GetRecordService() RecordService[T, PT]
	GetPolicyService() PolicyService
}
