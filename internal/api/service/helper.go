package service

import (
	"fmt"
	"net/http"
	"time"

	"mini-api/internal/api/constant"
	"mini-api/lib"
)

func GetUnixTimestamp() uint32 {
	unixTime := time.Now().Unix()
	return uint32(unixTime)
}

func GetAgentErrorMessage(agentUsername, memberUsername string, err error) string {
	return fmt.Sprintf("[agent: %s - member: %s] %s", agentUsername, memberUsername, err)
}

func GetStaffErrorMessage(staffId uint32, err error) string {
	return fmt.Sprintf("staff#%d: %s", staffId, err)
}

func NewCommonErrorSomethingWentWrong[T any](result T, err error) (T, error) {
	return result, lib.CommonError{
		StatusCode:    http.StatusInternalServerError,
		ErrorCode:     constant.ErrCodeInternalError,
		ErrorInstance: err,
	}
}

func NewGameApiError[T any](result T, errorCode string, err error) (T, error) {
	return result, lib.CommonError{
		StatusCode:    http.StatusBadRequest,
		ErrorCode:     errorCode,
		ErrorInstance: err,
	}
}
