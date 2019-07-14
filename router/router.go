package router

import (
	"apiserver/handler/api"
	"apiserver/pkg/upload"
	"net/http"

	_ "apiserver/docs" // docs is generated by Swag CLI, you have to import it.
	"apiserver/handler/api/sd"
	"apiserver/handler/api/v1/user"
	"apiserver/router/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 中间件.
	g.Use(gin.Recovery())     // 恢复Api server服务器
	g.Use(middleware.NoCache) // 强制浏览器不适用缓存
	g.Use(middleware.Options) // 浏览器跨域OPTIONS请求设置
	g.Use(middleware.Secure)  // 一些安全设置
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	// api for authentication functionalities
	g.POST("/login", user.Login)

	// 图片上传
	g.POST("/upload/image", api.UplaodImage)

	// 前端预览
	g.StaticFS("/upload/preview_images", http.Dir(upload.GetImageFullPath()))

	// 视频上传
	g.POST("/upload/video", api.UploadVideo)

	// 前端预览
	g.StaticFS("/upload/preview_videos", http.Dir(upload.GetVideoFullPath()))

	// 用户路由设置
	u := g.Group("v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)       // 创建用户
		u.DELETE("/:id", user.Delete) // 删除用户
		u.PUT("/:id", user.Update)    // 更新用户
		u.GET("", user.List)          // 用户列表
		u.GET("/:username", user.Get) // 获取指定用户的详细信息
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
