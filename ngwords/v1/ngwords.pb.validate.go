// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ngwords/v1/ngwords.proto

package ngwordsv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on ValidateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ValidateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ValidateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ValidateRequestMultiError, or nil if none found.
func (m *ValidateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ValidateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Msg

	if len(errors) > 0 {
		return ValidateRequestMultiError(errors)
	}
	return nil
}

// ValidateRequestMultiError is an error wrapping multiple validation errors
// returned by ValidateRequest.ValidateAll() if the designated constraints
// aren't met.
type ValidateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValidateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValidateRequestMultiError) AllErrors() []error { return m }

// ValidateRequestValidationError is the validation error returned by
// ValidateRequest.Validate if the designated constraints aren't met.
type ValidateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidateRequestValidationError) ErrorName() string { return "ValidateRequestValidationError" }

// Error satisfies the builtin error interface
func (e ValidateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidateRequestValidationError{}

// Validate checks the field values on ValidateResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ValidateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ValidateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ValidateResponseMultiError, or nil if none found.
func (m *ValidateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ValidateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsPass

	if len(errors) > 0 {
		return ValidateResponseMultiError(errors)
	}
	return nil
}

// ValidateResponseMultiError is an error wrapping multiple validation errors
// returned by ValidateResponse.ValidateAll() if the designated constraints
// aren't met.
type ValidateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValidateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValidateResponseMultiError) AllErrors() []error { return m }

// ValidateResponseValidationError is the validation error returned by
// ValidateResponse.Validate if the designated constraints aren't met.
type ValidateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidateResponseValidationError) ErrorName() string { return "ValidateResponseValidationError" }

// Error satisfies the builtin error interface
func (e ValidateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidateResponseValidationError{}

// Validate checks the field values on ConvertRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ConvertRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConvertRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConvertRequestMultiError,
// or nil if none found.
func (m *ConvertRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ConvertRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Msg

	// no validation rules for Mark

	if len(errors) > 0 {
		return ConvertRequestMultiError(errors)
	}
	return nil
}

// ConvertRequestMultiError is an error wrapping multiple validation errors
// returned by ConvertRequest.ValidateAll() if the designated constraints
// aren't met.
type ConvertRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConvertRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConvertRequestMultiError) AllErrors() []error { return m }

// ConvertRequestValidationError is the validation error returned by
// ConvertRequest.Validate if the designated constraints aren't met.
type ConvertRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConvertRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConvertRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConvertRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConvertRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConvertRequestValidationError) ErrorName() string { return "ConvertRequestValidationError" }

// Error satisfies the builtin error interface
func (e ConvertRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConvertRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConvertRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConvertRequestValidationError{}

// Validate checks the field values on ConvertResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConvertResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConvertResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConvertResponseMultiError, or nil if none found.
func (m *ConvertResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ConvertResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Msg

	if len(errors) > 0 {
		return ConvertResponseMultiError(errors)
	}
	return nil
}

// ConvertResponseMultiError is an error wrapping multiple validation errors
// returned by ConvertResponse.ValidateAll() if the designated constraints
// aren't met.
type ConvertResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConvertResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConvertResponseMultiError) AllErrors() []error { return m }

// ConvertResponseValidationError is the validation error returned by
// ConvertResponse.Validate if the designated constraints aren't met.
type ConvertResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConvertResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConvertResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConvertResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConvertResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConvertResponseValidationError) ErrorName() string { return "ConvertResponseValidationError" }

// Error satisfies the builtin error interface
func (e ConvertResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConvertResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConvertResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConvertResponseValidationError{}

// Validate checks the field values on BuildRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BuildRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BuildRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BuildRequestMultiError, or
// nil if none found.
func (m *BuildRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BuildRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	if len(errors) > 0 {
		return BuildRequestMultiError(errors)
	}
	return nil
}

// BuildRequestMultiError is an error wrapping multiple validation errors
// returned by BuildRequest.ValidateAll() if the designated constraints aren't met.
type BuildRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BuildRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BuildRequestMultiError) AllErrors() []error { return m }

// BuildRequestValidationError is the validation error returned by
// BuildRequest.Validate if the designated constraints aren't met.
type BuildRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BuildRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BuildRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BuildRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BuildRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BuildRequestValidationError) ErrorName() string { return "BuildRequestValidationError" }

// Error satisfies the builtin error interface
func (e BuildRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBuildRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BuildRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BuildRequestValidationError{}
