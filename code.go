package goexception

type Code string

const (
	InvalidArgumentCode  Code = "INVALID_ARGUMENT"
	InvalidInputDataCode Code = "INVALID_INPUT_DATA"
	NotFoundCode         Code = "NOT_FOUND"
	AlreadyExistsCode    Code = "ALREADY_EXISTS"
	PermissionDeniedCode Code = "PERMISSION_DENIED"
	UnauthenticatedCode  Code = "UNAUTHENTICATED"
	InternalErrorCode    Code = "INTERNAL_ERROR"
)

func (e Code) ToString() string {
	return string(e)
}
