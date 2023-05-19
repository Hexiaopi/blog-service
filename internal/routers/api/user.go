package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/internal/app"
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

type LoginRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Captcha  string `json:"captcha"`
	Cid      string `json:"cid"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// @Summary 鉴权接口
// @Description 请求鉴权获取Token
// @Tags Auth
// @Produce json
// @Accept json
// @param LoginRequest body LoginRequest true "用户信息"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /auth/login [post]
func (c *UserController) Login(writer http.ResponseWriter, request *http.Request) {
	var req LoginRequest

	data, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(data, &req); err != nil {
		app.ToResponseCode(writer, retcode.RequestUnMarshalError)
		return
	}
	param := service.AuthRequest{
		UserName: req.UserName,
		PassWord: req.PassWord,
		Captcha:  req.Captcha,
		Cid:      req.Cid,
	}
	if err := c.srv.Users().CheckAuth(request.Context(), &param); err != nil {
		log.Errorf("check user auth err:%v", err)
		app.ToResponseCode(writer, retcode.RequestAuthCheckFail)
		return
	}

	token, err := app.GenerateToken(param.UserName, param.PassWord)
	if err != nil {
		app.ToResponseCode(writer, retcode.GenerateAuthTokenFail)
	}

	res := LoginResponse{Token: token}

	app.ToResponseData(writer, res)
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
func (c *UserController) Info(writer http.ResponseWriter, request *http.Request) {
	name := request.Context().Value("name")
	if name == nil {
		app.ToResponseCode(writer, retcode.RequestTokenAuthFail)
		return
	}
	user, err := c.srv.Users().GetUser(request.Context(), name.(string))
	if err != nil {
		app.ToResponseCode(writer, retcode.RequestUserGetFail)
		return
	}
	user.Roles = []string{"admin"}
	app.ToResponseData(writer, user)
}

func Logout(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`{"ret_code":"000000","ret_desc":"Success"}`))
}
