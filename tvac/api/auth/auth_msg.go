package auth

import (
	"tvac/web"
)

const (
	errUserNotFound = web.Unauthorized("User not found")
	errPassword     = web.Unauthorized("Mật khẩu không đúng")
	errEmailExites  = web.Unauthorized("Email đã tồn tại")
	registerSuccess = web.Unauthorized("Đăng kí thành công")
)
