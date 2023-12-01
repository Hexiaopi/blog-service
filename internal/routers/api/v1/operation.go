package v1

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/model"
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
	param := service.OperationListRequest{UserName: username, ListOption: model.ListOption{Object: object, Action: action, Result: result, Limit: limit, Page: page, Sort: sort}}
	logs, total, err := c.srv.Operations().List(ctx.Request.Context(), &param)
	if err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.GetOperationsFail)
		return nil, 0, retcode.GetOperationsFail
	}
	//app.ToResponseList(ctx.Writer, total, logs)
	return logs, total, nil
}

// @Summary 创建操作日志
// @Description 创建操作日志
// @Tags Operation
// @Produce json
// @Accept json
// @Security JWT
// @param CreateOperationRequest body service.CreateOperationRequest true "创建标签"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/operation [post]
func (c *OperationController) Create(ctx *gin.Context) {
	var param service.CreateOperationRequest
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return
	}
	if err := c.srv.Operations().Create(ctx.Request.Context(), &param); err != nil {
		app.ToResponseCode(ctx.Writer, retcode.CreateOperationFail)
		return
	}
	app.ToResponseData(ctx.Writer, nil)
}

// @Summary 修改操作日志
// @Description 修改操作日志
// @Tags Operation
// @Produce json
// @Accept json
// @Security JWT
// @Param id path integer true "操作日志ID"
// @param UpdateOperationRequest body service.UpdateOperationRequest true "修改文章"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/operation [put]
func (c *OperationController) Update(ctx *gin.Context) {
	var param service.UpdateOperationRequest
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return
	}
	if err := c.srv.Operations().Update(ctx.Request.Context(), &param); err != nil {
		app.ToResponseCode(ctx.Writer, retcode.UpdateOperationFail)
		return
	}
	app.ToResponseData(ctx.Writer, nil)
}

// @Summary 删除操作日志
// @Description 删除操作日志
// @Tags Operation
// @Produce json
// @Security JWT
// @Param id path integer true "操作日志ID"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/operation [delete]
func (c *OperationController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	param := service.DeleteOperationRequest{OneOption: model.OneOption{Id: id}}
	if err := c.srv.Operations().Delete(ctx.Request.Context(), &param); err != nil {
		app.ToResponseCode(ctx.Writer, retcode.DeleteOperationFail)
		return
	}
	app.ToResponseData(ctx.Writer, nil)
}
