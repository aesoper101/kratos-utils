// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package httppb

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsBadRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_BAD_REQUEST.String() && e.Code == 400
}

func ErrorBadRequest(format string, args ...interface{}) *errors.Error {
	return errors.New(400, APIErrorReason_BAD_REQUEST.String(), fmt.Sprintf(format, args...))
}

func IsAuthenticationFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_AUTHENTICATION_FAILED.String() && e.Code == 401
}

func ErrorAuthenticationFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(401, APIErrorReason_AUTHENTICATION_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsAuthorizationFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_AUTHORIZATION_FAILED.String() && e.Code == 403
}

func ErrorAuthorizationFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(403, APIErrorReason_AUTHORIZATION_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_NOT_FOUND.String() && e.Code == 404
}

func ErrorNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, APIErrorReason_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsMethodNotAllowed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_METHOD_NOT_ALLOWED.String() && e.Code == 405
}

func ErrorMethodNotAllowed(format string, args ...interface{}) *errors.Error {
	return errors.New(405, APIErrorReason_METHOD_NOT_ALLOWED.String(), fmt.Sprintf(format, args...))
}

func IsNotAcceptable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_NOT_ACCEPTABLE.String() && e.Code == 406
}

func ErrorNotAcceptable(format string, args ...interface{}) *errors.Error {
	return errors.New(406, APIErrorReason_NOT_ACCEPTABLE.String(), fmt.Sprintf(format, args...))
}

func IsRequestTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_REQUEST_TIMEOUT.String() && e.Code == 408
}

func ErrorRequestTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(408, APIErrorReason_REQUEST_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

func IsConflict(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_CONFLICT.String() && e.Code == 409
}

func ErrorConflict(format string, args ...interface{}) *errors.Error {
	return errors.New(409, APIErrorReason_CONFLICT.String(), fmt.Sprintf(format, args...))
}

func IsGone(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_GONE.String() && e.Code == 410
}

func ErrorGone(format string, args ...interface{}) *errors.Error {
	return errors.New(410, APIErrorReason_GONE.String(), fmt.Sprintf(format, args...))
}

func IsLengthRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_LENGTH_REQUIRED.String() && e.Code == 411
}

func ErrorLengthRequired(format string, args ...interface{}) *errors.Error {
	return errors.New(411, APIErrorReason_LENGTH_REQUIRED.String(), fmt.Sprintf(format, args...))
}

func IsPreconditionFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_PRECONDITION_FAILED.String() && e.Code == 412
}

func ErrorPreconditionFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(412, APIErrorReason_PRECONDITION_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsRequestEntityTooLarge(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_REQUEST_ENTITY_TOO_LARGE.String() && e.Code == 413
}

func ErrorRequestEntityTooLarge(format string, args ...interface{}) *errors.Error {
	return errors.New(413, APIErrorReason_REQUEST_ENTITY_TOO_LARGE.String(), fmt.Sprintf(format, args...))
}

func IsRequestUriTooLong(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_REQUEST_URI_TOO_LONG.String() && e.Code == 414
}

func ErrorRequestUriTooLong(format string, args ...interface{}) *errors.Error {
	return errors.New(414, APIErrorReason_REQUEST_URI_TOO_LONG.String(), fmt.Sprintf(format, args...))
}

func IsUnsupportedMediaType(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_UNSUPPORTED_MEDIA_TYPE.String() && e.Code == 415
}

func ErrorUnsupportedMediaType(format string, args ...interface{}) *errors.Error {
	return errors.New(415, APIErrorReason_UNSUPPORTED_MEDIA_TYPE.String(), fmt.Sprintf(format, args...))
}

func IsRequestedRangeNotSatisfiable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_REQUESTED_RANGE_NOT_SATISFIABLE.String() && e.Code == 416
}

func ErrorRequestedRangeNotSatisfiable(format string, args ...interface{}) *errors.Error {
	return errors.New(416, APIErrorReason_REQUESTED_RANGE_NOT_SATISFIABLE.String(), fmt.Sprintf(format, args...))
}

func IsExpectationFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_EXPECTATION_FAILED.String() && e.Code == 417
}

func ErrorExpectationFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(417, APIErrorReason_EXPECTATION_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsMisdirectedRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_MISDIRECTED_REQUEST.String() && e.Code == 421
}

func ErrorMisdirectedRequest(format string, args ...interface{}) *errors.Error {
	return errors.New(421, APIErrorReason_MISDIRECTED_REQUEST.String(), fmt.Sprintf(format, args...))
}

func IsUnprocessableEntity(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_UNPROCESSABLE_ENTITY.String() && e.Code == 422
}

func ErrorUnprocessableEntity(format string, args ...interface{}) *errors.Error {
	return errors.New(422, APIErrorReason_UNPROCESSABLE_ENTITY.String(), fmt.Sprintf(format, args...))
}

func IsLocked(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_LOCKED.String() && e.Code == 423
}

func ErrorLocked(format string, args ...interface{}) *errors.Error {
	return errors.New(423, APIErrorReason_LOCKED.String(), fmt.Sprintf(format, args...))
}

func IsFailedDependency(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_FAILED_DEPENDENCY.String() && e.Code == 424
}

func ErrorFailedDependency(format string, args ...interface{}) *errors.Error {
	return errors.New(424, APIErrorReason_FAILED_DEPENDENCY.String(), fmt.Sprintf(format, args...))
}

func IsUpgradeRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_UPGRADE_REQUIRED.String() && e.Code == 426
}

func ErrorUpgradeRequired(format string, args ...interface{}) *errors.Error {
	return errors.New(426, APIErrorReason_UPGRADE_REQUIRED.String(), fmt.Sprintf(format, args...))
}

func IsPreconditionRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_PRECONDITION_REQUIRED.String() && e.Code == 428
}

func ErrorPreconditionRequired(format string, args ...interface{}) *errors.Error {
	return errors.New(428, APIErrorReason_PRECONDITION_REQUIRED.String(), fmt.Sprintf(format, args...))
}

func IsTooManyRequests(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_TOO_MANY_REQUESTS.String() && e.Code == 429
}

func ErrorTooManyRequests(format string, args ...interface{}) *errors.Error {
	return errors.New(429, APIErrorReason_TOO_MANY_REQUESTS.String(), fmt.Sprintf(format, args...))
}

func IsRequestHeaderFieldsTooLarge(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_REQUEST_HEADER_FIELDS_TOO_LARGE.String() && e.Code == 431
}

func ErrorRequestHeaderFieldsTooLarge(format string, args ...interface{}) *errors.Error {
	return errors.New(431, APIErrorReason_REQUEST_HEADER_FIELDS_TOO_LARGE.String(), fmt.Sprintf(format, args...))
}

func IsUnavailableForLegalReasons(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_UNAVAILABLE_FOR_LEGAL_REASONS.String() && e.Code == 451
}

func ErrorUnavailableForLegalReasons(format string, args ...interface{}) *errors.Error {
	return errors.New(451, APIErrorReason_UNAVAILABLE_FOR_LEGAL_REASONS.String(), fmt.Sprintf(format, args...))
}

func IsInternalServerError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_INTERNAL_SERVER_ERROR.String() && e.Code == 500
}

func ErrorInternalServerError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, APIErrorReason_INTERNAL_SERVER_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsNotImplemented(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_NOT_IMPLEMENTED.String() && e.Code == 501
}

func ErrorNotImplemented(format string, args ...interface{}) *errors.Error {
	return errors.New(501, APIErrorReason_NOT_IMPLEMENTED.String(), fmt.Sprintf(format, args...))
}

func IsBadGateway(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_BAD_GATEWAY.String() && e.Code == 502
}

func ErrorBadGateway(format string, args ...interface{}) *errors.Error {
	return errors.New(502, APIErrorReason_BAD_GATEWAY.String(), fmt.Sprintf(format, args...))
}

func IsServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_SERVICE_UNAVAILABLE.String() && e.Code == 503
}

func ErrorServiceUnavailable(format string, args ...interface{}) *errors.Error {
	return errors.New(503, APIErrorReason_SERVICE_UNAVAILABLE.String(), fmt.Sprintf(format, args...))
}

func IsGatewayTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_GATEWAY_TIMEOUT.String() && e.Code == 504
}

func ErrorGatewayTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(504, APIErrorReason_GATEWAY_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

func IsHttpVersionNotSupported(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_HTTP_VERSION_NOT_SUPPORTED.String() && e.Code == 505
}

func ErrorHttpVersionNotSupported(format string, args ...interface{}) *errors.Error {
	return errors.New(505, APIErrorReason_HTTP_VERSION_NOT_SUPPORTED.String(), fmt.Sprintf(format, args...))
}

func IsVariantAlsoNegotiates(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_VARIANT_ALSO_NEGOTIATES.String() && e.Code == 506
}

func ErrorVariantAlsoNegotiates(format string, args ...interface{}) *errors.Error {
	return errors.New(506, APIErrorReason_VARIANT_ALSO_NEGOTIATES.String(), fmt.Sprintf(format, args...))
}

func IsInsufficientStorage(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_INSUFFICIENT_STORAGE.String() && e.Code == 507
}

func ErrorInsufficientStorage(format string, args ...interface{}) *errors.Error {
	return errors.New(507, APIErrorReason_INSUFFICIENT_STORAGE.String(), fmt.Sprintf(format, args...))
}

func IsLoopDetected(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_LOOP_DETECTED.String() && e.Code == 508
}

func ErrorLoopDetected(format string, args ...interface{}) *errors.Error {
	return errors.New(508, APIErrorReason_LOOP_DETECTED.String(), fmt.Sprintf(format, args...))
}

func IsNotExtended(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_NOT_EXTENDED.String() && e.Code == 510
}

func ErrorNotExtended(format string, args ...interface{}) *errors.Error {
	return errors.New(510, APIErrorReason_NOT_EXTENDED.String(), fmt.Sprintf(format, args...))
}

func IsNetworkAuthenticationRequired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == APIErrorReason_NETWORK_AUTHENTICATION_REQUIRED.String() && e.Code == 511
}

func ErrorNetworkAuthenticationRequired(format string, args ...interface{}) *errors.Error {
	return errors.New(511, APIErrorReason_NETWORK_AUTHENTICATION_REQUIRED.String(), fmt.Sprintf(format, args...))
}
