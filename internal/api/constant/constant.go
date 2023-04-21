package constant

import "time"

const TTLJWTExpires = 31 * 24 * time.Hour
const KeySession = "session"

const (
	ErrCodeSomethingWentWrong = "SOMETHING_WENT_WRONG"
	ErrCodeBadRequest         = "BAD_REQUEST"
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeInternalError      = "INTERNAL_ERROR"
)

const (
	CodeLoginFailed         = "Login failed"
	CodeCreateTokenFailed   = "CreateTokenFailed"
	CodeCreateSessionFailed = "CreateSessionFailed"
	CodeLogoutSuccess       = "LogoutSuccess"
)

const (
	RolesCreate = "ROLES_CREATE"
	RolesRead   = "ROLES_READ"
	RolesUpdate = "ROLES_UPDATE"
	RolesDelete = "ROLES_DELETE"

	StaffsCreate = "STAFFS_CREATE"
	StaffsRead   = "STAFFS_READ"
	StaffsUpdate = "STAFFS_UPDATE"
	StaffsDelete = "STAFFS_DELETE"
)

const (
	RoleAdmin  = "ADMIN"
	RoleStaff  = "STAFF"
	RoleMember = "MEMBER"
)
