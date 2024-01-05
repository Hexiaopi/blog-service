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

type RoleController struct {
	srv    service.Service
	logger log.Logger
}

func NewRoleController(store store.Factory, logger log.Logger) *RoleController {
	return &RoleController{
		srv:    service.NewService(store, nil, logger),
		logger: logger,
	}
}

// @Summary 获取多个角色
// @Description 获取多个角色
// @Tags Role
// @Produce json
// @Security JWT
// @param name query string false "角色名称"
// @param state query integer false "状态"
// @param sort query string false "排序方式"
// @param page query integer false "页码"
// @param limit query integer false "页面大小"
// @Success 200 {object} app.ListResponse{data=[]model.Role} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/roles [get]
func (c *RoleController) List(ctx *gin.Context) (res interface{}, total int64, err error) {
	values := ctx.Request.URL.Query()
	name := values.Get("name")
	state, _ := strconv.Atoi(values.Get("state"))
	page, _ := strconv.Atoi(values.Get("page"))
	limit, _ := strconv.Atoi(values.Get("limit"))
	sort := values.Get("sort")
	param := service.ListRoleRequest{ListOption: entity.ListOption{Name: name, State: uint8(state), Limit: limit, Page: page, Sort: sort}}
	roles, total, err := c.srv.Roles().List(ctx.Request.Context(), &param)
	if err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.GetTagsFail)
		return nil, 0, retcode.GetRolesFail
	}
	//app.ToResponseList(ctx.Writer, total, roles)
	return roles, total, nil
}

// @Summary 创建角色
// @Description 创建角色
// @Tags Role
// @Produce json
// @Accept json
// @Security JWT
// @param CreateRoleRequest body service.CreateRoleRequest true "创建标签"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/role [post]
func (c *RoleController) Create(ctx *gin.Context) (res interface{}, err error) {
	var param service.CreateRoleRequest
	data, _ := io.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return nil, retcode.RequestUnMarshalError
	}
	if err := c.srv.Roles().Create(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.CreateTagFail)
		return nil, retcode.CreateRoleFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}

// @Summary 修改角色
// @Description 修改角色
// @Tags Role
// @Produce json
// @Accept json
// @Security JWT
// @Param id path integer true "角色ID"
// @param UpdateRoleRequest body service.UpdateRoleRequest true "修改文章"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/role [put]
func (c *RoleController) Update(ctx *gin.Context) (res interface{}, err error) {
	var param service.UpdateRoleRequest
	data, _ := io.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return nil, retcode.RequestUnMarshalError
	}
	if err := c.srv.Roles().Update(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.UpdateTagFail)
		return nil, retcode.UpdateRoleFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}

// @Summary 修改角色菜单
// @Description 修改角色菜单
// @Tags Role
// @Produce json
// @Accept json
// @Security JWT
// @Param id path integer true "角色ID"
// @param UpdateRoleRequest body service.UpdateRoleMenuRequest true "修改角色菜单"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/role/:id/menu [put]
func (c *RoleController) UpdateMenu(ctx *gin.Context) (res interface{}, err error) {
	var param service.UpdateRoleMenuRequest
	data, _ := io.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		return nil, retcode.RequestUnMarshalError
	}
	roleId := ctx.Param("id")
	param.RoleId, err = strconv.Atoi(roleId)
	if err != nil {
		return nil, retcode.RequestIllegal
	}
	if err := c.srv.Roles().UpdateMenu(ctx.Request.Context(), &param); err != nil {
		return nil, retcode.UpdateRoleFail
	}
	return nil, nil
}

// @Summary 修改角色接口
// @Description 修改角色接口
// @Tags Role
// @Produce json
// @Accept json
// @Security JWT
// @Param id path integer true "角色ID"
// @param UpdateRoleRequest body service.UpdateRoleRestRequest true "修改角色接口"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/role/:id/rest [put]
func (c *RoleController) UpdateRest(ctx *gin.Context) (res interface{}, err error) {
	var param service.UpdateRoleRestRequest
	data, _ := io.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		return nil, retcode.RequestUnMarshalError
	}
	roleId := ctx.Param("id")
	param.RoleId, err = strconv.Atoi(roleId)
	if err != nil {
		return nil, retcode.RequestIllegal
	}
	if err := c.srv.Roles().UpdateRest(ctx.Request.Context(), &param); err != nil {
		return nil, retcode.UpdateRoleFail
	}
	return nil, nil
}

// @Summary 删除角色
// @Description 删除角色
// @Tags Role
// @Produce json
// @Security JWT
// @Param id path integer true "角色ID"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/role [delete]
func (c *RoleController) Delete(ctx *gin.Context) (res interface{}, err error) {
	id, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	param := service.DeleteRoleRequest{OneOption: entity.OneOption{Id: id}}
	if err := c.srv.Roles().Delete(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.DeleteTagFail)
		return nil, retcode.DeleteRoleFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}
