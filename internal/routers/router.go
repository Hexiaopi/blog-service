package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/hexiaopi/blog-service/docs"

	"github.com/hexiaopi/blog-service/internal/middleware"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/routers/api"
	v1 "github.com/hexiaopi/blog-service/internal/routers/api/v1"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	router.Use(middleware.Logger)
	router.Use(middleware.Recovery)
	router.Use(middleware.Tracer)
	router.HandleFunc("/user/login", api.Login).Methods(http.MethodPost)
	router.HandleFunc("/user/info", api.Info).Methods(http.MethodGet)
	router.HandleFunc("/user/logout", api.Logout).Methods(http.MethodPost)
	router.HandleFunc("/table/list", api.TableList).Methods(http.MethodGet)

	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.Timeout)
	apiV1.Use(middleware.JWT)
	{
		apiV1.HandleFunc("/articles", v1.ListArticle).Methods(http.MethodGet)
		apiV1.HandleFunc("/articles", v1.CreateArticle).Methods(http.MethodPost)
		apiV1.HandleFunc("/articles/{id}", v1.GetArticle).Methods(http.MethodGet)
		apiV1.HandleFunc("/articles/{id}", v1.UpdateArticle).Methods(http.MethodPut)
		apiV1.HandleFunc("/articles/{id}", v1.DeleteArticle).Methods(http.MethodDelete)
		apiV1.HandleFunc("/tags", v1.ListTag).Methods(http.MethodGet)
		apiV1.HandleFunc("/tags", v1.CreateTag).Methods(http.MethodPost)
		apiV1.HandleFunc("/tags/{id}", v1.UpdateTag).Methods(http.MethodPut)
		apiV1.HandleFunc("/tags/{id}", v1.DeleteTag).Methods(http.MethodDelete)
	}
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
