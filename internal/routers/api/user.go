package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/config"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store/dao"
)

type LoginRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
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
// @Router /user/login [post]
func Login(writer http.ResponseWriter, request *http.Request) {
	var req LoginRequest

	data, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(data, &req); err != nil {
		app.ToResponseCode(writer, retcode.RequestUnMarshalError)
		return
	}
	param := service.AuthRequest{
		AppKey:    req.UserName,
		AppSecret: req.PassWord,
	}
	svc := service.NewAuthService(dao.NewDao(config.DBEngine))
	if err := svc.CheckAuth(request.Context(), &param); err != nil {
		app.ToResponseCode(writer, retcode.RequestAuthNotExists)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		app.ToResponseCode(writer, retcode.GenerateAuthTokenFail)
	}

	res := LoginResponse{Token: token}

	app.ToResponseData(writer, res)
}

func Info(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`{"ret_code":"000000","ret_desc":"Success","data":{"roles":["admin"],"introduction":"I am a super administrator","avatar":"https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/6a426929fa654a9ab8a58015ea9f573f~tplv-k3u1fbpfcp-zoom-crop-mark:3024:3024:3024:1702.awebp?","name":"Super Admin"}}`))
}

func Logout(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`{"ret_code":"000000","ret_desc":"Success"}`))
}
