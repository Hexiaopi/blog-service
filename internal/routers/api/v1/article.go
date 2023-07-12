package v1

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hexiaopi/blog-service/internal/cache"
	"github.com/hexiaopi/blog-service/internal/model"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store"
)

type ArticleController struct {
	srv service.Service
}

func NewArticleController(store store.Factory, cache cache.Factory) *ArticleController {
	return &ArticleController{
		srv: service.NewService(store, cache),
	}
}

// @Summary 获取多个文章
// @Description 根据条件获取多个文章详细信息
// @Tags Article
// @Produce json
// @Security JWT
// @param name query string false "文章名称"
// @param state query integer false "状态"
// @param page query integer false "页码"
// @param limit query integer false "每页数量"
// @Success 200 {object} app.ListResponse{data=[]model.Article} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/articles [get]
func (c *ArticleController) List(ctx *gin.Context) (res interface{}, total int64, err error) {
	values := ctx.Request.URL.Query()
	name := values.Get("name")
	state, _ := strconv.Atoi(values.Get("state"))
	page, _ := strconv.Atoi(values.Get("page"))
	limit, _ := strconv.Atoi(values.Get("limit"))
	sort := values.Get("sort")
	param := service.ArticleListRequest{ListOption: model.ListOption{Name: name, State: uint8(state), Limit: limit, Page: page, Sort: sort}}
	article, count, err := c.srv.Articles().List(ctx.Request.Context(), &param)
	if err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.GetArticlesFail)
		return nil, 0, retcode.GetArticlesFail
	}
	//app.ToResponseList(ctx.Writer, count, article)
	return article, count, nil
}

// @Summary 创建文章
// @Description 创建带标签的文章
// @Tags Article
// @Produce json
// @Accept json
// @Security JWT
// @param CreateArticleRequest body service.CreateArticleRequest true "创建文章"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/articles [post]
func (c *ArticleController) Create(ctx *gin.Context) (res interface{}, err error) {
	var param service.CreateArticleRequest
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return nil, retcode.RequestUnMarshalError
	}
	if err := c.srv.Articles().Create(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.CreateArticleFail)
		return nil, retcode.CreateArticleFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}

// @Summary 获取单个文章
// @Description 获取单个文章详细信息
// @Tags Article
// @Accept  json
// @Produce  json
// @Security JWT
// @Param id path integer true "文章ID"
// @Success 200 {object} app.CommResponse{data=model.Article} "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/article/{id} [get]
func (c *ArticleController) Get(ctx *gin.Context) (res interface{}, err error) {
	id, _ := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	param := service.ArticleRequest{OneOption: model.OneOption{Id: id}}
	article, err := c.srv.Articles().Get(ctx.Request.Context(), &param)
	if err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.GetArticleFail)
		return nil, retcode.GetArticleFail
	}
	//app.ToResponseData(ctx.Writer, article)
	return article, nil
}

// @Summary 修改文章
// @Description 修改带标签的文章
// @Tags Article
// @Produce json
// @Accept json
// @Security JWT
// @Param id path integer true "文章ID"
// @param UpdateArticleRequest body service.UpdateArticleRequest true "修改文章"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/articles/{id} [put]
func (c *ArticleController) Update(ctx *gin.Context) (res interface{}, err error) {
	var param service.UpdateArticleRequest
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.RequestUnMarshalError)
		return nil, retcode.RequestUnMarshalError
	}
	if err := c.srv.Articles().Update(ctx.Request.Context(), &param); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.UpdateArticleFail)
		return nil, retcode.UpdateArticleFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}

// @Summary 删除文章
// @Description 删除文章
// @Tags Article
// @Produce json
// @Security JWT
// @Param id path integer true "文章ID"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (c *ArticleController) Delete(ctx *gin.Context) (res interface{}, err error) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.srv.Articles().Delete(ctx.Request.Context(), id); err != nil {
		//app.ToResponseCode(ctx.Writer, retcode.DeleteArticleFail)
		return nil, retcode.DeleteArticleFail
	}
	//app.ToResponseData(ctx.Writer, nil)
	return nil, nil
}
