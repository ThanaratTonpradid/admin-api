package repository

import (
	"mini-api/model"
)

func (h Handler) CreateRole(entity *model.Role) error {
	return h.mysql.DB.Create(&entity).Error
}

func (h Handler) CreateRoleHasPermissions(entity *model.RolesHasPermission) error {
	return h.mysql.DB.Create(&entity).Error
}

func (h Handler) FindOneRoleByID(roleID uint32) (model.Role, error) {
	entity := model.Role{}
	err := h.mysql.DB.
		Where(&model.Role{
			ID: roleID,
		}, "ID").
		Take(&entity).Error
	return entity, err
}

func (h Handler) FindOneRoleByLabel(label string) (model.Role, error) {
	entity := model.Role{}
	err := h.mysql.DB.
		Where(&model.Role{
			Label: label,
		}, "Label").
		Take(&entity).Error
	return entity, err
}
