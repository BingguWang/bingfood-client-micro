// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 为某个枚举单独设置错误码
func IsInvalidArgument(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_ARGUMENT.String() && e.Code == 400
}

// 为某个枚举单独设置错误码
func ErrorInvalidArgument(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_INVALID_ARGUMENT.String(), fmt.Sprintf(format, args...))
}

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsUnauthenticated(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UNAUTHENTICATED.String() && e.Code == 401
}

func ErrorUnauthenticated(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_UNAUTHENTICATED.String(), fmt.Sprintf(format, args...))
}

func IsPasswordFalse(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PASSWORD_FALSE.String() && e.Code == 405
}

func ErrorPasswordFalse(format string, args ...interface{}) *errors.Error {
	return errors.New(405, ErrorReason_PASSWORD_FALSE.String(), fmt.Sprintf(format, args...))
}

// 服务器内部错误
func IsInternal(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INTERNAL.String() && e.Code == 500
}

// 服务器内部错误
func ErrorInternal(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_INTERNAL.String(), fmt.Sprintf(format, args...))
}