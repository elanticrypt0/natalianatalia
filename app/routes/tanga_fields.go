package routes

import (
	"github.com/k23dev/natalianatalia/app/features"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func tanga_fieldsRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	tanga_fields := rootPath.Group("/tanga_fields/")

	tanga_fields.GET("", func(ctx echo.Context) error {
		return features.FindAllTanga_fields(ctx, tapp)
	})

	tanga_fields.GET(":id", func(ctx echo.Context) error {
		return features.FindOneTanga_field(ctx, tapp)
	})

	tanga_fields.GET("new", func(ctx echo.Context) error {
		return features.ShowFormTanga_field(ctx, tapp, true)
	})

	tanga_fields.GET("edit/:id", func(ctx echo.Context) error {
		return features.ShowFormTanga_field(ctx, tapp, false)
	})

	tanga_fields.POST("create", func(ctx echo.Context) error {
		return features.CreateTanga_field(ctx, tapp)
	})

	tanga_fields.POST("update/:id", func(ctx echo.Context) error {
		return features.UpdateTanga_field(ctx, tapp)
	})

	tanga_fields.GET("delete/:id", func(ctx echo.Context) error {
		return features.DeleteTanga_field(ctx, tapp)
	})
}
