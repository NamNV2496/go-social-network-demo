// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: post_core/v1/post.proto

package postv1

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

// Validate checks the field values on Post with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Post) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Post with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PostMultiError, or nil if none found.
func (m *Post) ValidateAll() error {
	return m.validate(true)
}

func (m *Post) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for ContentText

	// no validation rules for Visible

	// no validation rules for Date

	if len(errors) > 0 {
		return PostMultiError(errors)
	}

	return nil
}

// PostMultiError is an error wrapping multiple validation errors returned by
// Post.ValidateAll() if the designated constraints aren't met.
type PostMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PostMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PostMultiError) AllErrors() []error { return m }

// PostValidationError is the validation error returned by Post.Validate if the
// designated constraints aren't met.
type PostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostValidationError) ErrorName() string { return "PostValidationError" }

// Error satisfies the builtin error interface
func (e PostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostValidationError{}

// Validate checks the field values on CreatePostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreatePostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePostRequestMultiError, or nil if none found.
func (m *CreatePostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreatePostRequestValidationError{
					field:  "Post",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreatePostRequestValidationError{
					field:  "Post",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatePostRequestValidationError{
				field:  "Post",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreatePostRequestMultiError(errors)
	}

	return nil
}

// CreatePostRequestMultiError is an error wrapping multiple validation errors
// returned by CreatePostRequest.ValidateAll() if the designated constraints
// aren't met.
type CreatePostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePostRequestMultiError) AllErrors() []error { return m }

// CreatePostRequestValidationError is the validation error returned by
// CreatePostRequest.Validate if the designated constraints aren't met.
type CreatePostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePostRequestValidationError) ErrorName() string {
	return "CreatePostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePostRequestValidationError{}

// Validate checks the field values on CreatePostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreatePostResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePostResponseMultiError, or nil if none found.
func (m *CreatePostResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePostResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PostId

	if len(errors) > 0 {
		return CreatePostResponseMultiError(errors)
	}

	return nil
}

// CreatePostResponseMultiError is an error wrapping multiple validation errors
// returned by CreatePostResponse.ValidateAll() if the designated constraints
// aren't met.
type CreatePostResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePostResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePostResponseMultiError) AllErrors() []error { return m }

// CreatePostResponseValidationError is the validation error returned by
// CreatePostResponse.Validate if the designated constraints aren't met.
type CreatePostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePostResponseValidationError) ErrorName() string {
	return "CreatePostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePostResponseValidationError{}

// Validate checks the field values on GetPostRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetPostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetPostRequestMultiError,
// or nil if none found.
func (m *GetPostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return GetPostRequestMultiError(errors)
	}

	return nil
}

// GetPostRequestMultiError is an error wrapping multiple validation errors
// returned by GetPostRequest.ValidateAll() if the designated constraints
// aren't met.
type GetPostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPostRequestMultiError) AllErrors() []error { return m }

// GetPostRequestValidationError is the validation error returned by
// GetPostRequest.Validate if the designated constraints aren't met.
type GetPostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostRequestValidationError) ErrorName() string { return "GetPostRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetPostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostRequestValidationError{}

// Validate checks the field values on GetPostResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetPostResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPostResponseMultiError, or nil if none found.
func (m *GetPostResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPostResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsValid

	if all {
		switch v := interface{}(m.GetPost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetPostResponseValidationError{
					field:  "Post",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetPostResponseValidationError{
					field:  "Post",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPostResponseValidationError{
				field:  "Post",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetPostResponseMultiError(errors)
	}

	return nil
}

// GetPostResponseMultiError is an error wrapping multiple validation errors
// returned by GetPostResponse.ValidateAll() if the designated constraints
// aren't met.
type GetPostResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPostResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPostResponseMultiError) AllErrors() []error { return m }

// GetPostResponseValidationError is the validation error returned by
// GetPostResponse.Validate if the designated constraints aren't met.
type GetPostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPostResponseValidationError) ErrorName() string { return "GetPostResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetPostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPostResponseValidationError{}
