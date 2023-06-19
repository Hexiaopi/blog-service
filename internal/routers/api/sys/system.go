package sys

import (
	"github.com/gin-gonic/gin"
	
	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store"
)

type SystemController struct {
	srv service.Service
}

func NewSystemController(store store.Factory) *SystemController {
	return &SystemController{
		srv: service.NewService(store, nil),
	}
}

// @Summary 获取系统配置项
// @Description 获取系统配置项，一般是key-value格式
// @Tags System
// @Produce json
// @param name query string false "名称"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/sys/config [get]
func (c *SystemController) Get(ctx *gin.Context) {
	values := ctx.Request.URL.Query()
	name := values.Get("name")
	param := service.SystemGetRequest{OneOption: model.OneOption{Name: name}}
	config, err := c.srv.Systems().Get(ctx.Request.Context(), &param)
	if err != nil {
		app.ToResponseCode(ctx.Writer, retcode.GetSystemConfigFail)
		return
	}
	app.ToResponseData(ctx.Writer, config.Value)
}
