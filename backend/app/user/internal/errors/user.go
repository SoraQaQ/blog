package errors

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// 用户验证错误码
const (
	ErrorReasonUserNotFound    = "USER_NOT_FOUND"
	ErrorReasonUserExist       = "USER_ALREADY_EXISTS"
	ErrorReasonInvalidUsername = "INVALID_USERNAME"
	ErrorReasonInvalidPassword = "INVALID_PASSWORD"
	ErrorReasonInvalidEmail    = "INVALID_EMAIL"
)

// 错误定义
var (
	ErrUserNotFound    = errors.New(404, ErrorReasonUserNotFound, "用户不存在")
	ErrUserExists      = errors.New(409, ErrorReasonUserExist, "用户已存在")
	ErrInvalidUsername = errors.New(400, ErrorReasonInvalidUsername, "用户名不能为空")
	ErrInvalidPassword = errors.New(400, ErrorReasonInvalidPassword, "密码不能为空或格式不正确")
	ErrInvalidEmail    = errors.New(400, ErrorReasonInvalidEmail, "邮箱格式不正确，请输入有效的邮箱地址")
)
