package routers

import (
	"clients/handlers/cron"
	"clients/handlers/monitor"
	"clients/handlers/work"
	"clients/middleware"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Not Found
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404")
	})

	// pprof router
	pprof.Register(g)

	// Api for authenttication functionalities
	// 登陆
	//g.POST("/login", user.Login)
	//g.POST("/logout", user.Logout)

	apis := g.Group("/api")
	//apis.GET("add_cron", cron.Create)
	//apis.GET("reload_cron", cron.Reload)
	apis.Use(middleware.AuthMiddleware())
	{
		apis.POST("job_create", work.Create) // 创建任务

		apis.GET("reload_cron", cron.Reload) // cron 重载

		apis.GET("active", monitor.Active) // 热数据
		apis.GET("info", monitor.Info)     // 基础数据
	}

	// The health check handlers
	//svcd := g.Group("/test")
	//{
	//	svcd.GET("/health", test.HealthCheck) // 获取健康状态
	//	svcd.GET("/disk", test.DiskCheck)     // 获取硬盘信息
	//	svcd.GET("/cpu", test.CPUCheck)       // 获取cpu信息
	//	svcd.GET("/ram", test.RAMCheck)       // 获取内存信息
	//	svcd.GET("/test", test.APICheck)
	//}

	return g
}
