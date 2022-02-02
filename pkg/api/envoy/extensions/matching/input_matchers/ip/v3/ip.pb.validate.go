// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/matching/input_matchers/ip/v3/ip.proto

package ipv3

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

// Validate checks the field values on Ip with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Ip) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Ip with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IpMultiError, or nil if none found.
func (m *Ip) ValidateAll() error {
	return m.validate(true)
}

func (m *Ip) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetCidrRanges()) < 1 {
		err := IpValidationError{
			field:  "CidrRanges",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetCidrRanges() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IpValidationError{
						field:  fmt.Sprintf("CidrRanges[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IpValidationError{
						field:  fmt.Sprintf("CidrRanges[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IpValidationError{
					field:  fmt.Sprintf("CidrRanges[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := IpValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IpMultiError(errors)
	}
	return nil
}

// IpMultiError is an error wrapping multiple validation errors returned by
// Ip.ValidateAll() if the designated constraints aren't met.
type IpMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IpMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IpMultiError) AllErrors() []error { return m }

// IpValidationError is the validation error returned by Ip.Validate if the
// designated constraints aren't met.
type IpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IpValidationError) ErrorName() string { return "IpValidationError" }

// Error satisfies the builtin error interface
func (e IpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IpValidationError{}
