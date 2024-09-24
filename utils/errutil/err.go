package errutil

import (
	"errors"
	"net/http"
)

var (
	ErrForbidden             = errors.New("request not granted on this resource")
	ErrNotFound              = errors.New("resource not found")
	ErrRecordNotFound        = errors.New("record not found")
	ErrBadRequest            = errors.New("bad request, check param or body")
	ErrSomethingWentWrong    = errors.New("something went wrong")
	ErrEmptyRedisKeyValue    = errors.New("empty cache key or value")
	ErrRedisNil              = errors.New("cache: nil")
	ErrMarshalRequestBody    = errors.New("failed to marshal request body")
	ErrUnmarshalResponseBody = errors.New("failed to unmarshal response body")
	ErrDuplicateTask         = errors.New("task already exists")
	ErrTaskIDConflict        = errors.New("task ID conflicts with another task")

	ErrInvalidEmail              = errors.New("invalid email")
	ErrInvalidPassword           = errors.New("invalid password")
	ErrUserRolePermissions       = errors.New("failed to fetch role permissions")
	ErrCreateJwt                 = errors.New("failed to create JWT token")
	ErrAccessTokenSign           = errors.New("failed to sign access_token")
	ErrRefreshTokenSign          = errors.New("failed to sign refresh_token")
	ErrStoreTokenUuid            = errors.New("failed to store token uuid")
	ErrUpdateLastLogin           = errors.New("failed to update last login")
	ErrNoContextUser             = errors.New("failed to get user from context")
	ErrInvalidRefreshToken       = errors.New("invalid refresh_token")
	ErrInvalidAccessToken        = errors.New("invalid access_token")
	ErrInvalidPasswordResetToken = errors.New("invalid reset_token")
	ErrInvalidRefreshUuid        = errors.New("invalid refresh_uuid")
	ErrInvalidAccessUuid         = errors.New("invalid refresh_uuid")
	ErrInvalidJwtSigningMethod   = errors.New("invalid signing method while parsing jwt")
	ErrParseJwt                  = errors.New("failed to parse JWT token")
	ErrDeleteOldTokenUuid        = errors.New("failed to delete old token uuids")
	ErrSendingEmail              = errors.New("failed to send email")
	ErrNotAdmin                  = errors.New("not admin")
	ErrNotSuperAdmin             = errors.New("not super admin")
)

type ErrResponse struct {
	Error string `json:"error"`
}

type RestErr struct {
	// example: error message
	Message string `json:"message"`
	// example: 400
	Status int `json:"status"`
	// example: bad_request
	Error string `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}
