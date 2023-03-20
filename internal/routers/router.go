package routers

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/hexiaopi/blog-service/docs"
	cache "github.com/hexiaopi/blog-service/internal/cache/redis"
	"github.com/hexiaopi/blog-service/internal/config"
	"github.com/hexiaopi/blog-service/internal/middleware"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/routers/api"
	v1 "github.com/hexiaopi/blog-service/internal/routers/api/v1"
	dao "github.com/hexiaopi/blog-service/internal/store/mysql"
)

func NewRouter() http.Handler {
	storeIns := dao.NewDao(config.DBEngine)
	cacheIns := cache.NewCache(config.RedisEngine)

	router := mux.NewRouter()
	// swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	// profiling
	router.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)
	// metrics
	router.Handle("/metrics", promhttp.Handler()).Methods("GET")
	//router.Use(middleware.Logger)
	router.Use(middleware.Recovery)

	userController := api.NewUserController(storeIns)

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.RequestId)
	apiRouter.Use(middleware.Logger)
	apiRouter.Use(middleware.Metrics)
	apiRouter.Use(middleware.Timeout)
	apiRouter.Use(middleware.Tracer)

	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", userController.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/logout", api.Logout).Methods(http.MethodPost)

	apiV1 := apiRouter.PathPrefix("/v1").Subrouter()
	apiV1.Use(middleware.JWT)
	{
		apiV1.HandleFunc("/user", userController.Info).Methods(http.MethodGet)
		articleController := v1.NewArticleController(storeIns, cacheIns)
		apiV1.HandleFunc("/articles", articleController.List).Methods(http.MethodGet)
		apiV1.HandleFunc("/article", articleController.Create).Methods(http.MethodPost)
		apiV1.HandleFunc("/article", articleController.Get).Methods(http.MethodGet)
		apiV1.HandleFunc("/article", articleController.Update).Methods(http.MethodPut)
		apiV1.HandleFunc("/article", articleController.Delete).Methods(http.MethodDelete)
		tagController := v1.NewTagController(storeIns)
		apiV1.HandleFunc("/tags", tagController.List).Methods(http.MethodGet)
		apiV1.HandleFunc("/tag", tagController.Create).Methods(http.MethodPost)
		apiV1.HandleFunc("/tag", tagController.Update).Methods(http.MethodPut)
		apiV1.HandleFunc("/tag", tagController.Delete).Methods(http.MethodDelete)
	}
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/dist"))))

	router.NotFoundHandler = pathNotFound{}
	router.MethodNotAllowedHandler = methodNotAllow{}
	return router
}

type pathNotFound struct{}

func (pathNotFound) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Errorf("request path:%s not found!", request.RequestURI)
	writer.Write(retcode.RequestPathNotFound.Marshal())
}

type methodNotAllow struct{}

func (methodNotAllow) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Errorf("request path:%s method:%s not allowed!", request.RequestURI, request.Method)
	writer.Write(retcode.RequestMethodNotAllow.Marshal())
}
