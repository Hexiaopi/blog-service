package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store"
)

type OperationController struct {
	srv service.Service
}

func NewOperationController(store store.Factory) *OperationController {
	return &OperationController{
		srv: service.NewService(store, nil),
	}
}

// @Summary 获取多个操作日志
// @Description 获取多个操作日志
// @Tags Operation
// @Produce json
// @Security JWT
// @param user query string false "用户"
// @param object query integer false "对象"
// @param page_num query integer false "页码"
// @param page_size query integer false "每页数量"
// @Success 200 {object} app.ListResponse{data=[]model.OperationLog} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/operation [get]
func (c *OperationController) List(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	object := values.Get("object")
	action := values.Get("action")
	userId, _ := strconv.Atoi(values.Get("user_id"))
	page, _ := strconv.Atoi(values.Get("page"))
	limit, _ := strconv.Atoi(values.Get("limit"))
	sort := values.Get("sort")
	param := service.OperationListRequest{ListOption: model.ListOption{UserId: userId, Object: object, Action: action, Limit: limit, Page: page, Sort: sort}}
	logs, total, err := c.srv.Operations().List(request.Context(), &param)
	if err != nil {
		app.ToResponseCode(writer, retcode.GetOperationsFail)
		return
	}
	app.ToResponseList(writer, total, logs)
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
func (c *OperationController) Create(writer http.ResponseWriter, request *http.Request) {
	var param service.CreateOperationRequest
	data, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(writer, retcode.RequestUnMarshalError)
		return
	}
	if err := c.srv.Operations().Create(request.Context(), &param); err != nil {
		app.ToResponseCode(writer, retcode.CreateOperationFail)
		return
	}
	app.ToResponseData(writer, nil)
}

// @Summary 修改操作日志
// @Description 修改操作日志
// @Tags Operation
// @Produce json
// @Accept json
// @Security JWT
// @Param id path integer true "操作日志ID"
// @param UpdateOperationRequest body service.UpdateTagRequest true "修改文章"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/operation [put]
func (c *OperationController) Update(writer http.ResponseWriter, request *http.Request) {
	var param service.UpdateOperationRequest
	data, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(writer, retcode.RequestUnMarshalError)
		return
	}
	if err := c.srv.Operations().Update(request.Context(), &param); err != nil {
		app.ToResponseCode(writer, retcode.UpdateOperationFail)
		return
	}
	app.ToResponseData(writer, nil)
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
func (c *OperationController) Delete(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(request.URL.Query().Get("id"))
	param := service.DeleteOperationRequest{OneOption: model.OneOption{Id: id}}
	if err := c.srv.Operations().Delete(request.Context(), &param); err != nil {
		app.ToResponseCode(writer, retcode.DeleteOperationFail)
		return
	}
	app.ToResponseData(writer, nil)
}
