package routers

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	cache "github.com/hexiaopi/blog-service/internal/cache/redis"
	"github.com/hexiaopi/blog-service/internal/config"
	"github.com/hexiaopi/blog-service/internal/middleware"
	_ "github.com/hexiaopi/blog-service/internal/routers/api/docs"
	"github.com/hexiaopi/blog-service/internal/routers/api/sys"
	v1 "github.com/hexiaopi/blog-service/internal/routers/api/v1"
	dao "github.com/hexiaopi/blog-service/internal/store/mysql"
)

func NewRouter() *gin.Engine {
	storeIns := dao.NewDao(config.DBEngine)
	cacheIns := cache.NewCache(config.RedisEngine)

	router := gin.New()
	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// profiling
	router.GET("/debug/pprof/*any", gin.WrapH(http.DefaultServeMux))
	// metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// static
	router.StaticFS("/static", http.Dir("./web/dist"))
	// 全局错误处理
	router.NoMethod(middleware.MethodNotAllow)
	router.NoRoute(middleware.PathNotFound)
	router.Use(middleware.Recovery())

	loginController := sys.NewLoginController(storeIns)
	systemController := sys.NewSystemController(storeIns)
	captchaController := sys.NewCaptchaController()

	apiRouter := router.Group("/api")
	apiRouter.Use(middleware.RequestId())
	apiRouter.Use(middleware.Logger())
	apiRouter.Use(middleware.Metrics())
	apiRouter.Use(middleware.Timeout())
	apiRouter.Use(middleware.Tracer())

	sysRouter := apiRouter.Group("/sys")
	{
		sysRouter.GET("/config", systemController.Get)
		sysRouter.GET("/captcha", captchaController.Get)
		sysRouter.POST("/login", loginController.Login)
		sysRouter.POST("/logout", loginController.Logout)
	}

	apiV1 := apiRouter.Group("/v1")
	apiV1.Use(middleware.JWT())
	apiV1.Use(middleware.NewOperation(storeIns).RecordOperation())
	{
		apiV1.GET("/user", loginController.Info)
		userController := v1.NewUserController(storeIns)
		{
			apiV1.GET("/users", userController.List)
			apiV1.POST("/user", userController.Create)
			apiV1.PUT("/user", userController.Update)
			apiV1.DELETE("/user", userController.Delete)
		}

		articleController := v1.NewArticleController(storeIns, cacheIns)
		{
			apiV1.GET("/articles", articleController.List)
			apiV1.POST("/article", articleController.Create)
			apiV1.GET("/article", articleController.Get)
			apiV1.PUT("/article", articleController.Update)
			apiV1.DELETE("/article", articleController.Delete)
		}
		tagController := v1.NewTagController(storeIns)
		{
			apiV1.GET("/tags", tagController.List)
			apiV1.POST("/tag", tagController.Create)
			apiV1.PUT("/tag", tagController.Update)
			apiV1.DELETE("/tag", tagController.Delete)
		}
		roleController := v1.NewRoleController(storeIns)
		{
			apiV1.GET("/roles", roleController.List)
			apiV1.POST("/role", roleController.Create)
			apiV1.PUT("/role", roleController.Update)
			apiV1.DELETE("/role", roleController.Delete)
		}
		resourceController := v1.NewResourceController(storeIns)
		{
			apiV1.GET("/resources", resourceController.List)
			apiV1.POST("/resource", resourceController.Create)
			apiV1.GET("/resource", resourceController.Get)
			apiV1.PUT("/resource", resourceController.Update)
			apiV1.DELETE("/resource", resourceController.Delete)
		}
		operationController := v1.NewOperationController(storeIns)
		{
			apiV1.GET("/operations", operationController.List)
			apiV1.POST("/operation", operationController.Create)
			apiV1.PUT("/operation", operationController.Update)
			apiV1.DELETE("/operation", operationController.Delete)
		}
	}
	return router
}
