// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: zanzi/domain/relation_graph.proto

package domain

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on RelationNode with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RelationNode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RelationNode with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RelationNodeMultiError, or
// nil if none found.
func (m *RelationNode) ValidateAll() error {
	return m.validate(true)
}

func (m *RelationNode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Node.(type) {
	case *RelationNode_EntitySet:
		if v == nil {
			err := RelationNodeValidationError{
				field:  "Node",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetEntitySet()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RelationNodeValidationError{
						field:  "EntitySet",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RelationNodeValidationError{
						field:  "EntitySet",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEntitySet()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RelationNodeValidationError{
					field:  "EntitySet",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *RelationNode_Entity:
		if v == nil {
			err := RelationNodeValidationError{
				field:  "Node",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetEntity()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RelationNodeValidationError{
						field:  "Entity",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RelationNodeValidationError{
						field:  "Entity",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEntity()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RelationNodeValidationError{
					field:  "Entity",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *RelationNode_Wildcard:
		if v == nil {
			err := RelationNodeValidationError{
				field:  "Node",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetWildcard()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RelationNodeValidationError{
						field:  "Wildcard",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RelationNodeValidationError{
						field:  "Wildcard",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetWildcard()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RelationNodeValidationError{
					field:  "Wildcard",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return RelationNodeMultiError(errors)
	}

	return nil
}

// RelationNodeMultiError is an error wrapping multiple validation errors
// returned by RelationNode.ValidateAll() if the designated constraints aren't met.
type RelationNodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RelationNodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RelationNodeMultiError) AllErrors() []error { return m }

// RelationNodeValidationError is the validation error returned by
// RelationNode.Validate if the designated constraints aren't met.
type RelationNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationNodeValidationError) ErrorName() string { return "RelationNodeValidationError" }

// Error satisfies the builtin error interface
func (e RelationNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationNodeValidationError{}

// Validate checks the field values on EntitySetNode with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EntitySetNode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EntitySetNode with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EntitySetNodeMultiError, or
// nil if none found.
func (m *EntitySetNode) ValidateAll() error {
	return m.validate(true)
}

func (m *EntitySetNode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetObject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntitySetNodeValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntitySetNodeValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntitySetNodeValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Relation

	if len(errors) > 0 {
		return EntitySetNodeMultiError(errors)
	}

	return nil
}

// EntitySetNodeMultiError is an error wrapping multiple validation errors
// returned by EntitySetNode.ValidateAll() if the designated constraints
// aren't met.
type EntitySetNodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EntitySetNodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EntitySetNodeMultiError) AllErrors() []error { return m }

// EntitySetNodeValidationError is the validation error returned by
// EntitySetNode.Validate if the designated constraints aren't met.
type EntitySetNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntitySetNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntitySetNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntitySetNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntitySetNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntitySetNodeValidationError) ErrorName() string { return "EntitySetNodeValidationError" }

// Error satisfies the builtin error interface
func (e EntitySetNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntitySetNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntitySetNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntitySetNodeValidationError{}

// Validate checks the field values on EntityNode with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EntityNode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EntityNode with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EntityNodeMultiError, or
// nil if none found.
func (m *EntityNode) ValidateAll() error {
	return m.validate(true)
}

func (m *EntityNode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetObject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntityNodeValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntityNodeValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntityNodeValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return EntityNodeMultiError(errors)
	}

	return nil
}

// EntityNodeMultiError is an error wrapping multiple validation errors
// returned by EntityNode.ValidateAll() if the designated constraints aren't met.
type EntityNodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EntityNodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EntityNodeMultiError) AllErrors() []error { return m }

// EntityNodeValidationError is the validation error returned by
// EntityNode.Validate if the designated constraints aren't met.
type EntityNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityNodeValidationError) ErrorName() string { return "EntityNodeValidationError" }

// Error satisfies the builtin error interface
func (e EntityNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntityNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityNodeValidationError{}

// Validate checks the field values on WildcardNode with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WildcardNode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WildcardNode with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WildcardNodeMultiError, or
// nil if none found.
func (m *WildcardNode) ValidateAll() error {
	return m.validate(true)
}

func (m *WildcardNode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Resource

	if len(errors) > 0 {
		return WildcardNodeMultiError(errors)
	}

	return nil
}

// WildcardNodeMultiError is an error wrapping multiple validation errors
// returned by WildcardNode.ValidateAll() if the designated constraints aren't met.
type WildcardNodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WildcardNodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WildcardNodeMultiError) AllErrors() []error { return m }

// WildcardNodeValidationError is the validation error returned by
// WildcardNode.Validate if the designated constraints aren't met.
type WildcardNodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WildcardNodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WildcardNodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WildcardNodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WildcardNodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WildcardNodeValidationError) ErrorName() string { return "WildcardNodeValidationError" }

// Error satisfies the builtin error interface
func (e WildcardNodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWildcardNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WildcardNodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WildcardNodeValidationError{}

// Validate checks the field values on RelationTree with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RelationTree) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RelationTree with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RelationTreeMultiError, or
// nil if none found.
func (m *RelationTree) ValidateAll() error {
	return m.validate(true)
}

func (m *RelationTree) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RelationTreeValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RelationTreeValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelationTreeValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetChildren() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RelationTreeValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RelationTreeValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RelationTreeValidationError{
					field:  fmt.Sprintf("Children[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RelationTreeMultiError(errors)
	}

	return nil
}

// RelationTreeMultiError is an error wrapping multiple validation errors
// returned by RelationTree.ValidateAll() if the designated constraints aren't met.
type RelationTreeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RelationTreeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RelationTreeMultiError) AllErrors() []error { return m }

// RelationTreeValidationError is the validation error returned by
// RelationTree.Validate if the designated constraints aren't met.
type RelationTreeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelationTreeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelationTreeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelationTreeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelationTreeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelationTreeValidationError) ErrorName() string { return "RelationTreeValidationError" }

// Error satisfies the builtin error interface
func (e RelationTreeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelationTree.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelationTreeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelationTreeValidationError{}

// Validate checks the field values on AccessRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AccessRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccessRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AccessRequestMultiError, or
// nil if none found.
func (m *AccessRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AccessRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetObject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccessRequestValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccessRequestValidationError{
					field:  "Object",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetObject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessRequestValidationError{
				field:  "Object",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Relation

	if all {
		switch v := interface{}(m.GetSubject()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AccessRequestValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AccessRequestValidationError{
					field:  "Subject",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSubject()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessRequestValidationError{
				field:  "Subject",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AccessRequestMultiError(errors)
	}

	return nil
}

// AccessRequestMultiError is an error wrapping multiple validation errors
// returned by AccessRequest.ValidateAll() if the designated constraints
// aren't met.
type AccessRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccessRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccessRequestMultiError) AllErrors() []error { return m }

// AccessRequestValidationError is the validation error returned by
// AccessRequest.Validate if the designated constraints aren't met.
type AccessRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccessRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccessRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccessRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccessRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccessRequestValidationError) ErrorName() string { return "AccessRequestValidationError" }

// Error satisfies the builtin error interface
func (e AccessRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccessRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccessRequestValidationError{}
