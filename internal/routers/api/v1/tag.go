package v1

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store"
)

type TagController struct {
	srv service.Service
}

func NewTagController(store store.Factory) *TagController {
	return &TagController{
		srv: service.NewService(store, nil),
	}
}

// @Summary 获取多个标签
// @Description 获取多个标签
// @Tags Tag
// @Produce json
// @Security JWT
// @param name query string false "标签名称"
// @param state query integer false "状态"
// @param sort query string false "排序方式"
// @param page query integer false "页码"
// @param limit query integer false "页面大小"
// @Success 200 {object} app.ListResponse{data=[]model.Tag} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/tags [get]
func (c *TagController) List(ctx *gin.Context) (res interface{}, total int64, err error) {
	values := ctx.Request.URL.Query()
	name := values.Get("name")
	state, _ := strconv.Atoi(values.Get("state"))
	page, _ := strconv.Atoi(values.Get("page"))
	limit, _ := strconv.Atoi(values.Get("limit"))
	sort := values.Get("sort")
	param := service.TagListRequest{ListOption: model.ListOption{Name: name, State: uint8(state), Limit: limit, Page: page, Sort: sort}}
	tags, total, err := c.srv.Tags().List(ctx.Request.Context(), &param)
	if err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.GetTagsFail)
		return nil, 0, retcode.GetTagsFail
	}
	//app.ToResponseList(ctx.Writer, total, tags)
	return tags, total, nil
}

// @Summary 创建标签
// @Description 创建标签
// @Tags Tag
// @Produce json
// @Accept json
// @Security JWT
// @param CreateTagRequest body service.CreateTagRequest true "创建标签"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/tag [post]
func (c *TagController) Create(ctx *gin.Context) (res interface{}, err error) {
	var param service.CreateTagRequest
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return nil, retcode.RequestUnMarshalError
	}
	if err := c.srv.Tags().Create(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.CreateTagFail)
		return nil, retcode.CreateTagFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}

// @Summary 修改标签
// @Description 修改标签
// @Tags Tag
// @Produce json
// @Accept json
// @Security JWT
// @Param id path integer true "标签ID"
// @param UpdateTagRequest body service.UpdateTagRequest true "修改文章"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/tag [put]
func (c *TagController) Update(ctx *gin.Context) (res interface{}, err error) {
	var param service.UpdateTagRequest
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return nil, retcode.RequestUnMarshalError
	}
	if err := c.srv.Tags().Update(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.UpdateTagFail)
		return nil, retcode.UpdateTagFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}

// @Summary 删除标签
// @Description 删除标签
// @Tags Tag
// @Produce json
// @Security JWT
// @Param id path integer true "标签ID"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/tag [delete]
func (c *TagController) Delete(ctx *gin.Context) (res interface{}, err error) {
	id, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	param := service.DeleteTagRequest{OneOption: model.OneOption{Id: id}}
	if err := c.srv.Tags().Delete(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.DeleteTagFail)
		return nil, retcode.DeleteTagFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}
