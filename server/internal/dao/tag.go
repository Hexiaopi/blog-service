package dao

import (
	"blog-service/internal/app"
	"blog-service/internal/model"
)

func (d *Dao) GetTag(id uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{ID: id, State: state}
	return tag.Get(d.engine)
}

func (d *Dao) ListTag(name string, state uint8, pageNum, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.List(d.engine, app.GetPageOffset(pageNum, pageSize), pageSize)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{Name: name, CreatedBy: createdBy, State: state}
	_, err := tag.Create(d.engine)
	return err
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{ID: id}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{ID: id}
	return tag.Delete(d.engine)
}
