package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type OperationController struct {
	srv    service.Service
	logger log.Logger
}

func NewOperationController(store store.Factory, logger log.Logger) *OperationController {
	return &OperationController{
		srv:    service.NewService(store, nil, logger),
		logger: logger,
	}
}

// @Summary 获取多个操作日志
// @Description 获取多个操作日志
// @Tags Operation
// @Produce json
// @Security JWT
// @param username query string false "用户名"
// @param action query string false "操作"
// @param object query string false "对象"
// @param result query string false "结果"
// @param sort query string false "排序方式"
// @param page query integer false "页码"
// @param limit query integer false "页面大小"
// @Success 200 {object} app.ListResponse{data=[]model.OperationLog} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/operation [get]
func (c *OperationController) List(ctx *gin.Context) (res interface{}, total int64, err error) {
	values := ctx.Request.URL.Query()
	username := values.Get("username")
	object := values.Get("object")
	action := values.Get("action")
	result := values.Get("result")
	page, _ := strconv.Atoi(values.Get("page"))
	limit, _ := strconv.Atoi(values.Get("limit"))
	sort := values.Get("sort")
	param := service.OperationListRequest{UserName: username, ListOption: entity.ListOption{Object: object, Action: action, Result: result, Limit: limit, Page: page, Sort: sort}}
	logs, total, err := c.srv.Operations().List(ctx.Request.Context(), &param)
	if err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.GetOperationsFail)
		return nil, 0, retcode.GetOperationsFail
	}
	//app.ToResponseList(ctx.Writer, total, logs)
	return logs, total, nil
}
