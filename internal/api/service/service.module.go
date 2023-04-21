package service

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthService),
	fx.Provide(NewPermissionsService),
	fx.Provide(NewRolesService),
	fx.Provide(NewStaffsService),
)
