package v1

import (
	"io"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type ResourceController struct {
	srv    service.Service
	logger log.Logger
}

func NewResourceController(store store.Factory, logger log.Logger) *ResourceController {
	return &ResourceController{
		srv:    service.NewService(store, nil, logger),
		logger: logger,
	}
}

// @Summary 获取多个资源
// @Description 获取多个资源
// @Tags Resource
// @Produce json
// @Security JWT
// @param name query string false "标签名称"
// @param state query integer false "状态"
// @param sort query string false "排序方式"
// @param page query integer false "页码"
// @param limit query integer false "页面大小"
// @Success 200 {object} app.ListResponse{data=[]entity.Resource} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/resources [get]
func (c *ResourceController) List(ctx *gin.Context) (res interface{}, total int64, err error) {
	values := ctx.Request.URL.Query()
	name := values.Get("name")
	state, _ := strconv.Atoi(values.Get("state"))
	page, _ := strconv.Atoi(values.Get("page"))
	limit, _ := strconv.Atoi(values.Get("limit"))
	sort := values.Get("sort")
	param := service.ResourceListRequest{ListOption: entity.ListOption{Name: name, State: uint8(state), Limit: limit, Page: page, Sort: sort}}
	resources, total, err := c.srv.Resources().List(ctx.Request.Context(), &param)
	if err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.GetResourcesFail)
		return nil, 0, retcode.GetResourcesFail
	}
	//app.ToResponseList(ctx.Writer, total, resources)
	return resources, total, nil
}

// @Summary 获取单个资源
// @Description 获取单个资源详细信息
// @Tags Resource
// @Accept  json
// @Produce  json
// @Security JWT
// @Param id path integer true "资源ID"
// @Success 200 {object} app.CommResponse{data=entity.Resource} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/resource/{id} [get]
func (c *ResourceController) Get(ctx *gin.Context) (res interface{}, err error) {
	id, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	param := service.ResourceRequest{OneOption: entity.OneOption{Id: id}}
	resource, err := c.srv.Resources().Get(ctx.Request.Context(), &param)
	if err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.GetResourceFail)
		return nil, retcode.GetResourceFail
	}
	//app.ToResponseData(ctx.Writer, resource)
	return resource, nil
}

// @Summary 创建资源
// @Description 创建资源
// @Tags Resource
// @Produce json
// @Accept multipart/form-data
// @Security JWT
// @param file formData file true "file"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/resource [post]
func (c *ResourceController) Create(ctx *gin.Context) (res interface{}, err error) {
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.RequestIllegal)
		return nil, retcode.RequestIllegal
	}
	data, err := io.ReadAll(file)
	if err != nil {
		// app.ToResponseCode(ctx.Writer, retcode.RequestIllegal)
		return nil, retcode.RequestIllegal
	}
	stateValue := ctx.Request.FormValue("state")
	state, _ := strconv.Atoi(stateValue)
	param := service.CreateResourceRequest{
		Resource: entity.Resource{
			Name:  fileHeader.Filename,
			Blob:  data,
			Type:  fileHeader.Header.Get("Content-Type"),
			Size:  fileHeader.Size,
			State: uint8(state),
		},
	}
	if err := c.srv.Resources().Create(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.CreateResourceFail)
		return nil, retcode.CreateResourceFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
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
func (c *ResourceController) Update(ctx *gin.Context) (res interface{}, err error) {
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.RequestIllegal)
		return nil, retcode.RequestIllegal
	}
	data, err := io.ReadAll(file)
	if err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.RequestIllegal)
		return nil, retcode.RequestIllegal
	}
	id, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	stateValue := ctx.Request.FormValue("state")
	state, _ := strconv.Atoi(stateValue)
	param := service.UpdateResourceRequest{
		Resource: entity.Resource{
			ID:    id,
			Name:  fileHeader.Filename,
			Blob:  data,
			Type:  fileHeader.Header.Get("Content-Type"),
			Size:  fileHeader.Size,
			State: uint8(state),
		},
	}
	if err := c.srv.Resources().Update(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.UpdateResourceFail)
		return nil, retcode.UpdateResourceFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
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
func (c *ResourceController) Delete(ctx *gin.Context) (res interface{}, err error) {
	id, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	if err := c.srv.Resources().Delete(ctx.Request.Context(), id); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.DeleteResourceFail)
		return nil, retcode.DeleteResourceFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}
