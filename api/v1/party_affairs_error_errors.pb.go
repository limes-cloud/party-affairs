// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotFoundError.String() && e.Code == 200
}

func NotFoundErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_NotFoundError.String(), "数据不存在:"+fmt.Sprintf(format, args...))
}

func NotFoundError() *errors.Error {
	return errors.New(200, ErrorReason_NotFoundError.String(), "数据不存在")
}

func IsDatabaseError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DatabaseError.String() && e.Code == 200
}

func DatabaseErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_DatabaseError.String(), "数据库错误:"+fmt.Sprintf(format, args...))
}

func DatabaseError() *errors.Error {
	return errors.New(200, ErrorReason_DatabaseError.String(), "数据库错误")
}

func IsMetadataError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_MetadataError.String() && e.Code == 200
}

func MetadataErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_MetadataError.String(), "元数据异常:"+fmt.Sprintf(format, args...))
}

func MetadataError() *errors.Error {
	return errors.New(200, ErrorReason_MetadataError.String(), "元数据异常")
}

func IsTransformError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_TransformError.String() && e.Code == 200
}

func TransformErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_TransformError.String(), "数据转换失败:"+fmt.Sprintf(format, args...))
}

func TransformError() *errors.Error {
	return errors.New(200, ErrorReason_TransformError.String(), "数据转换失败")
}

func IsDataError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DataError.String() && e.Code == 200
}

func DataErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_DataError.String(), "数据错误:"+fmt.Sprintf(format, args...))
}

func DataError() *errors.Error {
	return errors.New(200, ErrorReason_DataError.String(), "数据错误")
}

func IsResourceServiceError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ResourceServiceError.String() && e.Code == 200
}

func ResourceServiceErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_ResourceServiceError.String(), "资源服务异常:"+fmt.Sprintf(format, args...))
}

func ResourceServiceError() *errors.Error {
	return errors.New(200, ErrorReason_ResourceServiceError.String(), "资源服务异常")
}

func IsGetAuthInfoError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GetAuthInfoError.String() && e.Code == 200
}

func GetAuthInfoErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_GetAuthInfoError.String(), "获取授权信息失败:"+fmt.Sprintf(format, args...))
}

func GetAuthInfoError() *errors.Error {
	return errors.New(200, ErrorReason_GetAuthInfoError.String(), "获取授权信息失败")
}

func IsUnBindError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UnBindError.String() && e.Code == 401
}

func UnBindErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(401, ErrorReason_UnBindError.String(), "用户未绑定:"+fmt.Sprintf(format, args...))
}

func UnBindError() *errors.Error {
	return errors.New(401, ErrorReason_UnBindError.String(), "用户未绑定")
}

func IsBindError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_BindError.String() && e.Code == 401
}

func BindErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(401, ErrorReason_BindError.String(), "用户未绑定:"+fmt.Sprintf(format, args...))
}

func BindError() *errors.Error {
	return errors.New(401, ErrorReason_BindError.String(), "用户未绑定")
}

func IsLoginError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_LoginError.String() && e.Code == 200
}

func LoginErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_LoginError.String(), "登录失败:"+fmt.Sprintf(format, args...))
}

func LoginError() *errors.Error {
	return errors.New(200, ErrorReason_LoginError.String(), "登录失败")
}

func IsAlreadyBindError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_AlreadyBindError.String() && e.Code == 200
}

func AlreadyBindErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_AlreadyBindError.String(), "已存在绑定信息:"+fmt.Sprintf(format, args...))
}

func AlreadyBindError() *errors.Error {
	return errors.New(200, ErrorReason_AlreadyBindError.String(), "已存在绑定信息")
}

func IsUsernameFormatError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UsernameFormatError.String() && e.Code == 200
}

func UsernameFormatErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_UsernameFormatError.String(), "用户名格式错误:"+fmt.Sprintf(format, args...))
}

func UsernameFormatError() *errors.Error {
	return errors.New(200, ErrorReason_UsernameFormatError.String(), "用户名格式错误")
}

func IsCaptchaError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CaptchaError.String() && e.Code == 200
}

func CaptchaErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_CaptchaError.String(), "验证码错误:"+fmt.Sprintf(format, args...))
}

func CaptchaError() *errors.Error {
	return errors.New(200, ErrorReason_CaptchaError.String(), "验证码错误")
}

func IsUsernameNotExistError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UsernameNotExistError.String() && e.Code == 200
}

func UsernameNotExistErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_UsernameNotExistError.String(), "账户信息不存在:"+fmt.Sprintf(format, args...))
}

func UsernameNotExistError() *errors.Error {
	return errors.New(200, ErrorReason_UsernameNotExistError.String(), "账户信息不存在")
}

func IsSendEmailCaptchaError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SendEmailCaptchaError.String() && e.Code == 200
}

func SendEmailCaptchaErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(200, ErrorReason_SendEmailCaptchaError.String(), "发送验证码失败:"+fmt.Sprintf(format, args...))
}

func SendEmailCaptchaError() *errors.Error {
	return errors.New(200, ErrorReason_SendEmailCaptchaError.String(), "发送验证码失败")
}

func IsRefreshTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RefreshTokenError.String() && e.Code == 401
}

func RefreshTokenErrorFormat(format string, args ...any) *errors.Error {
	return errors.New(401, ErrorReason_RefreshTokenError.String(), "刷新token失败:"+fmt.Sprintf(format, args...))
}

func RefreshTokenError() *errors.Error {
	return errors.New(401, ErrorReason_RefreshTokenError.String(), "刷新token失败")
}
