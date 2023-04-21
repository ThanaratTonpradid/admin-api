package repository

import (
	"github.com/dollarsignteam/go-logger"

	"mini-api/lib"
)

type Handler struct {
	redis  *lib.Redis
	mysql  *lib.MySQL
	logger *logger.Logger
}

func NewHandler(
	redis *lib.Redis,
	mysql *lib.MySQL,
	logger *logger.Logger,
) *Handler {
	return &Handler{
		redis:  redis,
		mysql:  mysql,
		logger: logger,
	}
}
