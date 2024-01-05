package sys

import (
	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/blog-service/internal/entity"
	"github.com/hexiaopi/blog-service/internal/pkg/captcha"
)

type CaptchaController struct {
}

func NewCaptchaController() *CaptchaController {
	return &CaptchaController{}
}

// @Summary 获取验证码
// @Description 获取验证码
// @Tags System
// @Produce json
// @param name query string false "名称"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/sys/captcha [get]
func (c *CaptchaController) Get(ctx *gin.Context) (res interface{}, err error) {
	id, base := captcha.Generate()
	captcha := entity.Captcha{
		Captcha: base,
		Cid:     id,
	}
	//app.ToResponseData(ctx.Writer, captcha)
	return captcha, nil
}
