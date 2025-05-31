package goexception

// InvalidArgument creates a new Exception with the InvalidArgumentCode error code.
//
// Deprecated: Use InvalidParameter instead.
func InvalidArgument(message string, errs ...any) *Exception {
	return _createException(InvalidArgumentCode, message, errs...)
}

// InvalidParameter creates a new Exception with the InvalidParameterCode error code.
// To be used when a parameter is invalid or does not meet the required criteria.
func InvalidParameter(message string, errs ...any) *Exception {
	return _createException(InvalidParameterCode, message, errs...)
}

// InvalidData creates a new Exception with the InvalidDataCode error code.
// To be used when the provided data is not valid.
func InvalidData(message string, errs ...any) *Exception {
	return _createException(InvalidDataCode, message, errs...)
}

// NotFound creates a new Exception with the NotFoundCode error code.
func NotFound(message string, errs ...any) *Exception {
	return _createException(NotFoundCode, message, errs...)
}

// AlreadyExists creates a new Exception with the AlreadyExistsCode error code.
func AlreadyExists(message string, errs ...any) *Exception {
	return _createException(AlreadyExistsCode, message, errs...)
}

// PermissionDenied creates a new Exception with the PermissionDeniedCode error code.
func PermissionDenied(message string, errs ...any) *Exception {
	return _createException(PermissionDeniedCode, message, errs...)
}

// Unauthenticated creates a new Exception with the UnauthenticatedCode error code.
func Unauthenticated(message string, errs ...any) *Exception {
	return _createException(UnauthenticatedCode, message, errs...)
}

// Internal creates a new Exception with the ErrorInternalCode error code.
// The original error that caused the exception is also included.
func Internal(message string, errs ...any) *Exception {
	return _createException(InternalErrorCode, message, errs...)
}
