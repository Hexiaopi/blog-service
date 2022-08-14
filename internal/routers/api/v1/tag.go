package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/hexiaopi/blog-service/global"
	"github.com/hexiaopi/blog-service/internal/app"
	"github.com/hexiaopi/blog-service/internal/retcode"
	"github.com/hexiaopi/blog-service/internal/service"
	"github.com/hexiaopi/blog-service/internal/store/dao"
)

// @Summary 获取多个标签
// @Description 获取多个标签
// @Tags Tag
// @Produce json
// @param name query string false "标签名称"
// @param state query integer false "状态"
// @param page_num query integer false "页码"
// @param page_size query integer false "每页数量"
// @Success 200 {object} app.ListResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/tags [get]
func ListTag(writer http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()
	name := values.Get("name")
	state, _ := strconv.Atoi(values.Get("state"))
	state = 1
	pageNum, _ := strconv.Atoi(values.Get("page_num"))
	pageSize, _ := strconv.Atoi(values.Get("page_size"))
	param := service.TagListRequest{Name: name, State: uint8(state)}
	page := app.CorrectPage(pageSize, pageNum)
	svc := service.NewTagService(dao.NewDao(global.DBEngine))

	tags, total, err := svc.List(request.Context(), &param, &page)
	if err != nil {
		app.ToResponseCode(writer, retcode.GetTagsFail)
		return
	}
	app.ToResponseList(writer, total, tags)
}

// @Summary 创建标签
// @Description 创建标签
// @Tags Tag
// @Produce json
// @Accept json
// @param CreateTagRequest body service.CreateTagRequest true "创建标签"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/tags [post]
func CreateTag(writer http.ResponseWriter, request *http.Request) {
	var param service.CreateTagRequest
	data, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(writer, retcode.RequestUnMarshalError)
		return
	}

	svc := service.NewTagService(dao.NewDao(global.DBEngine))
	if err := svc.Create(request.Context(), &param); err != nil {
		app.ToResponseCode(writer, retcode.CreateTagFail)
		return
	}
	app.ToResponseData(writer, nil)
}

// @Summary 修改标签
// @Description 修改标签
// @Tags Tag
// @Produce json
// @Accept json
// @Param id path integer true "标签ID"
// @param UpdateTagRequest body service.UpdateTagRequest true "修改文章"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/tags/{id} [put]
func UpdateTag(writer http.ResponseWriter, request *http.Request) {
	var param service.UpdateTagRequest
	data, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(data, &param); err != nil {
		app.ToResponseCode(writer, retcode.RequestUnMarshalError)
		return
	}

	id, _ := strconv.ParseUint(mux.Vars(request)["id"], 10, 32)
	param.ID = uint32(id)

	svc := service.NewTagService(dao.NewDao(global.DBEngine))
	if err := svc.Update(request.Context(), &param); err != nil {
		app.ToResponseCode(writer, retcode.UpdateTagFail)
		return
	}
	app.ToResponseData(writer, nil)
}

// @Summary 删除标签
// @Description 删除标签
// @Tags Tag
// @Produce json
// @Param id path integer true "标签ID"
// @Success 200 {object} app.CommResponse "成功"
// @Failure 400 {object} app.ErrResponse "请求错误"
// @Failure 500 {object} app.ErrResponse "内部错误"
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(request)["id"])
	param := service.DeleteTagRequest{ID: id}
	svc := service.NewTagService(dao.NewDao(global.DBEngine))
	if err := svc.Delete(request.Context(), &param); err != nil {
		app.ToResponseCode(writer, retcode.DeleteTagFail)
		return
	}
	app.ToResponseData(writer, nil)
}
