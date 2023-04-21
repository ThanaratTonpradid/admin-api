package repository

import (
	"mini-api/model"
)

type StaffProfile struct {
	model.Staff
}

func (h Handler) CreateStaff(entity *model.Staff) error {
	return h.mysql.DB.Create(&entity).Error
}

func (h Handler) FindOneStaffByID(staffID uint32) (model.Staff, error) {
	entity := model.Staff{}
	if err := h.mysql.DB.
		Where(&model.Staff{
			ID: staffID,
		}, "ID").
		First(&entity).Error; err != nil {
		return model.Staff{}, err
	}
	return entity, nil
}

func (h Handler) FindOneStaffByUsername(username string) (model.Staff, error) {
	entity := model.Staff{}
	if err := h.mysql.DB.
		Where(&model.Staff{
			Username: username,
		}, "Username").
		First(&entity).Error; err != nil {
		return model.Staff{}, err
	}
	return entity, nil
}

func (h Handler) UpdateStaffLastLoginByID(staffID uint32, ip string) error {
	unixTimeNow := GetUnixTimestamp()
	return h.mysql.DB.
		Where(&model.Staff{
			ID: staffID,
		}, "ID").
		Updates(model.Staff{
			LastIP:    &ip,
			LastLogin: &unixTimeNow,
			UpdatedAt: unixTimeNow,
		}).Error
}
