// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errors

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_NotFound.String() && e.Code == 200
}

func NotFoundFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_NotFound.String(), "数据不存在:"+fmt.Sprintf(format, args...))
}

func NotFound() *errors.Error {
	return errors.New(200, Reason_NotFound.String(), "数据不存在")
}

func IsDatabase(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_Database.String() && e.Code == 200
}

func DatabaseFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_Database.String(), "数据库错误:"+fmt.Sprintf(format, args...))
}

func Database() *errors.Error {
	return errors.New(200, Reason_Database.String(), "数据库错误")
}

func IsMetadata(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_Metadata.String() && e.Code == 200
}

func MetadataFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_Metadata.String(), "元数据异常:"+fmt.Sprintf(format, args...))
}

func Metadata() *errors.Error {
	return errors.New(200, Reason_Metadata.String(), "元数据异常")
}

func IsTransform(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_Transform.String() && e.Code == 200
}

func TransformFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_Transform.String(), "数据转换失败:"+fmt.Sprintf(format, args...))
}

func Transform() *errors.Error {
	return errors.New(200, Reason_Transform.String(), "数据转换失败")
}

func IsData(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_Data.String() && e.Code == 200
}

func DataFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_Data.String(), "数据错误:"+fmt.Sprintf(format, args...))
}

func Data() *errors.Error {
	return errors.New(200, Reason_Data.String(), "数据错误")
}

func IsResourceService(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_ResourceService.String() && e.Code == 200
}

func ResourceServiceFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_ResourceService.String(), "资源服务异常:"+fmt.Sprintf(format, args...))
}

func ResourceService() *errors.Error {
	return errors.New(200, Reason_ResourceService.String(), "资源服务异常")
}

func IsGetAuthInfo(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_GetAuthInfo.String() && e.Code == 200
}

func GetAuthInfoFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_GetAuthInfo.String(), "获取授权信息失败:"+fmt.Sprintf(format, args...))
}

func GetAuthInfo() *errors.Error {
	return errors.New(200, Reason_GetAuthInfo.String(), "获取授权信息失败")
}

func IsUnBind(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_UnBind.String() && e.Code == 401
}

func UnBindFormat(format string, args ...any) *errors.Error {
	return errors.New(401, Reason_UnBind.String(), "用户未绑定:"+fmt.Sprintf(format, args...))
}

func UnBind() *errors.Error {
	return errors.New(401, Reason_UnBind.String(), "用户未绑定")
}

func IsBind(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_Bind.String() && e.Code == 401
}

func BindFormat(format string, args ...any) *errors.Error {
	return errors.New(401, Reason_Bind.String(), "用户未绑定:"+fmt.Sprintf(format, args...))
}

func Bind() *errors.Error {
	return errors.New(401, Reason_Bind.String(), "用户未绑定")
}

func IsLogin(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_Login.String() && e.Code == 200
}

func LoginFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_Login.String(), "登录失败:"+fmt.Sprintf(format, args...))
}

func Login() *errors.Error {
	return errors.New(200, Reason_Login.String(), "登录失败")
}

func IsAlreadyBind(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_AlreadyBind.String() && e.Code == 200
}

func AlreadyBindFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_AlreadyBind.String(), "已存在绑定信息:"+fmt.Sprintf(format, args...))
}

func AlreadyBind() *errors.Error {
	return errors.New(200, Reason_AlreadyBind.String(), "已存在绑定信息")
}

func IsUsernameFormat(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_UsernameFormat.String() && e.Code == 200
}

func UsernameFormatFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_UsernameFormat.String(), "用户名格式错误:"+fmt.Sprintf(format, args...))
}

func UsernameFormat() *errors.Error {
	return errors.New(200, Reason_UsernameFormat.String(), "用户名格式错误")
}

func IsCaptcha(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_Captcha.String() && e.Code == 200
}

func CaptchaFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_Captcha.String(), "验证码错误:"+fmt.Sprintf(format, args...))
}

func Captcha() *errors.Error {
	return errors.New(200, Reason_Captcha.String(), "验证码错误")
}

func IsUsernameNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_UsernameNotExist.String() && e.Code == 200
}

func UsernameNotExistFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_UsernameNotExist.String(), "账户信息不存在:"+fmt.Sprintf(format, args...))
}

func UsernameNotExist() *errors.Error {
	return errors.New(200, Reason_UsernameNotExist.String(), "账户信息不存在")
}

func IsSendEmailCaptcha(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_SendEmailCaptcha.String() && e.Code == 200
}

func SendEmailCaptchaFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_SendEmailCaptcha.String(), "发送验证码失败:"+fmt.Sprintf(format, args...))
}

func SendEmailCaptcha() *errors.Error {
	return errors.New(200, Reason_SendEmailCaptcha.String(), "发送验证码失败")
}

func IsRefreshToken(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_RefreshToken.String() && e.Code == 401
}

func RefreshTokenFormat(format string, args ...any) *errors.Error {
	return errors.New(401, Reason_RefreshToken.String(), "刷新token失败:"+fmt.Sprintf(format, args...))
}

func RefreshToken() *errors.Error {
	return errors.New(401, Reason_RefreshToken.String(), "刷新token失败")
}

func IsAuthInfo(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_AuthInfo.String() && e.Code == 200
}

func AuthInfoFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_AuthInfo.String(), "授权信息错误:"+fmt.Sprintf(format, args...))
}

func AuthInfo() *errors.Error {
	return errors.New(200, Reason_AuthInfo.String(), "授权信息错误")
}

func IsUserCenter(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_UserCenter.String() && e.Code == 200
}

func UserCenterFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_UserCenter.String(), "用户中心连接失败:"+fmt.Sprintf(format, args...))
}

func UserCenter() *errors.Error {
	return errors.New(200, Reason_UserCenter.String(), "用户中心连接失败")
}

func IsResource(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Reason_Resource.String() && e.Code == 200
}

func ResourceFormat(format string, args ...any) *errors.Error {
	return errors.New(200, Reason_Resource.String(), "资源中心连接失败:"+fmt.Sprintf(format, args...))
}

func Resource() *errors.Error {
	return errors.New(200, Reason_Resource.String(), "资源中心连接失败")
}
