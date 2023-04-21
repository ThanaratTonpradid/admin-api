package route

import (
	"net/http"

	"github.com/dollarsignteam/go-logger"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag"

	"mini-api/config"
	"mini-api/internal/api/doc"
	"mini-api/lib"
)

// @securityDefinitions.apikey Bearer
// @in                         header
// @name                       Authorization
// @schemes                    http https
// @basePath                   /
type SwaggerRoute struct {
	logger  *logger.Logger
	handler *lib.HttpHandler
}

func NewSwaggerRoute(
	logger *logger.Logger,
	handler *lib.HttpHandler,
) SwaggerRoute {
	return SwaggerRoute{
		logger:  logger,
		handler: handler,
	}
}

func (r SwaggerRoute) Setup() {
	doc.SwaggerInfoAPI.Title = "MINI API"
	doc.SwaggerInfoAPI.Version = config.Version
	swag.Register(swag.Name, doc.SwaggerInfoAPI)

	r.logger.Info("Setting up swagger routes")
	r.handler.Engine.GET("/swagger/*", echoSwagger.WrapHandler)
	r.handler.Engine.GET("/swagger", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
}
