package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"blog-service/internal/app"
	"blog-service/internal/retcode"
	"blog-service/internal/service"
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
	svc := service.New(request.Context())
	if err := svc.CheckAuth(&param); err != nil {
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
	writer.Write([]byte(`{"ret_code":"000000","ret_desc":"Success","data":{"roles":["admin"],"introduction":"I am a super administrator","avatar":"https://img2.baidu.com/it/u=1314332406,1737009348&fm=26&fmt=auto","name":"Super Admin"}}`))
}

func Logout(writer http.ResponseWriter,request *http.Request){
	writer.Write([]byte(`{"ret_code":"000000","ret_desc":"Success"}`))
}
