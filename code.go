package goexception

type Code string

const (
	// Deprecated: Use InvalidParameterCode instead
	InvalidArgumentCode  Code = "INVALID_ARGUMENT"
	InvalidParameterCode Code = "INVALID_PARAMETER"
	InvalidDataCode      Code = "INVALID_DATA"
	NotFoundCode         Code = "NOT_FOUND"
	AlreadyExistsCode    Code = "ALREADY_EXISTS"
	PermissionDeniedCode Code = "PERMISSION_DENIED"
	UnauthenticatedCode  Code = "UNAUTHENTICATED"
	InternalErrorCode    Code = "INTERNAL_ERROR"
)

func (e Code) ToString() string {
	return string(e)
}
