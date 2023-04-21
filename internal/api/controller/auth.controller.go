package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"mini-api/internal/api/constant"
	"mini-api/internal/api/dto"
	"mini-api/internal/api/service"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(
	authService service.AuthService,
) AuthController {
	return AuthController{
		authService: authService,
	}
}

// @Tags    Auth
// @Summary Login staff by username and password
// @Produce json
// @Param   data body     dto.LoginRequest            true "Request payload"
// @Success 200  {object} dto.LoginResponse           "OK"
// @Failure 400  {object} dto.ErrorValidationResponse "Bad Request"
// @Failure 401  {object} dto.ErrorResponse           "Unauthorized"
// @Router  /api/auth/login [post]
func (ctrl AuthController) Login(c echo.Context) error {
	req := new(dto.LoginRequest)
	if err := c.Bind(req); err != nil {
		return NewCommonErrorBadRequest(err)
	}
	resp, err := ctrl.authService.Login(req, c.RealIP())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

// @Tags     Auth
// @Summary  Logout staff and invalidate token
// @Security Bearer
// @Produce  json
// @Success  200 {object} dto.SuccessResponse "OK"
// @Failure  401 {object} dto.ErrorResponse   "Unauthorized"
// @Router   /api/auth/logout [post]
func (ctrl AuthController) Logout(c echo.Context) error {
	session := GetSession(c)
	ctrl.authService.Logout(session)
	return c.JSON(http.StatusOK, dto.SuccessResponse{
		Success: true,
		Message: constant.CodeLogoutSuccess,
	})
}
