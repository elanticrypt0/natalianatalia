package routes

import (
	"github.com/k23dev/natalianatalia/app/features"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func tangasRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	tangas := rootPath.Group("/tangas/")

	tangas.GET("", func(ctx echo.Context) error {
		return features.FindAllTangas(ctx, tapp)
	})

	tangas.GET(":id", func(ctx echo.Context) error {
		return features.FindOneTanga(ctx, tapp)
	})

	tangas.GET("new", func(ctx echo.Context) error {
		return features.ShowFormTanga(ctx, tapp, true)
	})

	tangas.GET("edit/:id", func(ctx echo.Context) error {
		return features.ShowFormTanga(ctx, tapp, false)
	})

	tangas.POST("create", func(ctx echo.Context) error {
		return features.CreateTanga(ctx, tapp)
	})

	tangas.POST("update/:id", func(ctx echo.Context) error {
		return features.UpdateTanga(ctx, tapp)
	})

	tangas.GET("delete/:id", func(ctx echo.Context) error {
		return features.DeleteTanga(ctx, tapp)
	})
}
