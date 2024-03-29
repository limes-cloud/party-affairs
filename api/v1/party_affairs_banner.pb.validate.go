// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: party_affairs_banner.proto

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

// Validate checks the field values on Banner with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Banner) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Banner with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BannerMultiError, or nil if none found.
func (m *Banner) ValidateAll() error {
	return m.validate(true)
}

func (m *Banner) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Src

	// no validation rules for Path

	// no validation rules for Weight

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BannerValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BannerValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BannerValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BannerMultiError(errors)
	}

	return nil
}

// BannerMultiError is an error wrapping multiple validation errors returned by
// Banner.ValidateAll() if the designated constraints aren't met.
type BannerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BannerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BannerMultiError) AllErrors() []error { return m }

// BannerValidationError is the validation error returned by Banner.Validate if
// the designated constraints aren't met.
type BannerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BannerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BannerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BannerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BannerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BannerValidationError) ErrorName() string { return "BannerValidationError" }

// Error satisfies the builtin error interface
func (e BannerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBanner.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BannerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BannerValidationError{}

// Validate checks the field values on AllBannerReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AllBannerReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AllBannerReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AllBannerReplyMultiError,
// or nil if none found.
func (m *AllBannerReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AllBannerReply) validate(all bool) error {
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
					errors = append(errors, AllBannerReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AllBannerReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AllBannerReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AllBannerReplyMultiError(errors)
	}

	return nil
}

// AllBannerReplyMultiError is an error wrapping multiple validation errors
// returned by AllBannerReply.ValidateAll() if the designated constraints
// aren't met.
type AllBannerReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AllBannerReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AllBannerReplyMultiError) AllErrors() []error { return m }

// AllBannerReplyValidationError is the validation error returned by
// AllBannerReply.Validate if the designated constraints aren't met.
type AllBannerReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AllBannerReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AllBannerReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AllBannerReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AllBannerReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AllBannerReplyValidationError) ErrorName() string { return "AllBannerReplyValidationError" }

// Error satisfies the builtin error interface
func (e AllBannerReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAllBannerReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AllBannerReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AllBannerReplyValidationError{}

// Validate checks the field values on AddBannerRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddBannerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddBannerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddBannerRequestMultiError, or nil if none found.
func (m *AddBannerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddBannerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := AddBannerRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetSrc()) < 1 {
		err := AddBannerRequestValidationError{
			field:  "Src",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Path != nil {
		// no validation rules for Path
	}

	if len(errors) > 0 {
		return AddBannerRequestMultiError(errors)
	}

	return nil
}

// AddBannerRequestMultiError is an error wrapping multiple validation errors
// returned by AddBannerRequest.ValidateAll() if the designated constraints
// aren't met.
type AddBannerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddBannerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddBannerRequestMultiError) AllErrors() []error { return m }

// AddBannerRequestValidationError is the validation error returned by
// AddBannerRequest.Validate if the designated constraints aren't met.
type AddBannerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddBannerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddBannerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddBannerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddBannerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddBannerRequestValidationError) ErrorName() string { return "AddBannerRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddBannerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddBannerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddBannerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddBannerRequestValidationError{}

// Validate checks the field values on UpdateBannerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateBannerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateBannerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateBannerRequestMultiError, or nil if none found.
func (m *UpdateBannerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateBannerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := UpdateBannerRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetSrc()) < 1 {
		err := UpdateBannerRequestValidationError{
			field:  "Src",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetId() <= 0 {
		err := UpdateBannerRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Path != nil {
		// no validation rules for Path
	}

	if len(errors) > 0 {
		return UpdateBannerRequestMultiError(errors)
	}

	return nil
}

// UpdateBannerRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateBannerRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateBannerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateBannerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateBannerRequestMultiError) AllErrors() []error { return m }

// UpdateBannerRequestValidationError is the validation error returned by
// UpdateBannerRequest.Validate if the designated constraints aren't met.
type UpdateBannerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateBannerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateBannerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateBannerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateBannerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateBannerRequestValidationError) ErrorName() string {
	return "UpdateBannerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateBannerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateBannerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateBannerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateBannerRequestValidationError{}

// Validate checks the field values on DeleteBannerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteBannerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteBannerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteBannerRequestMultiError, or nil if none found.
func (m *DeleteBannerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteBannerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteBannerRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteBannerRequestMultiError(errors)
	}

	return nil
}

// DeleteBannerRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteBannerRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteBannerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteBannerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteBannerRequestMultiError) AllErrors() []error { return m }

// DeleteBannerRequestValidationError is the validation error returned by
// DeleteBannerRequest.Validate if the designated constraints aren't met.
type DeleteBannerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteBannerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteBannerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteBannerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteBannerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteBannerRequestValidationError) ErrorName() string {
	return "DeleteBannerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteBannerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteBannerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteBannerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteBannerRequestValidationError{}
