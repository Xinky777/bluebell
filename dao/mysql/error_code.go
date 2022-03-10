package mysql

import "errors"

var (
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorUserExist       = errors.New("用户已存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
)
