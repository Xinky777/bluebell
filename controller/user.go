package controller

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

//SignUpHandler 处理注册请求
func SignUpHandler(c *gin.Context) {
	//1.获取参数并参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误 记录日志
		zap.L().Error("signUp with invalid param", zap.Error(err))
		//判断error是不是validator.ValidationErrors类型的
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//如果不是
			ResponseError(c, CodeInvalidParam)
			return
		}
		//如果是
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) //翻译错误
		return
	}

	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

//LoginHandler 处理登陆请求
func LoginHandler(c *gin.Context) {
	//1.获取参数并校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误 记录日志
		zap.L().Error("login with invalid param", zap.Error(err))
		//判断error是不是validator.ValidationErrors类型的
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//如果不是
			ResponseError(c, CodeInvalidParam)
			return
		}
		//如果是
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.业务处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			//用户已存在
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3.返回响应
	ResponseSuccess(c, CodeSuccess)
}
