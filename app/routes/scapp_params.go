package routes

import (
	"github.com/k23dev/natalianatalia/app/features"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func scapp_paramsRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	scapp_params := rootPath.Group("/scapp_params/")

	scapp_params.GET("", func(ctx echo.Context) error {
		return features.FindAllScapp_params(ctx, tapp)
	})

	scapp_params.GET(":id", func(ctx echo.Context) error {
		return features.FindOneScapp_param(ctx, tapp)
	})

	scapp_params.GET("new", func(ctx echo.Context) error {
		return features.ShowFormScapp_param(ctx, tapp, true)
	})

	scapp_params.GET("edit/:id", func(ctx echo.Context) error {
		return features.ShowFormScapp_param(ctx, tapp, false)
	})

	scapp_params.POST("create", func(ctx echo.Context) error {
		return features.CreateScapp_param(ctx, tapp)
	})

	scapp_params.POST("update/:id", func(ctx echo.Context) error {
		return features.UpdateScapp_param(ctx, tapp)
	})

	scapp_params.GET("delete/:id", func(ctx echo.Context) error {
		return features.DeleteScapp_param(ctx, tapp)
	})
}
