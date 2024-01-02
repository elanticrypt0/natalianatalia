package routes

import (
	"github.com/k23dev/natalianatalia/app/features"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func tangas_fieldsRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	tangas_fields := rootPath.Group("/tangas_fields/")

	tangas_fields.GET("", func(ctx echo.Context) error {
		return features.FindAllTangas_fields(ctx, tapp)
	})

	tangas_fields.GET(":id", func(ctx echo.Context) error {
		return features.FindOneTangas_field(ctx, tapp)
	})

	tangas_fields.GET("new", func(ctx echo.Context) error {
		return features.ShowFormTangas_field(ctx, tapp, true)
	})

	tangas_fields.GET("edit/:id", func(ctx echo.Context) error {
		return features.ShowFormTangas_field(ctx, tapp, false)
	})

	tangas_fields.POST("create", func(ctx echo.Context) error {
		return features.CreateTangas_field(ctx, tapp)
	})

	tangas_fields.POST("update/:id", func(ctx echo.Context) error {
		return features.UpdateTangas_field(ctx, tapp)
	})

	tangas_fields.GET("delete/:id", func(ctx echo.Context) error {
		return features.DeleteTangas_field(ctx, tapp)
	})
}
