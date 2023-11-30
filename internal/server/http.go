package server

import (
	"context"

	"github.com/hexiaopi/blog-service/internal/config"
	"github.com/hexiaopi/blog-service/internal/routers"
	"github.com/hexiaopi/blog-service/pkg/app"
	"github.com/hexiaopi/blog-service/pkg/server/http"
)

func Run() {
	router := routers.NewRouter()
	httpServer := http.NewServer(router,
		config.Logger,
		http.WithServerHost(config.AppEngine.HTTP.Host),
		http.WithServerPort(config.AppEngine.HTTP.Port),
	)
	if err := app.NewApp(app.WithServer(httpServer)).Run(context.Background()); err != nil {
		config.Logger.Errorf("start http server err:%v", err)
	}
}
