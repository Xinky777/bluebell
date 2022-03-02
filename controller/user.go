package controller

import (
	"net/http"
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
)

//SignUpHandler 处理注册请求
func SignUpHandler(c *gin.Context) {
	//1.获取参数并参数校验
	var p models.ParamSignUp
	c.ShouldBindJSON(&p)
	//2.业务处理
	logic.SingUp()
	//3.返回响应
	c.JSON(http.StatusOK, "ok")
}
