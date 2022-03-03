package controller

import (
	"net/http"
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
		//记录日志
		zap.L().Error("signUp with invalid param", zap.Error(err))
		//判断error是不是validator.ValidationErrors类型的
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//如果不是
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		//如果是
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})
		return
	}

	//2.业务处理
	if err := logic.SingUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
