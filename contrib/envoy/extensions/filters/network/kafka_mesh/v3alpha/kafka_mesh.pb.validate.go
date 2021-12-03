// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: contrib/envoy/extensions/filters/network/kafka_mesh/v3alpha/kafka_mesh.proto

package v3alpha

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

// Validate checks the field values on KafkaMesh with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *KafkaMesh) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KafkaMesh with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in KafkaMeshMultiError, or nil
// if none found.
func (m *KafkaMesh) ValidateAll() error {
	return m.validate(true)
}

func (m *KafkaMesh) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetAdvertisedHost()) < 1 {
		err := KafkaMeshValidationError{
			field:  "AdvertisedHost",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAdvertisedPort() <= 0 {
		err := KafkaMeshValidationError{
			field:  "AdvertisedPort",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetUpstreamClusters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, KafkaMeshValidationError{
						field:  fmt.Sprintf("UpstreamClusters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, KafkaMeshValidationError{
						field:  fmt.Sprintf("UpstreamClusters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return KafkaMeshValidationError{
					field:  fmt.Sprintf("UpstreamClusters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetForwardingRules() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, KafkaMeshValidationError{
						field:  fmt.Sprintf("ForwardingRules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, KafkaMeshValidationError{
						field:  fmt.Sprintf("ForwardingRules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return KafkaMeshValidationError{
					field:  fmt.Sprintf("ForwardingRules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return KafkaMeshMultiError(errors)
	}
	return nil
}

// KafkaMeshMultiError is an error wrapping multiple validation errors returned
// by KafkaMesh.ValidateAll() if the designated constraints aren't met.
type KafkaMeshMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KafkaMeshMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KafkaMeshMultiError) AllErrors() []error { return m }

// KafkaMeshValidationError is the validation error returned by
// KafkaMesh.Validate if the designated constraints aren't met.
type KafkaMeshValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KafkaMeshValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KafkaMeshValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KafkaMeshValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KafkaMeshValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KafkaMeshValidationError) ErrorName() string { return "KafkaMeshValidationError" }

// Error satisfies the builtin error interface
func (e KafkaMeshValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKafkaMesh.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KafkaMeshValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KafkaMeshValidationError{}

// Validate checks the field values on KafkaClusterDefinition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *KafkaClusterDefinition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on KafkaClusterDefinition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// KafkaClusterDefinitionMultiError, or nil if none found.
func (m *KafkaClusterDefinition) ValidateAll() error {
	return m.validate(true)
}

func (m *KafkaClusterDefinition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetClusterName()) < 1 {
		err := KafkaClusterDefinitionValidationError{
			field:  "ClusterName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetBootstrapServers()) < 1 {
		err := KafkaClusterDefinitionValidationError{
			field:  "BootstrapServers",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPartitionCount() <= 0 {
		err := KafkaClusterDefinitionValidationError{
			field:  "PartitionCount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ProducerConfig

	if len(errors) > 0 {
		return KafkaClusterDefinitionMultiError(errors)
	}
	return nil
}

// KafkaClusterDefinitionMultiError is an error wrapping multiple validation
// errors returned by KafkaClusterDefinition.ValidateAll() if the designated
// constraints aren't met.
type KafkaClusterDefinitionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m KafkaClusterDefinitionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m KafkaClusterDefinitionMultiError) AllErrors() []error { return m }

// KafkaClusterDefinitionValidationError is the validation error returned by
// KafkaClusterDefinition.Validate if the designated constraints aren't met.
type KafkaClusterDefinitionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KafkaClusterDefinitionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KafkaClusterDefinitionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KafkaClusterDefinitionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KafkaClusterDefinitionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KafkaClusterDefinitionValidationError) ErrorName() string {
	return "KafkaClusterDefinitionValidationError"
}

// Error satisfies the builtin error interface
func (e KafkaClusterDefinitionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKafkaClusterDefinition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KafkaClusterDefinitionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KafkaClusterDefinitionValidationError{}

// Validate checks the field values on ForwardingRule with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ForwardingRule) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForwardingRule with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ForwardingRuleMultiError,
// or nil if none found.
func (m *ForwardingRule) ValidateAll() error {
	return m.validate(true)
}

func (m *ForwardingRule) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TargetCluster

	switch m.Trigger.(type) {

	case *ForwardingRule_TopicPrefix:
		// no validation rules for TopicPrefix

	}

	if len(errors) > 0 {
		return ForwardingRuleMultiError(errors)
	}
	return nil
}

// ForwardingRuleMultiError is an error wrapping multiple validation errors
// returned by ForwardingRule.ValidateAll() if the designated constraints
// aren't met.
type ForwardingRuleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForwardingRuleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForwardingRuleMultiError) AllErrors() []error { return m }

// ForwardingRuleValidationError is the validation error returned by
// ForwardingRule.Validate if the designated constraints aren't met.
type ForwardingRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForwardingRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForwardingRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForwardingRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForwardingRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForwardingRuleValidationError) ErrorName() string { return "ForwardingRuleValidationError" }

// Error satisfies the builtin error interface
func (e ForwardingRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForwardingRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForwardingRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForwardingRuleValidationError{}
