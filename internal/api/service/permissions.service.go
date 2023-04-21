package service

import (
	"github.com/dollarsignteam/go-logger"

	"mini-api/internal/api/constant"
	"mini-api/internal/repository"
	"mini-api/lib"
	"mini-api/model"
)

type PermissionsService struct {
	logger     *logger.Logger
	repository *repository.Handler
	jwtHandler *lib.JWTHandler
}

func NewPermissionsService(
	logger *logger.Logger,
	repository *repository.Handler,
	jwtHandler *lib.JWTHandler,
) PermissionsService {
	return PermissionsService{
		logger:     logger,
		repository: repository,
		jwtHandler: jwtHandler,
	}
}

func (svc PermissionsService) InitPermissions() {
	permissions := []string{
		constant.RolesRead,
		constant.RolesCreate,
		constant.RolesUpdate,
		constant.RolesDelete,
		constant.StaffsRead,
		constant.StaffsCreate,
		constant.StaffsUpdate,
		constant.StaffsDelete,
	}
	svc.logger.Info("Start init permission")
	for _, code := range permissions {
		svc.logger.Infof("Insert: %s", code)
		svc.CreatePermission(code)
	}
	svc.logger.Info("Init permission complete")
}

func (svc PermissionsService) FindOnePermissionByID(permissionID uint32) (model.Permission, error) {
	permission, err := svc.repository.FindOnePermissionByID(permissionID)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Permission{}, err)
	}
	return permission, nil
}

func (svc PermissionsService) CreatePermission(code string) (model.Permission, error) {
	unixTimeNow := GetUnixTimestamp()
	entity := model.Permission{
		CodeName:  code,
		CreatedAt: unixTimeNow,
		UpdatedAt: unixTimeNow,
	}
	if err := svc.repository.CreatePermission(&entity); err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Permission{}, err)
	}
	return svc.FindOnePermissionByID(entity.ID)
}
