package goexception

// InvalidArgument creates a new Exception with the InvalidArgumentCode error code.
func InvalidArgument(message string, errs ...any) *Exception {
	return _createException(InvalidArgumentCode, message, errs...)
}

// InvalidInputData creates a new Exception with the InvalidInputDataCode error code.
func InvalidInputData(message string, errs ...any) *Exception {
	return _createException(InvalidInputDataCode, message, errs...)
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
