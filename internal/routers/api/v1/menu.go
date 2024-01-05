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

type MenuController struct {
	srv    service.Service
	logger log.Logger
}

func NewMenuController(store store.Factory, logger log.Logger) *MenuController {
	return &MenuController{
		srv:    service.NewService(store, nil, logger),
		logger: logger,
	}
}

// @Summary 获取菜单树
// @Description 根据角色获取菜单树结构
// @Tags Menu
// @Produce json
// @Security JWT
// @param name query string false "菜单名称"
// @param sort query string false "排序方式"
// @Success 200 {object} app.ListResponse{data=[]entity.SysMenu} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/menu/tree [get]
func (c *MenuController) Tree(ctx *gin.Context) (res interface{}, total int64, err error) {
	values := ctx.Request.URL.Query()
	name := values.Get("name")
	sort := values.Get("sort")
	param := service.SysMenuListRequest{ListOption: entity.ListOption{Name: name, Sort: sort}}
	Rests, total, err := c.srv.Menus().Tree(ctx.Request.Context(), &param)
	if err != nil {
		return nil, 0, retcode.GetRestsFail
	}
	return Rests, total, nil
}

// @Summary 获取菜单列表
// @Description 获取菜单列表及其子菜单
// @Tags Menu
// @Produce json
// @Security JWT
// @param name query string false "菜单名称"
// @param sort query string false "排序方式"
// @Success 200 {object} app.ListResponse{data=[]entity.SysMenu} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/menus [get]
func (c *MenuController) List(ctx *gin.Context) (res interface{}, total int64, err error) {
	values := ctx.Request.URL.Query()
	name := values.Get("name")
	sort := values.Get("sort")
	param := service.SysMenuListRequest{ListOption: entity.ListOption{Name: name, Sort: sort}}
	Rests, total, err := c.srv.Menus().List(ctx.Request.Context(), &param)
	if err != nil {
		return nil, 0, retcode.GetRestsFail
	}
	return Rests, total, nil
}

// @Summary 创建菜单
// @Description 创建菜单
// @Tags Menu
// @Produce json
// @Accept json
// @Security JWT
// @param CreateRestRequest body service.CreateSysMenuRequest true "创建接口"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/menu [post]
func (c *MenuController) Create(ctx *gin.Context) (res interface{}, err error) {
	var param service.CreateSysMenuRequest
	data, _ := io.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		c.logger.Errorf("unmarshal request err:%v", err)
		return nil, retcode.RequestUnMarshalError
	}
	if err := c.srv.Menus().Create(ctx.Request.Context(), &param); err != nil {
		return nil, retcode.CreateRestFail
	}
	return nil, nil
}

// @Summary 修改菜单
// @Description 修改菜单
// @Tags Menu
// @Produce json
// @Accept json
// @Security JWT
// @Param id path integer true "才带你ID"
// @param UpdateRestRequest body service.UpdateSysMenuRequest true "修改接口"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/menu [put]
func (c *MenuController) Update(ctx *gin.Context) (res interface{}, err error) {
	var param service.UpdateSysMenuRequest
	data, _ := io.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		c.logger.Errorf("unmarshal request err:%v", err)
		return nil, retcode.RequestUnMarshalError
	}
	if err := c.srv.Menus().Update(ctx.Request.Context(), &param); err != nil {
		return nil, retcode.UpdateRestFail
	}
	return nil, nil
}

// @Summary 删除菜单
// @Description 删除菜单
// @Tags Menu
// @Produce json
// @Security JWT
// @Param id path integer true "菜单ID"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/menu [delete]
func (c *MenuController) Delete(ctx *gin.Context) (res interface{}, err error) {
	id, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	if err := c.srv.Menus().Delete(ctx.Request.Context(), id); err != nil {
		return nil, retcode.DeleteRestFail
	}
	return nil, nil
}
