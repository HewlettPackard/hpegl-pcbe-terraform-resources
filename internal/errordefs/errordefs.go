// (C) Copyright 2025 Hewlett Packard Enterprise Development LP

package errordefs

import (
	"fmt"
)

type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return e.message
}

func NewNotFoundError(resource string) *NotFoundError {
	return &NotFoundError{
		message: fmt.Sprintf(
			"resource %s, not found error", resource,
		),
	}
}

var NotFound *NotFoundError

type PollError struct {
	message string
}

func (e *PollError) Error() string {
	return e.message
}

func NewPollError(resource string) *PollError {
	return &PollError{
		message: fmt.Sprintf(
			"resource %s, poll error", resource,
		),
	}
}

var Poll *PollError

type CreateError struct {
	message string
}

func (e *CreateError) Error() string {
	return e.message
}

func NewCreateError(resource string) *CreateError {
	return &CreateError{
		message: fmt.Sprintf(
			"resource %s, create error", resource,
		),
	}
}

var Create *CreateError

type NoURIError struct {
	message string
}

func (e *NoURIError) Error() string {
	return e.message
}

func NewNoURIError(resource string) *NoURIError {
	return &NoURIError{
		message: fmt.Sprintf(
			"resource %s, nouri error", resource,
		),
	}
}

var NoURI *NoURIError

type ClientError struct {
	message string
}

func (e *ClientError) Error() string {
	return e.message
}

func NewClientError(resource string) *ClientError {
	return &ClientError{
		message: fmt.Sprintf(
			"resource %s, client error", resource,
		),
	}
}

var Client *ClientError

type ValueError struct {
	message string
}

func (e *ValueError) Error() string {
	return e.message
}

func NewValueError(value string) *ValueError {
	return &ValueError{
		message: fmt.Sprintf(
			"value error: %s", value,
		),
	}
}

var Value *ValueError
