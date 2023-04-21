package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dollarsignteam/go-logger"

	"mini-api/helper"
	"mini-api/internal/api/constant"
	"mini-api/internal/api/dto"
	"mini-api/internal/repository"
	"mini-api/lib"
	"mini-api/model"
)

type AuthService struct {
	logger     *logger.Logger
	repository *repository.Handler
	jwtHandler *lib.JWTHandler
}

func NewAuthService(
	logger *logger.Logger,
	repository *repository.Handler,
	jwtHandler *lib.JWTHandler,
) AuthService {
	return AuthService{
		logger:     logger,
		repository: repository,
		jwtHandler: jwtHandler,
	}
}

func (svc AuthService) Login(req *dto.LoginRequest, ip string) (dto.LoginResponse, error) {
	staff, err := svc.repository.FindOneStaffByUsername(req.Username)
	if err != nil {
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeLoginFailed,
			ErrorInstance: err,
		}
	}
	isValid := helper.CompareHashPassword(staff.Password, req.Password)
	if !isValid {
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeLoginFailed,
			ErrorInstance: errors.New("invalid credentials"),
		}
	}

	return svc.LoginByStaff(staff, ip)
}

func (svc AuthService) LoginByStaff(staff model.Staff, ip string) (dto.LoginResponse, error) {
	if !staff.IsActive || staff.DeletedAt != nil {
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeLoginFailed,
			ErrorInstance: errors.New("staff is deleted"),
		}
	}
	subject := fmt.Sprintf("%d", staff.ID)
	jwtToken, err := svc.jwtHandler.CreateToken(subject)
	if err != nil {
		svc.logger.Error(GetStaffErrorMessage(staff.ID, err))
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeCreateTokenFailed,
			ErrorInstance: err,
		}
	}
	session := dto.Session{
		ID:        jwtToken.ID,
		Username:  staff.Username,
		StaffID:   staff.ID,
		CreatedAt: jwtToken.IssuedAt,
	}
	if err := svc.repository.SetStaffSession(session); err != nil {
		svc.logger.Error(GetStaffErrorMessage(staff.ID, err))
		return dto.LoginResponse{}, lib.CommonError{
			StatusCode:    http.StatusUnauthorized,
			ErrorCode:     constant.CodeCreateSessionFailed,
			ErrorInstance: err,
		}
	}
	if err := svc.repository.UpdateStaffLastLoginByID(staff.ID, ip); err != nil {
		svc.logger.Warn(GetStaffErrorMessage(staff.ID, err))
	}

	return dto.LoginResponse{
		Token: jwtToken.Token,
	}, nil
}

func (svc AuthService) GetSession(staffId uint32) (dto.Session, error) {
	session, err := svc.repository.GetStaffSession(staffId)
	if err != nil {
		return dto.Session{}, err
	}
	if err := helper.ValidateStruct(session); err != nil {
		return dto.Session{}, err
	}
	return session, nil
}

func (svc AuthService) Logout(session dto.Session) {
	svc.repository.DelStaffSession(session.StaffID)
}
