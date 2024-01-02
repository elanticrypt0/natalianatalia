package routes

import (
	"github.com/k23dev/natalianatalia/app/features"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func categoriesRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	categories := rootPath.Group("/categories/")

	categories.GET("", func(ctx echo.Context) error {
		return features.FindAllCategories(ctx, tapp)
	})

	categories.GET(":id", func(ctx echo.Context) error {
		return features.FindOneCategory(ctx, tapp)
	})

	categories.GET("new", func(ctx echo.Context) error {
		return features.ShowFormCategory(ctx, tapp, true)
	})

	categories.GET("edit/:id", func(ctx echo.Context) error {
		return features.ShowFormCategory(ctx, tapp, false)
	})

	categories.POST("create", func(ctx echo.Context) error {
		return features.CreateCategory(ctx, tapp)
	})

	categories.POST("update/:id", func(ctx echo.Context) error {
		return features.UpdateCategory(ctx, tapp)
	})

	categories.GET("delete/:id", func(ctx echo.Context) error {
		return features.DeleteCategory(ctx, tapp)
	})
}
