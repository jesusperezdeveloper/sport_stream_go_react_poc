package domain

import "errors"

var (
	ErrNotFound            = errors.New("resource not found")
	ErrAlreadyExists       = errors.New("resource already exists")
	ErrInvalidInput        = errors.New("invalid input")
	ErrInvalidTransition   = errors.New("invalid status transition")
	ErrClubNotFound        = errors.New("club not found")
	ErrStreamNotFound      = errors.New("stream not found")
	ErrEventNotFound       = errors.New("event not found")
	ErrInvalidStreamStatus = errors.New("invalid stream status")
	ErrInvalidEventStatus  = errors.New("invalid event status")
	ErrInvalidStreamType   = errors.New("invalid stream type")
)
