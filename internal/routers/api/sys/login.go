package sys

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/pkg/captcha"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store"
)

type LoginController struct {
	srv service.Service
}

func NewLoginController(store store.Factory) *LoginController {
	return &LoginController{
		srv: service.NewService(store, nil),
	}
}

type LoginRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Captcha  string `json:"captcha"`
	Cid      string `json:"cid"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// @Summary 登录接口
// @Description 用户登录生成Token
// @Tags System
// @Produce json
// @Accept json
// @param LoginRequest body LoginRequest true "用户信息"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/sys/login [post]
func (c *LoginController) Login(ctx *gin.Context) {
	var req LoginRequest

	data, _ := ioutil.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &req); err != nil {
		app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return
	}
	param := service.AuthRequest{
		UserName: req.UserName,
		PassWord: req.PassWord,
		Captcha:  req.Captcha,
		Cid:      req.Cid,
	}
	if err := c.srv.Users().CheckAuth(ctx.Request.Context(), &param); err != nil {
		log.Errorf("check user auth err:%v", err)
		app.ToResponseCode(ctx.Writer, retcode.RequestAuthCheckFail)
		return
	}

	config, err := c.srv.Systems().Get(ctx.Request.Context(), &service.SystemGetRequest{OneOption: model.OneOption{Name: "EnableLoginCaptcha"}})
	if err != nil {
		log.Errorf("get system config err:%v", err)
		app.ToResponseCode(ctx.Writer, retcode.RequestAuthCheckFail)
		return
	}
	if config != nil && config.Value == "1" {
		if !captcha.Verify(param.Cid, param.Captcha) {
			log.Errorf("check user auth err:%v", err)
			app.ToResponseCode(ctx.Writer, retcode.RequestAuthCheckFail)
			return
		}
	}

	token, err := app.GenerateToken(param.UserId, param.UserName, param.PassWord)
	if err != nil {
		app.ToResponseCode(ctx.Writer, retcode.GenerateAuthTokenFail)
	}

	res := LoginResponse{Token: token}

	app.ToResponseData(ctx.Writer, res)
}

// @Summary 用户信息
// @Description 获取登录用户信息
// @Tags Auth
// @Produce json
// @Accept json
// @Security JWT
// @param LoginRequest body LoginRequest true "用户信息"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/user [get]
func (c *LoginController) Info(ctx *gin.Context) {
	name := ctx.GetString("username")
	if name == "" {
		app.ToResponseCode(ctx.Writer, retcode.RequestTokenAuthFail)
		return
	}
	user, err := c.srv.Users().GetUser(ctx.Request.Context(), name)
	if err != nil {
		app.ToResponseCode(ctx.Writer, retcode.RequestUserGetFail)
		return
	}
	user.Roles = []string{"admin"}
	app.ToResponseData(ctx.Writer, user)
}

// @Summary 退出接口
// @Description 用户退出清除cookie
// @Tags System
// @Produce json
// @Accept json
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/sys/logout [post]
func (c *LoginController) Logout(ctx *gin.Context) {
	app.ToResponseCode(ctx.Writer, retcode.Success)
}
