package policy

import "errors"

var ErrInvalidPolicy = errors.New("invalid policy")
var ErrInvalidRelationship = errors.New("invalid relationship")
var ErrSubjectRestriction = errors.New("subject restriction")
var ErrRelExpTree = errors.New("relation expression tree")
var ErrDuplicateDefinition = errors.New("duplicate definition")
var ErrSubjectNotAllowed = errors.New("subject not allowed")
var ErrPolicyNotFound = errors.New("policy not found")
var ErrResourceNotFound = errors.New("resource not found")
var ErrRelationshipNotFound = errors.New("relationship not found")
var ErrRelationNotFound = errors.New("relation not found")
var ErrPolicyExists = errors.New("policy exists")
