package v1

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store"
	log "github.com/hexiaopi/blog-service/pkg/logger"
)

type RestController struct {
	srv    service.Service
	logger log.Logger
}

func NewRestController(store store.Factory, logger log.Logger) *RestController {
	return &RestController{
		srv:    service.NewService(store, nil, logger),
		logger: logger,
	}
}

// @Summary 获取多个接口
// @Description 获取多个接口
// @Tags Rest
// @Produce json
// @Security JWT
// @param name query string false "标签名称"
// @param sort query string false "排序方式"
// @param page query integer false "页码"
// @param limit query integer false "页面大小"
// @Success 200 {object} app.ListResponse{data=[]model.SysRest} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/rests [get]
func (c *RestController) List(ctx *gin.Context) (res interface{}, total int64, err error) {
	values := ctx.Request.URL.Query()
	name := values.Get("name")
	page, _ := strconv.Atoi(values.Get("page"))
	limit, _ := strconv.Atoi(values.Get("limit"))
	sort := values.Get("sort")
	param := service.SysRestListRequest{ListOption: entity.ListOption{Name: name, Limit: limit, Page: page, Sort: sort}}
	Rests, total, err := c.srv.Rests().List(ctx.Request.Context(), &param)
	if err != nil {
		return nil, 0, retcode.GetRestsFail
	}
	return Rests, total, nil
}

// @Summary 创建接口
// @Description 创建接口
// @Tags Rest
// @Produce json
// @Accept json
// @Security JWT
// @param CreateRestRequest body service.CreateRestRequest true "创建接口"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/rest [post]
func (c *RestController) Create(ctx *gin.Context) (res interface{}, err error) {
	var param service.CreateSysRestRequest
	data, _ := io.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return nil, retcode.RequestUnMarshalError
	}
	if err := c.srv.Rests().Create(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.CreateRestFail)
		return nil, retcode.CreateRestFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}

// @Summary 修改接口
// @Description 修改接口
// @Tags Rest
// @Produce json
// @Accept json
// @Security JWT
// @Param id path integer true "接口ID"
// @param UpdateRestRequest body service.UpdateRestRequest true "修改接口"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/rest [put]
func (c *RestController) Update(ctx *gin.Context) (res interface{}, err error) {
	var param service.UpdateSysRestRequest
	data, _ := io.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return nil, retcode.RequestUnMarshalError
	}
	if err := c.srv.Rests().Update(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.UpdateRestFail)
		return nil, retcode.UpdateRestFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}

// @Summary 删除接口
// @Description 删除接口
// @Tags Rest
// @Produce json
// @Security JWT
// @Param id path integer true "接口ID"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/rest [delete]
func (c *RestController) Delete(ctx *gin.Context) (res interface{}, err error) {
	id, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	if err := c.srv.Rests().Delete(ctx.Request.Context(), id); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.DeleteRestFail)
		return nil, retcode.DeleteRestFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}
