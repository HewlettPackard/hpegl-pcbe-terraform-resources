// (C) Copyright 2025 Hewlett Packard Enterprise Development LP

package errordefs

import (
	"fmt"
)

type ResourceError struct {
	code    string
	message string
}

func (e *ResourceError) Error() string {
	return fmt.Sprintf("[%s] %s", e.code, e.message)
}

const (
	ErrNotFoundCode = "ERR_NOT_FOUND"
	ErrPollCode     = "ERR_POLL"
	ErrCreateCode   = "ERR_CREATE"
	ErrNoURICode    = "ERR_NO_URI"
	ErrClientCode   = "ERR_CLIENT"
)

func NewNotFoundError(resource string) *ResourceError {
	return &ResourceError{
		code: ErrNotFoundCode,
		message: fmt.Sprintf(
			"resource %s not found", resource,
		),
	}
}

func NewPollError(resource string) *ResourceError {
	return &ResourceError{
		code: ErrPollCode,
		message: fmt.Sprintf(
			"failed to poll task for resource %s", resource,
		),
	}
}

func NewCreateError(resource string) *ResourceError {
	return &ResourceError{
		code: ErrCreateCode,
		message: fmt.Sprintf(
			"failed to create resource %s", resource,
		),
	}
}

func NewNoURIError(resource string) *ResourceError {
	return &ResourceError{
		code: ErrCreateCode,
		message: fmt.Sprintf(
			"failed to get uri for resource %s", resource,
		),
	}
}

func NewClientError(resource string) *ResourceError {
	return &ResourceError{
		code: ErrCreateCode,
		message: fmt.Sprintf(
			"failed to create client for resource %s", resource,
		),
	}
}

func (e *ResourceError) IsNotFound() bool {
	return e.code == ErrNotFoundCode
}

func (e *ResourceError) IsPoll() bool {
	return e.code == ErrPollCode
}

func (e *ResourceError) IsCreate() bool {
	return e.code == ErrCreateCode
}

func (e *ResourceError) IsNoURI() bool {
	return e.code == ErrNoURICode
}

func (e *ResourceError) IsClient() bool {
	return e.code == ErrClientCode
}
