package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

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

//GetPostListHandler 获取帖子列表接口的处理函数
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

//GetPostListHandler2 升级版获取帖子列表接口
//根据前端传里的参数动态获取帖子列表
//根据创建时间排序 或者按照分数排序
func GetPostListHandler2(c *gin.Context) {
	//1.获取参数
	//get请求参数:  query(string)  /api/v1/post2?page=1&size=10&order=time
	//c.ShouldBind() 根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中是json格式的数据 才能用这个方法获取到数据

	//初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//获取数据
	//2.去redis查询id列表 -> 3.根据id去数据库查询帖子详细信息
	data, err := logic.GetPostListNew(p) //更新：查询接口合二为一
	if err != nil {
		zap.L().Error("logic.GetPostList2 failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

//GetCommunityPostListHandler 根据社区查询帖子列表
//func GetCommunityPostListHandler(c *gin.Context) {
//	//初始化结构体时指定初始参数
//	p := &models.ParamCommunityPostList{
//		ParamPostList: &models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//	}
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//	//获取数据
//	data, err := logic.GetCommunityPostList(p) //2.去redis查询id列表 -> 3.根据id去数据库查询帖子详细信息
//	if err != nil {
//		zap.L().Error("logic.GetPostList2 failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	ResponseSuccess(c, data)
//}
