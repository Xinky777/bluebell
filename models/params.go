package models

//定义用户请求的参数结构体
//用户输入的结构体字段

//ParamSignUp 定义注册请求（Signup）的参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

//ParamLogin 定义登陆请求（Login）的参数结构体
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
