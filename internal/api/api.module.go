package api

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"sync"

	"github.com/dollarsignteam/go-fileprobe"
	"github.com/dollarsignteam/go-logger"
	"github.com/imroc/req/v3"
	"go.uber.org/fx"

	"mini-api/config"
	"mini-api/internal/api/constant"
	"mini-api/internal/api/controller"
	"mini-api/internal/api/middleware"
	"mini-api/internal/api/route"
	"mini-api/internal/api/service"
	"mini-api/internal/repository"
	"mini-api/lib"
)

var Module = fx.Options(
	fx.Provide(NewLoggerOptions),
	fx.Provide(NewRedisOptions),
	fx.Provide(NewMySQLOptions),
	fx.Provide(NewJWTOptions),
	fx.Provide(config.NewAPIConfig),
	fx.Provide(lib.NewHttpHandler),
	fx.Provide(lib.NewRedis),
	fx.Provide(lib.NewMySQL),
	fx.Provide(lib.NewJWTHandler),
	fx.Provide(logger.NewLogger),
	fx.Provide(req.NewClient),
	route.Module,
	middleware.Module,
	controller.Module,
	service.Module,
	repository.Module,
	fx.Invoke(Run),
)

func NewLoggerOptions(config *config.APIConfig) logger.LoggerOptions {
	return logger.LoggerOptions{
		Level: config.LogLevel,
		Name:  "mini-api",
	}
}

func NewRedisOptions(config *config.APIConfig) lib.RedisOptions {
	return lib.RedisOptions{
		URL: config.RedisURL,
	}
}

func NewMySQLOptions(config *config.APIConfig) lib.MySQLOptions {
	return lib.MySQLOptions{
		DSN:      config.MySQLDSN,
		LogLevel: lib.NewMySQLLogLevel(config.LogLevel),
	}
}

func NewJWTOptions(config *config.APIConfig) lib.JWTOptions {
	return lib.JWTOptions{
		JWTSecret:     []byte(config.JWTSecret),
		JWTExpiresTTL: constant.TTLJWTExpires,
	}
}

func Run(
	lc fx.Lifecycle,
	routes route.Routes,
	handler *lib.HttpHandler,
	log *logger.Logger,
	cfg *config.APIConfig,
	redis *lib.Redis,
	mysql *lib.MySQL,
	svcPermission service.PermissionsService,
	svcRole service.RolesService,
	svcStaff service.StaffsService,
) {
	fp := fileprobe.NewHandler()
	isInitData, err := strconv.ParseBool(cfg.InitData)
	if err != nil {
		log.Error(err)
	}
	if isInitData {
		log.Debug("isInitData")
		svcPermission.InitPermissions()
		svcRole.InitRoles()
		svcStaff.InitStaffs()
	}
	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			go func() {
				routes.Setup()
				handler.Engine.HTTPErrorHandler = customHTTPErrorHandler(log)
				log.Infof(`http server listen on %s`, cfg.ListenAddr())
				if err := handler.Engine.Start(cfg.ListenAddr()); err != nil {
					if errors.Is(err, http.ErrServerClosed) {
						log.Warn("Shutting down...")
					} else {
						log.Fatalf("Failed to start: %s", err)
					}
				}
			}()
			return fp.Create()
		},
		OnStop: func(c context.Context) error {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				mysql.Close()
			}()
			go func() {
				defer wg.Done()
				redis.Close()
			}()
			wg.Wait()
			if err := handler.Engine.Shutdown(c); err != nil {
				log.Warn(err)
			}
			return fp.Remove()
		},
	})
}
