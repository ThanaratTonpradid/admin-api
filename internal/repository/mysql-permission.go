package repository

import (
	"mini-api/model"
)

func (h Handler) CreatePermission(entity *model.Permission) error {
	return h.mysql.DB.Create(&entity).Error
}

func (h Handler) FindOnePermissionByID(permissionID uint32) (model.Permission, error) {
	entity := model.Permission{}
	err := h.mysql.DB.
		Where(&model.Permission{
			ID: permissionID,
		}, "ID").
		Take(&entity).Error
	return entity, err
}

func (h Handler) FindOnePermissionByCodeName(codeName string) (model.Permission, error) {
	entity := model.Permission{}
	err := h.mysql.DB.
		Where(&model.Permission{
			CodeName: codeName,
		}, "CodeName").
		Take(&entity).Error
	return entity, err
}
