// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/user/v1/user.proto

package v1

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

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for GroupId

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Mobile

	// no validation rules for Email

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on CreateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserRequestMultiError, or nil if none found.
func (m *CreateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for GroupId

	if l := utf8.RuneCountInString(m.GetUsername()); l < 5 || l > 128 {
		err := CreateUserRequestValidationError{
			field:  "Username",
			reason: "value length must be between 5 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 6 || l > 18 {
		err := CreateUserRequestValidationError{
			field:  "Password",
			reason: "value length must be between 6 and 18 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetMobile()); l < 6 || l > 18 {
		err := CreateUserRequestValidationError{
			field:  "Mobile",
			reason: "value length must be between 6 and 18 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetEmail()); l < 3 || l > 128 {
		err := CreateUserRequestValidationError{
			field:  "Email",
			reason: "value length must be between 3 and 128 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateUserRequestMultiError(errors)
	}

	return nil
}

// CreateUserRequestMultiError is an error wrapping multiple validation errors
// returned by CreateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserRequestMultiError) AllErrors() []error { return m }

// CreateUserRequestValidationError is the validation error returned by
// CreateUserRequest.Validate if the designated constraints aren't met.
type CreateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRequestValidationError) ErrorName() string {
	return "CreateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRequestValidationError{}

// Validate checks the field values on CreateUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateUserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserReplyMultiError, or nil if none found.
func (m *CreateUserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateUserReplyMultiError(errors)
	}

	return nil
}

// CreateUserReplyMultiError is an error wrapping multiple validation errors
// returned by CreateUserReply.ValidateAll() if the designated constraints
// aren't met.
type CreateUserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserReplyMultiError) AllErrors() []error { return m }

// CreateUserReplyValidationError is the validation error returned by
// CreateUserReply.Validate if the designated constraints aren't met.
type CreateUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserReplyValidationError) ErrorName() string { return "CreateUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserReplyValidationError{}

// Validate checks the field values on UpdateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserRequestMultiError, or nil if none found.
func (m *UpdateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateUserRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateUserRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateUserRequestValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateUserRequestValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateUserRequestMultiError(errors)
	}

	return nil
}

// UpdateUserRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserRequestMultiError) AllErrors() []error { return m }

// UpdateUserRequestValidationError is the validation error returned by
// UpdateUserRequest.Validate if the designated constraints aren't met.
type UpdateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserRequestValidationError) ErrorName() string {
	return "UpdateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserRequestValidationError{}

// Validate checks the field values on UpdateUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserReplyMultiError, or nil if none found.
func (m *UpdateUserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return UpdateUserReplyMultiError(errors)
	}

	return nil
}

// UpdateUserReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateUserReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateUserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserReplyMultiError) AllErrors() []error { return m }

// UpdateUserReplyValidationError is the validation error returned by
// UpdateUserReply.Validate if the designated constraints aren't met.
type UpdateUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserReplyValidationError) ErrorName() string { return "UpdateUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdateUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserReplyValidationError{}

// Validate checks the field values on DeleteUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUserRequestMultiError, or nil if none found.
func (m *DeleteUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteUserRequestMultiError(errors)
	}

	return nil
}

// DeleteUserRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteUserRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserRequestMultiError) AllErrors() []error { return m }

// DeleteUserRequestValidationError is the validation error returned by
// DeleteUserRequest.Validate if the designated constraints aren't met.
type DeleteUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserRequestValidationError) ErrorName() string {
	return "DeleteUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserRequestValidationError{}

// Validate checks the field values on DeleteUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteUserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUserReplyMultiError, or nil if none found.
func (m *DeleteUserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteUserReplyMultiError(errors)
	}

	return nil
}

// DeleteUserReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteUserReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteUserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserReplyMultiError) AllErrors() []error { return m }

// DeleteUserReplyValidationError is the validation error returned by
// DeleteUserReply.Validate if the designated constraints aren't met.
type DeleteUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserReplyValidationError) ErrorName() string { return "DeleteUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserReplyValidationError{}

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserRequestMultiError,
// or nil if none found.
func (m *GetUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetUserRequestMultiError(errors)
	}

	return nil
}

// GetUserRequestMultiError is an error wrapping multiple validation errors
// returned by GetUserRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserRequestMultiError) AllErrors() []error { return m }

// GetUserRequestValidationError is the validation error returned by
// GetUserRequest.Validate if the designated constraints aren't met.
type GetUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserRequestValidationError) ErrorName() string { return "GetUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserRequestValidationError{}

// Validate checks the field values on GetUserReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserReplyMultiError, or
// nil if none found.
func (m *GetUserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetUserReplyMultiError(errors)
	}

	return nil
}

// GetUserReplyMultiError is an error wrapping multiple validation errors
// returned by GetUserReply.ValidateAll() if the designated constraints aren't met.
type GetUserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserReplyMultiError) AllErrors() []error { return m }

// GetUserReplyValidationError is the validation error returned by
// GetUserReply.Validate if the designated constraints aren't met.
type GetUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReplyValidationError) ErrorName() string { return "GetUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReplyValidationError{}

// Validate checks the field values on ListUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserRequestMultiError, or nil if none found.
func (m *ListUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListUserRequestMultiError(errors)
	}

	return nil
}

// ListUserRequestMultiError is an error wrapping multiple validation errors
// returned by ListUserRequest.ValidateAll() if the designated constraints
// aren't met.
type ListUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserRequestMultiError) AllErrors() []error { return m }

// ListUserRequestValidationError is the validation error returned by
// ListUserRequest.Validate if the designated constraints aren't met.
type ListUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserRequestValidationError) ErrorName() string { return "ListUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserRequestValidationError{}

// Validate checks the field values on ListUserReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListUserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListUserReplyMultiError, or
// nil if none found.
func (m *ListUserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUserReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUserReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUserReplyValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListUserReplyMultiError(errors)
	}

	return nil
}

// ListUserReplyMultiError is an error wrapping multiple validation errors
// returned by ListUserReply.ValidateAll() if the designated constraints
// aren't met.
type ListUserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserReplyMultiError) AllErrors() []error { return m }

// ListUserReplyValidationError is the validation error returned by
// ListUserReply.Validate if the designated constraints aren't met.
type ListUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserReplyValidationError) ErrorName() string { return "ListUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserReplyValidationError{}
