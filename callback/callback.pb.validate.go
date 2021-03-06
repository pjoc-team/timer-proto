// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: callback/callback.proto

package callback

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

	"github.com/golang/protobuf/ptypes"

	job "github.com/pjoc-team/timer-proto/job"
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
	_ = ptypes.DynamicAny{}

	_ = job.Status(0)
)

// define the regex for a UUID once up-front
var _callback_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CallbackRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CallbackRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	switch m.Payload.(type) {

	case *CallbackRequest_Job:

		if v, ok := interface{}(m.GetJob()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CallbackRequestValidationError{
					field:  "Job",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CallbackRequest_Task:

		if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CallbackRequestValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CallbackRequestValidationError is the validation error returned by
// CallbackRequest.Validate if the designated constraints aren't met.
type CallbackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallbackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallbackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallbackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallbackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallbackRequestValidationError) ErrorName() string { return "CallbackRequestValidationError" }

// Error satisfies the builtin error interface
func (e CallbackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallbackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallbackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallbackRequestValidationError{}

// Validate checks the field values on CallbackResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CallbackResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CallbackResponseValidationError is the validation error returned by
// CallbackResponse.Validate if the designated constraints aren't met.
type CallbackResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CallbackResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CallbackResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CallbackResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CallbackResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CallbackResponseValidationError) ErrorName() string { return "CallbackResponseValidationError" }

// Error satisfies the builtin error interface
func (e CallbackResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCallbackResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CallbackResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CallbackResponseValidationError{}
