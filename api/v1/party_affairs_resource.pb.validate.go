// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: party_affairs_resource.proto

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

// Validate checks the field values on Resource with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Resource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Resource with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResourceMultiError, or nil
// if none found.
func (m *Resource) ValidateAll() error {
	return m.validate(true)
}

func (m *Resource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Desc

	// no validation rules for Url

	// no validation rules for DownloadCount

	// no validation rules for ClassifyId

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if all {
		switch v := interface{}(m.GetResourceClassify()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ResourceValidationError{
					field:  "ResourceClassify",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ResourceValidationError{
					field:  "ResourceClassify",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResourceClassify()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResourceValidationError{
				field:  "ResourceClassify",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ResourceValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ResourceValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResourceValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ResourceMultiError(errors)
	}

	return nil
}

// ResourceMultiError is an error wrapping multiple validation errors returned
// by Resource.ValidateAll() if the designated constraints aren't met.
type ResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResourceMultiError) AllErrors() []error { return m }

// ResourceValidationError is the validation error returned by
// Resource.Validate if the designated constraints aren't met.
type ResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceValidationError) ErrorName() string { return "ResourceValidationError" }

// Error satisfies the builtin error interface
func (e ResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceValidationError{}

// Validate checks the field values on GetResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetResourceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetResourceRequestMultiError, or nil if none found.
func (m *GetResourceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResourceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetResourceRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetResourceRequestMultiError(errors)
	}

	return nil
}

// GetResourceRequestMultiError is an error wrapping multiple validation errors
// returned by GetResourceRequest.ValidateAll() if the designated constraints
// aren't met.
type GetResourceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResourceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResourceRequestMultiError) AllErrors() []error { return m }

// GetResourceRequestValidationError is the validation error returned by
// GetResourceRequest.Validate if the designated constraints aren't met.
type GetResourceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResourceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResourceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResourceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResourceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResourceRequestValidationError) ErrorName() string {
	return "GetResourceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetResourceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResourceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResourceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResourceRequestValidationError{}

// Validate checks the field values on GetResourceReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetResourceReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResourceReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetResourceReplyMultiError, or nil if none found.
func (m *GetResourceReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResourceReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetResourceReplyValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetResourceReplyValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetResourceReplyValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetResourceReplyMultiError(errors)
	}

	return nil
}

// GetResourceReplyMultiError is an error wrapping multiple validation errors
// returned by GetResourceReply.ValidateAll() if the designated constraints
// aren't met.
type GetResourceReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResourceReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResourceReplyMultiError) AllErrors() []error { return m }

// GetResourceReplyValidationError is the validation error returned by
// GetResourceReply.Validate if the designated constraints aren't met.
type GetResourceReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResourceReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResourceReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResourceReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResourceReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResourceReplyValidationError) ErrorName() string { return "GetResourceReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetResourceReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResourceReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResourceReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResourceReplyValidationError{}

// Validate checks the field values on PageResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PageResourceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageResourceRequestMultiError, or nil if none found.
func (m *PageResourceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageResourceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := PageResourceRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val <= 0 || val > 50 {
		err := PageResourceRequestValidationError{
			field:  "PageSize",
			reason: "value must be inside range (0, 50]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.ClassifyId != nil {
		// no validation rules for ClassifyId
	}

	if m.Title != nil {
		// no validation rules for Title
	}

	if len(errors) > 0 {
		return PageResourceRequestMultiError(errors)
	}

	return nil
}

// PageResourceRequestMultiError is an error wrapping multiple validation
// errors returned by PageResourceRequest.ValidateAll() if the designated
// constraints aren't met.
type PageResourceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageResourceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageResourceRequestMultiError) AllErrors() []error { return m }

// PageResourceRequestValidationError is the validation error returned by
// PageResourceRequest.Validate if the designated constraints aren't met.
type PageResourceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageResourceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageResourceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageResourceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageResourceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageResourceRequestValidationError) ErrorName() string {
	return "PageResourceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PageResourceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageResourceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageResourceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageResourceRequestValidationError{}

// Validate checks the field values on PageResourceReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PageResourceReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageResourceReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageResourceReplyMultiError, or nil if none found.
func (m *PageResourceReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PageResourceReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PageResourceReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PageResourceReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PageResourceReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PageResourceReplyMultiError(errors)
	}

	return nil
}

// PageResourceReplyMultiError is an error wrapping multiple validation errors
// returned by PageResourceReply.ValidateAll() if the designated constraints
// aren't met.
type PageResourceReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageResourceReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageResourceReplyMultiError) AllErrors() []error { return m }

// PageResourceReplyValidationError is the validation error returned by
// PageResourceReply.Validate if the designated constraints aren't met.
type PageResourceReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageResourceReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageResourceReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageResourceReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageResourceReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageResourceReplyValidationError) ErrorName() string {
	return "PageResourceReplyValidationError"
}

// Error satisfies the builtin error interface
func (e PageResourceReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageResourceReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageResourceReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageResourceReplyValidationError{}

// Validate checks the field values on AddResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddResourceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddResourceRequestMultiError, or nil if none found.
func (m *AddResourceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddResourceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := AddResourceRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDesc()) < 1 {
		err := AddResourceRequestValidationError{
			field:  "Desc",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetUrl()) < 1 {
		err := AddResourceRequestValidationError{
			field:  "Url",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetClassifyId() <= 0 {
		err := AddResourceRequestValidationError{
			field:  "ClassifyId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddResourceRequestMultiError(errors)
	}

	return nil
}

// AddResourceRequestMultiError is an error wrapping multiple validation errors
// returned by AddResourceRequest.ValidateAll() if the designated constraints
// aren't met.
type AddResourceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddResourceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddResourceRequestMultiError) AllErrors() []error { return m }

// AddResourceRequestValidationError is the validation error returned by
// AddResourceRequest.Validate if the designated constraints aren't met.
type AddResourceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddResourceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddResourceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddResourceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddResourceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddResourceRequestValidationError) ErrorName() string {
	return "AddResourceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddResourceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddResourceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddResourceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddResourceRequestValidationError{}

// Validate checks the field values on UpdateResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateResourceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateResourceRequestMultiError, or nil if none found.
func (m *UpdateResourceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateResourceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := UpdateResourceRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDesc()) < 1 {
		err := UpdateResourceRequestValidationError{
			field:  "Desc",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetUrl()) < 1 {
		err := UpdateResourceRequestValidationError{
			field:  "Url",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetClassifyId() <= 0 {
		err := UpdateResourceRequestValidationError{
			field:  "ClassifyId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetId() <= 0 {
		err := UpdateResourceRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateResourceRequestMultiError(errors)
	}

	return nil
}

// UpdateResourceRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateResourceRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateResourceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateResourceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateResourceRequestMultiError) AllErrors() []error { return m }

// UpdateResourceRequestValidationError is the validation error returned by
// UpdateResourceRequest.Validate if the designated constraints aren't met.
type UpdateResourceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResourceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResourceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResourceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResourceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResourceRequestValidationError) ErrorName() string {
	return "UpdateResourceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateResourceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResourceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResourceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResourceRequestValidationError{}

// Validate checks the field values on DeleteResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteResourceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteResourceRequestMultiError, or nil if none found.
func (m *DeleteResourceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteResourceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteResourceRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteResourceRequestMultiError(errors)
	}

	return nil
}

// DeleteResourceRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteResourceRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteResourceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteResourceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteResourceRequestMultiError) AllErrors() []error { return m }

// DeleteResourceRequestValidationError is the validation error returned by
// DeleteResourceRequest.Validate if the designated constraints aren't met.
type DeleteResourceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResourceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResourceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResourceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResourceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResourceRequestValidationError) ErrorName() string {
	return "DeleteResourceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteResourceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResourceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResourceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResourceRequestValidationError{}
