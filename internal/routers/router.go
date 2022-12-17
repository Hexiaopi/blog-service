package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/hexiaopi/blog-service/docs"

	"github.com/hexiaopi/blog-service/internal/config"
	"github.com/hexiaopi/blog-service/internal/middleware"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/routers/api"
	v1 "github.com/hexiaopi/blog-service/internal/routers/api/v1"
	"github.com/hexiaopi/blog-service/internal/store/dao"
)

func NewRouter() http.Handler {
	storeIns := dao.NewDao(config.DBEngine)

	router := mux.NewRouter()
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	router.PathPrefix("/drawio/").Handler(http.StripPrefix("/drawio/", http.FileServer(http.Dir("./drawio/src/main/webapp"))))
	router.Use(middleware.Logger)
	router.Use(middleware.Recovery)
	router.Use(middleware.Tracer)
	userController := api.NewUserController(storeIns)
	router.HandleFunc("/user/login", userController.Login).Methods(http.MethodPost)
	router.HandleFunc("/user/info", api.Info).Methods(http.MethodGet)
	router.HandleFunc("/user/logout", api.Logout).Methods(http.MethodPost)

	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.Timeout)
	apiV1.Use(middleware.JWT)
	{
		articleController := v1.NewArticleController(storeIns)
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
