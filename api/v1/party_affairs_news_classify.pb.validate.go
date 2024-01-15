// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: party_affairs_news_classify.proto

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

// Validate checks the field values on NewsClassify with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NewsClassify) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewsClassify with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NewsClassifyMultiError, or
// nil if none found.
func (m *NewsClassify) ValidateAll() error {
	return m.validate(true)
}

func (m *NewsClassify) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if m.Weight != nil {
		// no validation rules for Weight
	}

	if len(errors) > 0 {
		return NewsClassifyMultiError(errors)
	}

	return nil
}

// NewsClassifyMultiError is an error wrapping multiple validation errors
// returned by NewsClassify.ValidateAll() if the designated constraints aren't met.
type NewsClassifyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewsClassifyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewsClassifyMultiError) AllErrors() []error { return m }

// NewsClassifyValidationError is the validation error returned by
// NewsClassify.Validate if the designated constraints aren't met.
type NewsClassifyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewsClassifyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewsClassifyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewsClassifyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewsClassifyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewsClassifyValidationError) ErrorName() string { return "NewsClassifyValidationError" }

// Error satisfies the builtin error interface
func (e NewsClassifyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewsClassify.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewsClassifyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewsClassifyValidationError{}

// Validate checks the field values on AllNewsClassifyReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AllNewsClassifyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AllNewsClassifyReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AllNewsClassifyReplyMultiError, or nil if none found.
func (m *AllNewsClassifyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AllNewsClassifyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AllNewsClassifyReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AllNewsClassifyReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AllNewsClassifyReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AllNewsClassifyReplyMultiError(errors)
	}

	return nil
}

// AllNewsClassifyReplyMultiError is an error wrapping multiple validation
// errors returned by AllNewsClassifyReply.ValidateAll() if the designated
// constraints aren't met.
type AllNewsClassifyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AllNewsClassifyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AllNewsClassifyReplyMultiError) AllErrors() []error { return m }

// AllNewsClassifyReplyValidationError is the validation error returned by
// AllNewsClassifyReply.Validate if the designated constraints aren't met.
type AllNewsClassifyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AllNewsClassifyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AllNewsClassifyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AllNewsClassifyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AllNewsClassifyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AllNewsClassifyReplyValidationError) ErrorName() string {
	return "AllNewsClassifyReplyValidationError"
}

// Error satisfies the builtin error interface
func (e AllNewsClassifyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAllNewsClassifyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AllNewsClassifyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AllNewsClassifyReplyValidationError{}

// Validate checks the field values on AddNewsClassifyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddNewsClassifyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddNewsClassifyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddNewsClassifyRequestMultiError, or nil if none found.
func (m *AddNewsClassifyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddNewsClassifyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := AddNewsClassifyRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetWeight() < 0 {
		err := AddNewsClassifyRequestValidationError{
			field:  "Weight",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddNewsClassifyRequestMultiError(errors)
	}

	return nil
}

// AddNewsClassifyRequestMultiError is an error wrapping multiple validation
// errors returned by AddNewsClassifyRequest.ValidateAll() if the designated
// constraints aren't met.
type AddNewsClassifyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddNewsClassifyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddNewsClassifyRequestMultiError) AllErrors() []error { return m }

// AddNewsClassifyRequestValidationError is the validation error returned by
// AddNewsClassifyRequest.Validate if the designated constraints aren't met.
type AddNewsClassifyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddNewsClassifyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddNewsClassifyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddNewsClassifyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddNewsClassifyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddNewsClassifyRequestValidationError) ErrorName() string {
	return "AddNewsClassifyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddNewsClassifyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddNewsClassifyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddNewsClassifyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddNewsClassifyRequestValidationError{}

// Validate checks the field values on UpdateNewsClassifyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateNewsClassifyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateNewsClassifyRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateNewsClassifyRequestMultiError, or nil if none found.
func (m *UpdateNewsClassifyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateNewsClassifyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateNewsClassifyRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := UpdateNewsClassifyRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetWeight() < 0 {
		err := UpdateNewsClassifyRequestValidationError{
			field:  "Weight",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateNewsClassifyRequestMultiError(errors)
	}

	return nil
}

// UpdateNewsClassifyRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateNewsClassifyRequest.ValidateAll() if the
// designated constraints aren't met.
type UpdateNewsClassifyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateNewsClassifyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateNewsClassifyRequestMultiError) AllErrors() []error { return m }

// UpdateNewsClassifyRequestValidationError is the validation error returned by
// UpdateNewsClassifyRequest.Validate if the designated constraints aren't met.
type UpdateNewsClassifyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNewsClassifyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNewsClassifyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNewsClassifyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNewsClassifyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNewsClassifyRequestValidationError) ErrorName() string {
	return "UpdateNewsClassifyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNewsClassifyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNewsClassifyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNewsClassifyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNewsClassifyRequestValidationError{}

// Validate checks the field values on DeleteNewsClassifyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteNewsClassifyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteNewsClassifyRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteNewsClassifyRequestMultiError, or nil if none found.
func (m *DeleteNewsClassifyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteNewsClassifyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteNewsClassifyRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteNewsClassifyRequestMultiError(errors)
	}

	return nil
}

// DeleteNewsClassifyRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteNewsClassifyRequest.ValidateAll() if the
// designated constraints aren't met.
type DeleteNewsClassifyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteNewsClassifyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteNewsClassifyRequestMultiError) AllErrors() []error { return m }

// DeleteNewsClassifyRequestValidationError is the validation error returned by
// DeleteNewsClassifyRequest.Validate if the designated constraints aren't met.
type DeleteNewsClassifyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteNewsClassifyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteNewsClassifyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteNewsClassifyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteNewsClassifyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteNewsClassifyRequestValidationError) ErrorName() string {
	return "DeleteNewsClassifyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteNewsClassifyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteNewsClassifyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteNewsClassifyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteNewsClassifyRequestValidationError{}