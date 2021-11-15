package service

import (
	"blog-service/internal/app"
	"blog-service/internal/model"
)

type CountTagRequest struct {
	Name  string
	State uint8
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

type ListTagRequest struct {
	Name  string
	State uint8
}

func (svc *Service) ListTag(param *ListTagRequest, page *app.Page) ([]*model.Tag, error) {
	return svc.dao.ListTag(param.Name, param.State, page.PageNum, page.PageSize)
}

type CreateTagRequest struct {
	Name      string
	CreatedBy string
	State     uint8
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

type UpdateTagRequest struct {
	ID         uint32
	Name       string
	State      uint8
	ModifiedBy string
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

type DeleteTagRequest struct {
	ID uint32
}

func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
