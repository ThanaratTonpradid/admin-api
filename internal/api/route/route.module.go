package route

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewSwaggerRoute),
	fx.Provide(NewAuthRoute),
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	swagger SwaggerRoute,
	auth AuthRoute,
) Routes {
	return Routes{
		swagger,
		auth,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
