package router

import (
	"net/http"
	"time"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

//Setup 注册路由
func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //gin设置为发布模式
	}
	r := gin.New()
	//r.Use()使用log日志收集组件
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	//注册业务路由
	v1.POST("/signup", controller.SignUpHandler)

	//登陆业务路由
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware(), middlewares.RateLimitMiddleware(2*time.Second, 1)) //应用JWT认证中间件以及令牌桶限流组件
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		//根据时间或分数获取帖子列表
		v1.GET("/post2", controller.GetPostListHandler2)
		v1.POST("/vote", controller.PostVoteController)
	}
	pprof.Register(r) //注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
