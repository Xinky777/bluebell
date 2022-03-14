package controller

import (
	"strconv"
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

//CreatePostHandler 创建帖子的处理函数
func CreatePostHandler(c *gin.Context) {
	//1.获取参数及参数的校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("create post with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从 c 取到当前发请求的用户的ID
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//2.创建帖子
	if err = logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

//GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数（从URL中获取帖子id）
	PostIdStr := c.Param("id")                         //获取URL参数(:id)
	PostId, err := strconv.ParseInt(PostIdStr, 10, 64) //格式转换
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam) //参数错误 格式转换失败
		return
	}
	//2. 根据id去除帖子数据（查数据库）
	data, err := logic.GetPostById(PostId)
	if err != nil {
		zap.L().Error("logic.GetPostById(PostId) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, data)
}

//GetPostListHandler 获取帖子列表借口的处理函数
func GetPostListHandler(c *gin.Context) {
	//获取分页参数
	page, size := getPageInfo(c)
	//获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}
