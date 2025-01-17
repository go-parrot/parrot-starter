package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" //nolint: goimports

	_ "github.com/go-parrot/parrot-starter/docs"
	"github.com/go-parrot/parrot-starter/internal/handler"
	"github.com/go-parrot/parrot/pkg/app"
	"github.com/go-parrot/parrot/pkg/middleware"
)

// NewRouter loads the middlewares, routes, handlers.
func NewRouter() *gin.Engine {
	g := gin.New()
	// 使用中间件
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.Logging())
	g.Use(middleware.RequestID())
	g.Use(middleware.Metrics(app.Conf.Name))
	g.Use(middleware.Tracing(app.Conf.Name))

	// 404 Handler.
	g.NoRoute(app.RouteNotFound)
	g.NoMethod(app.RouteNotFound)

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))
	// pprof router 性能分析路由
	// 默认关闭，开发环境下可以打开
	// 访问方式: HOST/debug/pprof
	// 通过 HOST/debug/pprof/profile 生成profile
	// 查看分析图 go tool pprof -http=:5000 profile
	// see: https://github.com/gin-contrib/pprof
	// pprof.Register(g)

	// HealthCheck 健康检查路由
	g.GET("/health", app.HealthCheck)
	// metrics router 可以在 prometheus 中进行监控
	// 通过 grafana 可视化查看 prometheus 的监控数据，使用插件6671查看
	g.GET("/metrics", gin.WrapH(promhttp.Handler()))
	g.GET("/hello", handler.Hello)

	// v1 router
	apiV1 := g.Group("/v1")
	apiV1.Use()
	{
		// here to add biz router
	}

	return g
}
