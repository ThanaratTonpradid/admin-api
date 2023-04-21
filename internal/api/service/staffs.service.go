package service

import (
	"github.com/dollarsignteam/go-logger"

	"mini-api/config"
	"mini-api/helper"
	"mini-api/internal/api/constant"
	"mini-api/internal/repository"
	"mini-api/lib"
	"mini-api/model"
)

type StaffInit struct {
	RolesName string
	Username  string
	Password  string
	Fullname  string
	IsAdmin   bool
}

type StaffsService struct {
	logger     *logger.Logger
	repository *repository.Handler
	jwtHandler *lib.JWTHandler
	cfg        *config.APIConfig
}

func NewStaffsService(
	logger *logger.Logger,
	repository *repository.Handler,
	jwtHandler *lib.JWTHandler,
	cfg *config.APIConfig,
) StaffsService {
	return StaffsService{
		logger:     logger,
		repository: repository,
		jwtHandler: jwtHandler,
		cfg:        cfg,
	}
}

func (svc StaffsService) InitStaffs() {
	staff := StaffInit{
		RolesName: constant.RoleAdmin,
		Username:  "admin@dev",
		Password:  svc.cfg.DefaultDevPassword,
		Fullname:  "dev",
		IsAdmin:   true,
	}
	svc.CreateStaff(staff)
}

func (svc StaffsService) CreateStaff(req StaffInit) (model.Staff, error) {
	role, err := svc.repository.FindOneRoleByLabel(req.RolesName)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Staff{}, err)
	}
	unixTimeNow := GetUnixTimestamp()
	password, err := helper.GenerateHashPassword(req.Password)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Staff{}, err)
	}
	entity := model.Staff{
		RolesID:   role.ID,
		Username:  req.Username,
		Password:  password,
		IsAdmin:   req.IsAdmin,
		CreatedAt: unixTimeNow,
		UpdatedAt: unixTimeNow,
	}
	if err := svc.repository.CreateStaff(&entity); err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(model.Staff{}, err)
	}
	return svc.repository.FindOneStaffByID(entity.ID)
}
