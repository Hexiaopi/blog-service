package v1

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store"
)

type ResourceController struct {
	srv service.Service
}

func NewResourceController(store store.Factory) *ResourceController {
	return &ResourceController{
		srv: service.NewService(store, nil),
	}
}

// @Summary 获取多个资源
// @Description 获取多个资源
// @Tags Resource
// @Produce json
// @Security JWT
// @param name query string false "标签名称"
// @param state query integer false "状态"
// @param page_num query integer false "页码"
// @param page_size query integer false "每页数量"
// @Success 200 {object} app.ListResponse{data=[]model.Tag} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/resources [get]
func (c *ResourceController) List(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	name := values.Get("name")
	state, _ := strconv.Atoi(values.Get("state"))
	page, _ := strconv.Atoi(values.Get("page"))
	limit, _ := strconv.Atoi(values.Get("limit"))
	sort := values.Get("sort")
	param := service.ResourceListRequest{ListOption: model.ListOption{Name: name, State: uint8(state), Limit: limit, Page: page, Sort: sort}}
	resources, total, err := c.srv.Resources().List(request.Context(), &param)
	if err != nil {
		app.ToResponseCode(writer, retcode.GetResourcesFail)
		return
	}
	app.ToResponseList(writer, total, resources)
}

// @Summary 获取单个资源
// @Description 获取单个资源详细信息
// @Tags Resource
// @Accept  json
// @Produce  json
// @Security JWT
// @Param id path integer true "资源ID"
// @Success 200 {object} app.CommResponse{data=model.Resource} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/resource/{id} [get]
func (c *ResourceController) Get(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(request.URL.Query().Get("id"))
	param := service.ResourceRequest{OneOption: model.OneOption{Id: id}}
	resource, err := c.srv.Resources().Get(request.Context(), &param)
	if err != nil {
		app.ToResponseCode(writer, retcode.GetResourceFail)
		return
	}
	app.ToResponseData(writer, resource)
}

// @Summary 创建资源
// @Description 创建资源
// @Tags Resource
// @Produce form-data
// @Accept json
// @Security JWT
// @param CreateResourceRequest body service.CreateResourceRequest true "创建标签"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/resource [post]
func (c *ResourceController) Create(writer http.ResponseWriter, request *http.Request) {
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		app.ToResponseCode(writer, retcode.RequestIllegal)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		app.ToResponseCode(writer, retcode.RequestIllegal)
		return
	}
	stateValue := request.FormValue("state")
	state, _ := strconv.Atoi(stateValue)
	param := service.CreateResourceRequest{
		Resource: model.Resource{
			Name:  fileHeader.Filename,
			Blob:  data,
			Type:  fileHeader.Header.Get("Content-Type"),
			Size:  fileHeader.Size,
			State: uint8(state),
		},
	}
	if err := c.srv.Resources().Create(request.Context(), &param); err != nil {
		app.ToResponseCode(writer, retcode.CreateResourceFail)
		return
	}
	app.ToResponseData(writer, nil)
}

// @Summary 修改资源
// @Description 修改资源
// @Tags Resource
// @Produce json
// @Accept json
// @Security JWT
// @Param id path integer true "标签ID"
// @param UpdateResourceRequest body service.UpdateResourceRequest true "修改资源"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/resource [put]
func (c *ResourceController) Update(writer http.ResponseWriter, request *http.Request) {
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		app.ToResponseCode(writer, retcode.RequestIllegal)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		app.ToResponseCode(writer, retcode.RequestIllegal)
		return
	}
	id, _ := strconv.Atoi(request.URL.Query().Get("id"))
	stateValue := request.FormValue("state")
	state, _ := strconv.Atoi(stateValue)
	param := service.UpdateResourceRequest{
		Resource: model.Resource{
			ID:    id,
			Name:  fileHeader.Filename,
			Blob:  data,
			Type:  fileHeader.Header.Get("Content-Type"),
			Size:  fileHeader.Size,
			State: uint8(state),
		},
	}
	if err := c.srv.Resources().Update(request.Context(), &param); err != nil {
		app.ToResponseCode(writer, retcode.UpdateResourceFail)
		return
	}
	app.ToResponseData(writer, nil)
}

// @Summary 删除资源
// @Description 删除资源
// @Tags Resource
// @Produce json
// @Security JWT
// @Param id path integer true "标签ID"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/resource [delete]
func (c *ResourceController) Delete(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(request.URL.Query().Get("id"))
	if err := c.srv.Resources().Delete(request.Context(), id); err != nil {
		app.ToResponseCode(writer, retcode.DeleteResourceFail)
		return
	}
	app.ToResponseData(writer, nil)
}
