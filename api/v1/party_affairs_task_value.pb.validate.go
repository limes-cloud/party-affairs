// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: party_affairs_task_value.proto

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

// Validate checks the field values on TaskValue with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TaskValue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TaskValue with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TaskValueMultiError, or nil
// if none found.
func (m *TaskValue) ValidateAll() error {
	return m.validate(true)
}

func (m *TaskValue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for TaskId

	// no validation rules for UserId

	// no validation rules for Value

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TaskValueValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TaskValueValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskValueValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TaskValueMultiError(errors)
	}

	return nil
}

// TaskValueMultiError is an error wrapping multiple validation errors returned
// by TaskValue.ValidateAll() if the designated constraints aren't met.
type TaskValueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TaskValueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TaskValueMultiError) AllErrors() []error { return m }

// TaskValueValidationError is the validation error returned by
// TaskValue.Validate if the designated constraints aren't met.
type TaskValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskValueValidationError) ErrorName() string { return "TaskValueValidationError" }

// Error satisfies the builtin error interface
func (e TaskValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTaskValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskValueValidationError{}

// Validate checks the field values on GetCurTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCurTaskValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCurTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCurTaskValueRequestMultiError, or nil if none found.
func (m *GetCurTaskValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCurTaskValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTaskId() <= 0 {
		err := GetCurTaskValueRequestValidationError{
			field:  "TaskId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetCurTaskValueRequestMultiError(errors)
	}

	return nil
}

// GetCurTaskValueRequestMultiError is an error wrapping multiple validation
// errors returned by GetCurTaskValueRequest.ValidateAll() if the designated
// constraints aren't met.
type GetCurTaskValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCurTaskValueRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCurTaskValueRequestMultiError) AllErrors() []error { return m }

// GetCurTaskValueRequestValidationError is the validation error returned by
// GetCurTaskValueRequest.Validate if the designated constraints aren't met.
type GetCurTaskValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCurTaskValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCurTaskValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCurTaskValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCurTaskValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCurTaskValueRequestValidationError) ErrorName() string {
	return "GetCurTaskValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCurTaskValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCurTaskValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCurTaskValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCurTaskValueRequestValidationError{}

// Validate checks the field values on GetTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTaskValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTaskValueRequestMultiError, or nil if none found.
func (m *GetTaskValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTaskValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTaskId() <= 0 {
		err := GetTaskValueRequestValidationError{
			field:  "TaskId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUserId() <= 0 {
		err := GetTaskValueRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetTaskValueRequestMultiError(errors)
	}

	return nil
}

// GetTaskValueRequestMultiError is an error wrapping multiple validation
// errors returned by GetTaskValueRequest.ValidateAll() if the designated
// constraints aren't met.
type GetTaskValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTaskValueRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTaskValueRequestMultiError) AllErrors() []error { return m }

// GetTaskValueRequestValidationError is the validation error returned by
// GetTaskValueRequest.Validate if the designated constraints aren't met.
type GetTaskValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTaskValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTaskValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTaskValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTaskValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTaskValueRequestValidationError) ErrorName() string {
	return "GetTaskValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTaskValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTaskValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTaskValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTaskValueRequestValidationError{}

// Validate checks the field values on PageTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PageTaskValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageTaskValueRequestMultiError, or nil if none found.
func (m *PageTaskValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageTaskValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := PageTaskValueRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val <= 0 || val > 50 {
		err := PageTaskValueRequestValidationError{
			field:  "PageSize",
			reason: "value must be inside range (0, 50]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTaskId() <= 0 {
		err := PageTaskValueRequestValidationError{
			field:  "TaskId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.UserId != nil {

		if m.GetUserId() <= 0 {
			err := PageTaskValueRequestValidationError{
				field:  "UserId",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return PageTaskValueRequestMultiError(errors)
	}

	return nil
}

// PageTaskValueRequestMultiError is an error wrapping multiple validation
// errors returned by PageTaskValueRequest.ValidateAll() if the designated
// constraints aren't met.
type PageTaskValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageTaskValueRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageTaskValueRequestMultiError) AllErrors() []error { return m }

// PageTaskValueRequestValidationError is the validation error returned by
// PageTaskValueRequest.Validate if the designated constraints aren't met.
type PageTaskValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageTaskValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageTaskValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageTaskValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageTaskValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageTaskValueRequestValidationError) ErrorName() string {
	return "PageTaskValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PageTaskValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageTaskValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageTaskValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageTaskValueRequestValidationError{}

// Validate checks the field values on PageTaskValueReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PageTaskValueReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageTaskValueReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageTaskValueReplyMultiError, or nil if none found.
func (m *PageTaskValueReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PageTaskValueReply) validate(all bool) error {
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
					errors = append(errors, PageTaskValueReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PageTaskValueReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PageTaskValueReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PageTaskValueReplyMultiError(errors)
	}

	return nil
}

// PageTaskValueReplyMultiError is an error wrapping multiple validation errors
// returned by PageTaskValueReply.ValidateAll() if the designated constraints
// aren't met.
type PageTaskValueReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageTaskValueReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageTaskValueReplyMultiError) AllErrors() []error { return m }

// PageTaskValueReplyValidationError is the validation error returned by
// PageTaskValueReply.Validate if the designated constraints aren't met.
type PageTaskValueReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageTaskValueReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageTaskValueReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageTaskValueReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageTaskValueReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageTaskValueReplyValidationError) ErrorName() string {
	return "PageTaskValueReplyValidationError"
}

// Error satisfies the builtin error interface
func (e PageTaskValueReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageTaskValueReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageTaskValueReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageTaskValueReplyValidationError{}

// Validate checks the field values on ExportTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExportTaskValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExportTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExportTaskValueRequestMultiError, or nil if none found.
func (m *ExportTaskValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ExportTaskValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTaskId() <= 0 {
		err := ExportTaskValueRequestValidationError{
			field:  "TaskId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ExportTaskValueRequestMultiError(errors)
	}

	return nil
}

// ExportTaskValueRequestMultiError is an error wrapping multiple validation
// errors returned by ExportTaskValueRequest.ValidateAll() if the designated
// constraints aren't met.
type ExportTaskValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExportTaskValueRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExportTaskValueRequestMultiError) AllErrors() []error { return m }

// ExportTaskValueRequestValidationError is the validation error returned by
// ExportTaskValueRequest.Validate if the designated constraints aren't met.
type ExportTaskValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExportTaskValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExportTaskValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExportTaskValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExportTaskValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExportTaskValueRequestValidationError) ErrorName() string {
	return "ExportTaskValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ExportTaskValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExportTaskValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExportTaskValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExportTaskValueRequestValidationError{}

// Validate checks the field values on AddTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddTaskValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddTaskValueRequestMultiError, or nil if none found.
func (m *AddTaskValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddTaskValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTaskId() <= 0 {
		err := AddTaskValueRequestValidationError{
			field:  "TaskId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetValue()) < 1 {
		err := AddTaskValueRequestValidationError{
			field:  "Value",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddTaskValueRequestMultiError(errors)
	}

	return nil
}

// AddTaskValueRequestMultiError is an error wrapping multiple validation
// errors returned by AddTaskValueRequest.ValidateAll() if the designated
// constraints aren't met.
type AddTaskValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddTaskValueRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddTaskValueRequestMultiError) AllErrors() []error { return m }

// AddTaskValueRequestValidationError is the validation error returned by
// AddTaskValueRequest.Validate if the designated constraints aren't met.
type AddTaskValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddTaskValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddTaskValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddTaskValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddTaskValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddTaskValueRequestValidationError) ErrorName() string {
	return "AddTaskValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddTaskValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddTaskValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddTaskValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddTaskValueRequestValidationError{}

// Validate checks the field values on UpdateTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateTaskValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTaskValueRequestMultiError, or nil if none found.
func (m *UpdateTaskValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTaskValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTaskId() <= 0 {
		err := UpdateTaskValueRequestValidationError{
			field:  "TaskId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetValue()) < 1 {
		err := UpdateTaskValueRequestValidationError{
			field:  "Value",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateTaskValueRequestMultiError(errors)
	}

	return nil
}

// UpdateTaskValueRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateTaskValueRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateTaskValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTaskValueRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTaskValueRequestMultiError) AllErrors() []error { return m }

// UpdateTaskValueRequestValidationError is the validation error returned by
// UpdateTaskValueRequest.Validate if the designated constraints aren't met.
type UpdateTaskValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTaskValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTaskValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTaskValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTaskValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTaskValueRequestValidationError) ErrorName() string {
	return "UpdateTaskValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTaskValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTaskValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTaskValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTaskValueRequestValidationError{}

// Validate checks the field values on DeleteTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTaskValueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTaskValueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTaskValueRequestMultiError, or nil if none found.
func (m *DeleteTaskValueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTaskValueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteTaskValueRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteTaskValueRequestMultiError(errors)
	}

	return nil
}

// DeleteTaskValueRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteTaskValueRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteTaskValueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTaskValueRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTaskValueRequestMultiError) AllErrors() []error { return m }

// DeleteTaskValueRequestValidationError is the validation error returned by
// DeleteTaskValueRequest.Validate if the designated constraints aren't met.
type DeleteTaskValueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTaskValueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTaskValueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTaskValueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTaskValueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTaskValueRequestValidationError) ErrorName() string {
	return "DeleteTaskValueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTaskValueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTaskValueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTaskValueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTaskValueRequestValidationError{}
