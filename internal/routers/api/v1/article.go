package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/hexiaopi/blog-service/internal/app"
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
// @Router /api/v1/articles/{id} [get]
func (c *ArticleController) Get(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(request.URL.Query().Get("id"))
	param := service.ArticleRequest{OneOption: model.OneOption{Id: id}}
	article, err := c.srv.Articles().Get(request.Context(), &param)
	if err != nil {
		app.ToResponseCode(writer, retcode.GetArticleFail)
		return
	}
	app.ToResponseData(writer, article)
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
func (c *ArticleController) List(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	name := values.Get("name")
	state, _ := strconv.Atoi(values.Get("state"))
	page, _ := strconv.Atoi(values.Get("page"))
	limit, _ := strconv.Atoi(values.Get("limit"))
	sort := values.Get("sort")
	param := service.ArticleListRequest{ListOption: model.ListOption{Name: name, State: uint8(state), Limit: limit, Page: page, Sort: sort}}
	article, count, err := c.srv.Articles().List(request.Context(), &param)
	if err != nil {
		app.ToResponseCode(writer, retcode.GetArticlesFail)
		return
	}
	app.ToResponseList(writer, count, article)
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
func (c *ArticleController) Create(writer http.ResponseWriter, request *http.Request) {
	var param service.CreateArticleRequest
	data, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(writer, retcode.RequestUnMarshalError)
		return
	}
	if err := c.srv.Articles().Create(request.Context(), &param); err != nil {
		app.ToResponseCode(writer, retcode.CreateArticleFail)
		return
	}
	app.ToResponseData(writer, nil)
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
func (c *ArticleController) Update(writer http.ResponseWriter, request *http.Request) {
	var param service.UpdateArticleRequest
	data, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(writer, retcode.RequestUnMarshalError)
		return
	}
	if err := c.srv.Articles().Update(request.Context(), &param); err != nil {
		app.ToResponseCode(writer, retcode.UpdateArticleFail)
		return
	}
	app.ToResponseData(writer, nil)
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
func (c *ArticleController) Delete(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(request)["id"])
	if err := c.srv.Articles().Delete(request.Context(), id); err != nil {
		app.ToResponseCode(writer, retcode.DeleteArticleFail)
		return
	}
	app.ToResponseData(writer, nil)
}
