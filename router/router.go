package router

import (
	"net/http"

	"apiserver/handler/sd"
	"apiserver/handler/user"
	"apiserver/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 中间件.
	g.Use(gin.Recovery())     // 恢复Api server服务器
	g.Use(middleware.NoCache) // 强制浏览器不适用缓存
	g.Use(middleware.Options) // 浏览器跨域OPTIONS请求设置
	g.Use(middleware.Secure)  // 一些安全设置
	g.Use(mw...)
	// api for authentication functionalities
	g.POST("/login",user.Login)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// 用户路由设置
	u := g.Group("v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("",user.Create) // 创建用户
		u.DELETE("/:id",user.Delete) // 删除用户
		u.PUT("/:id",user.Update) // 更新用户
		u.GET("",user.List) // 用户列表
		u.GET("/:username",user.Get) // 获取指定用户的详细信息
		// u.POST("/:username",user.Create)
	}
	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
