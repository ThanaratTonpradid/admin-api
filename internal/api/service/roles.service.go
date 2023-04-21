package service

import (
	"github.com/dollarsignteam/go-logger"

	"mini-api/internal/api/constant"
	"mini-api/internal/repository"
	"mini-api/lib"
	"mini-api/model"
)

type RoleDefault struct {
	Label       string
	Permissions []string
}

type RolesService struct {
	logger     *logger.Logger
	repository *repository.Handler
	jwtHandler *lib.JWTHandler
}

func NewRolesService(
	logger *logger.Logger,
	repository *repository.Handler,
	jwtHandler *lib.JWTHandler,
) RolesService {
	return RolesService{
		logger:     logger,
		repository: repository,
		jwtHandler: jwtHandler,
	}
}

func (svc RolesService) InitRoles() {
	roles := []RoleDefault{}
	roles = append(roles, RoleDefault{
		Label: constant.RoleAdmin,
		Permissions: []string{
			constant.RolesRead,
			constant.RolesCreate,
			constant.RolesUpdate,
			constant.RolesDelete,
			constant.StaffsRead,
			constant.StaffsCreate,
			constant.StaffsUpdate,
			constant.StaffsDelete,
		},
	})
	roles = append(roles, RoleDefault{
		Label: constant.RoleStaff,
		Permissions: []string{
			constant.RolesRead,
			constant.StaffsRead,
		},
	})
	roles = append(roles, RoleDefault{
		Label: constant.RoleMember,
		Permissions: []string{
			constant.RolesRead,
			constant.StaffsRead,
		},
	})
	svc.logger.Info("Start init role")
	for _, r := range roles {
		svc.logger.Infof("Insert: %s", r.Label)
		createdRole, _ := svc.CreateRole(r.Label)
		for _, p := range r.Permissions {
			findPermission, _ := svc.repository.FindOnePermissionByCodeName(p)
			svc.CreateRoleHasPermissions(createdRole.ID, findPermission.ID)
		}
	}
	svc.logger.Info("Init role complete")
}

func (svc RolesService) FindOneRoleByID(roleID uint32) (model.Role, error) {
	role, err := svc.repository.FindOneRoleByID(roleID)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Role{}, err)
	}
	return role, nil
}

func (svc RolesService) FindOneRoleByLabel(label string) (model.Role, error) {
	role, err := svc.repository.FindOneRoleByLabel(label)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Role{}, err)
	}
	return role, nil
}

func (svc RolesService) CreateRole(roleName string) (model.Role, error) {
	unixTimeNow := GetUnixTimestamp()
	entity := model.Role{
		Label:     roleName,
		CreatedAt: unixTimeNow,
		UpdatedAt: unixTimeNow,
	}
	if err := svc.repository.CreateRole(&entity); err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Role{}, err)
	}
	return svc.FindOneRoleByID(entity.ID)
}

func (svc RolesService) CreateRoleHasPermissions(createdRoleID uint32, permissionID uint32) (model.Role, error) {
	entity := model.RolesHasPermission{
		RolesID:       createdRoleID,
		PermissionsID: permissionID,
	}
	if err := svc.repository.CreateRoleHasPermissions(&entity); err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Role{}, err)
	}
	return svc.FindOneRoleByID(entity.ID)
}
