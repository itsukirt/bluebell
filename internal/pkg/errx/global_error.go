package errx

import "bluebell/internal/pkg/codes"

var (
	ErrInternalServerError   = New(codes.ErrInternalServerError)
	ErrRequestParamsInvalid  = New(codes.ErrRequestParamsInvalid)
	ErrUserNotFound          = New(codes.ErrUserNotFound)
	ErrEmailHasRegistered    = New(codes.ErrEmailHasRegistered)
	ErrUsernameHasRegistered = New(codes.ErrUsernameHasRegistered)
)
