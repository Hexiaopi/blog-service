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
// @Router /api/v1/menu/tree [get]
func (c *MenuController) List(ctx *gin.Context) (res interface{}, total int64, err error) {
	values := ctx.Request.URL.Query()
	name := values.Get("name")
	page, _ := strconv.Atoi(values.Get("page"))
	limit, _ := strconv.Atoi(values.Get("limit"))
	sort := values.Get("sort")
	param := service.SysMenuListRequest{ListOption: entity.ListOption{Name: name, Limit: limit, Page: page, Sort: sort}}
	Rests, total, err := c.srv.Menus().List(ctx.Request.Context(), &param)
	if err != nil {
		return nil, 0, retcode.GetRestsFail
	}
	return Rests, total, nil
}
