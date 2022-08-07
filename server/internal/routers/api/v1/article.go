package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
)

// @Summary 获取单个文章
// @Description 获取单个文章详细信息
// @Tags Article
// @Accept  json
// @Produce  json
// @Param id path integer true "文章ID"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/articles/{id} [get]
func GetArticle(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.ParseUint(mux.Vars(request)["id"], 10, 32)
	param := service.ArticleRequest{ID: uint32(id), State: 1}
	svc := service.New(request.Context())
	article, err := svc.GetArticle(&param)
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
// @param name query string false "文章名称"
// @param tag_id query integer false "标签ID"
// @param state query integer false "状态"
// @param page_num query integer false "页码"
// @param page_size query integer false "每页数量"
// @Success 200 {object} app.ListResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/articles [get]
func ListArticle(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	tagId, _ := strconv.Atoi(values.Get("tag_id"))
	state, _ := strconv.Atoi(values.Get("state"))
	pageNum, _ := strconv.Atoi(values.Get("page_num"))
	pageSize, _ := strconv.Atoi(values.Get("page_size"))
	param := service.ArticleListRequest{TagID: uint32(tagId), State: uint8(state)}
	page := app.CorrectPage(pageSize, pageNum)
	svc := service.New(request.Context())
	article, count, err := svc.ListArticle(&param, &page)
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
// @param CreateArticleRequest body service.CreateArticleRequest true "创建文章"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/articles [post]
func CreateArticle(writer http.ResponseWriter, request *http.Request) {
	var param service.CreateArticleRequest
	data, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(writer, retcode.RequestUnMarshalError)
		return
	}

	svc := service.New(request.Context())
	if err := svc.CreateArticle(&param); err != nil {
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
// @Param id path integer true "文章ID"
// @param UpdateArticleRequest body service.UpdateArticleRequest true "修改文章"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/articles/{id} [put]
func UpdateArticle(writer http.ResponseWriter, request *http.Request) {
	var param service.UpdateArticleRequest
	data, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(writer, retcode.RequestUnMarshalError)
		return
	}

	id, _ := strconv.ParseUint(mux.Vars(request)["id"], 10, 32)
	param.ID = uint32(id)

	svc := service.New(request.Context())
	if err := svc.UpdateArticle(&param); err != nil {
		app.ToResponseCode(writer, retcode.UpdateArticleFail)
		return
	}
	app.ToResponseData(writer, nil)
}

// @Summary 删除文章
// @Description 删除文章
// @Tags Article
// @Produce json
// @Param id path integer true "文章ID"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/articles/{id} [delete]
func DeleteArticle(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.ParseUint(mux.Vars(request)["id"], 10, 32)
	svc := service.New(request.Context())
	if err := svc.DeleteArticle(uint32(id)); err != nil {
		app.ToResponseCode(writer, retcode.DeleteArticleFail)
		return
	}
	app.ToResponseData(writer, nil)
}
