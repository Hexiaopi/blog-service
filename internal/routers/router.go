package routers

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/hexiaopi/blog-service/internal/app"
	cache "github.com/hexiaopi/blog-service/internal/cache/redis"
	"github.com/hexiaopi/blog-service/internal/config"
	"github.com/hexiaopi/blog-service/internal/middleware"
	"github.com/hexiaopi/blog-service/internal/retcode"
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
	router.StaticFS("/static", http.Dir("./web/dist"))
	//router.Use(middleware.Logger)
	router.NoMethod(MethodNotAllow)
	router.NoRoute(PathNotFound)
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
	sysRouter.GET("/config", systemController.Get)
	sysRouter.GET("/captcha", captchaController.Get)
	sysRouter.POST("/login", loginController.Login)
	sysRouter.POST("/logout", loginController.Logout)

	apiV1 := apiRouter.Group("/v1")
	apiV1.Use(middleware.JWT())
	apiV1.Use(middleware.NewOperation(storeIns).RecordOperation())
	{
		apiV1.GET("/user", loginController.Info)
		articleController := v1.NewArticleController(storeIns, cacheIns)
		apiV1.GET("/articles", articleController.List)
		apiV1.POST("/article", articleController.Create)
		apiV1.GET("/article", articleController.Get)
		apiV1.PUT("/article", articleController.Update)
		apiV1.DELETE("/article", articleController.Delete)
		tagController := v1.NewTagController(storeIns)
		apiV1.GET("/tags", tagController.List)
		apiV1.POST("/tag", tagController.Create)
		apiV1.PUT("/tag", tagController.Update)
		apiV1.DELETE("/tag", tagController.Delete)
		resourceController := v1.NewResourceController(storeIns)
		apiV1.GET("/resources", resourceController.List)
		apiV1.POST("/resource", resourceController.Create)
		apiV1.GET("/resource", resourceController.Get)
		apiV1.PUT("/resource", resourceController.Update)
		apiV1.DELETE("/resource", resourceController.Delete)
		operationController := v1.NewOperationController(storeIns)
		apiV1.GET("/operations", operationController.List)
		apiV1.POST("/operation", operationController.Create)
		apiV1.PUT("/operation", operationController.Update)
		apiV1.DELETE("/operation", operationController.Delete)
	}

	//router.NotFoundHandler = pathNotFound{}
	//router.MethodNotAllowedHandler = methodNotAllow{}
	return router
}

func PathNotFound(ctx *gin.Context) {
	log.Errorf("request path:%s not found!", ctx.Request.RequestURI)
	app.ToResponseCode(ctx.Writer, retcode.RequestPathNotFound)
}

func MethodNotAllow(ctx *gin.Context) {
	log.Errorf("request path:%s method:%s not allowed!", ctx.Request.RequestURI, ctx.Request.Method)
	app.ToResponseCode(ctx.Writer, retcode.RequestMethodNotAllow)
}
