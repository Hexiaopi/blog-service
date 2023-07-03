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
)

type UserController struct {
	srv service.Service
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{
		srv: service.NewService(store, nil),
	}
}

// @Summary 获取多个用户
// @Description 获取多个用户
// @Tags User
// @Produce json
// @Security JWT
// @param name query string false "用户名称"
// @param state query integer false "状态"
// @param page_num query integer false "页码"
// @param page_size query integer false "每页数量"
// @Success 200 {object} app.ListResponse{data=[]model.User} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/users [get]
func (c *UserController) List(ctx *gin.Context) {
	values := ctx.Request.URL.Query()
	name := values.Get("name")
	//state, _ := strconv.Atoi(values.Get("state"))
	page, _ := strconv.Atoi(values.Get("page"))
	limit, _ := strconv.Atoi(values.Get("limit"))
	sort := values.Get("sort")
	param := service.ListUserRequest{ListOption: model.ListOption{Name: name, Limit: limit, Page: page, Sort: sort}}
	users, total, err := c.srv.Users().List(ctx.Request.Context(), &param)
	if err != nil {
		app.ToResponseCode(ctx.Writer, retcode.GetUsersFail)
		return
	}
	app.ToResponseList(ctx.Writer, total, users)
}

// @Summary 创建用户
// @Description 创建用户
// @Tags User
// @Produce json
// @Accept json
// @Security JWT
// @param CreateUserRequest body service.CreateUserRequest true "创建用户"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/user [post]
func (c *UserController) Create(ctx *gin.Context) {
	var param service.CreateUserRequest
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return
	}
	if err := c.srv.Users().Create(ctx.Request.Context(), &param); err != nil {
		app.ToResponseCode(ctx.Writer, retcode.CreateUserFail)
		return
	}
	app.ToResponseData(ctx.Writer, nil)
}

// @Summary 修改用户
// @Description 修改用户信息
// @Tags User
// @Produce json
// @Accept json
// @Security JWT
// @Param id path integer true "用户ID"
// @param UpdateUserRequest body service.UpdateUserRequest true "修改用户"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/user [put]
func (c *UserController) Update(ctx *gin.Context) {
	var param service.UpdateUserRequest
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return
	}
	if err := c.srv.Users().Update(ctx.Request.Context(), &param); err != nil {
		app.ToResponseCode(ctx.Writer, retcode.UpdateUserFail)
		return
	}
	app.ToResponseData(ctx.Writer, nil)
}

// @Summary 删除用户
// @Description 删除用户
// @Tags User
// @Produce json
// @Security JWT
// @Param id path integer true "用户ID"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/user [delete]
func (c *UserController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	param := service.DeleteUserRequest{OneOption: model.OneOption{Id: id}}
	if err := c.srv.Users().Delete(ctx.Request.Context(), &param); err != nil {
		app.ToResponseCode(ctx.Writer, retcode.DeleteUserFail)
		return
	}
	app.ToResponseData(ctx.Writer, nil)
}
