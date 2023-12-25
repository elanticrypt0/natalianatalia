package routes

import (
	"github.com/k23dev/natalianatalia/app/features"
	"github.com/k23dev/natalianatalia/pkg/webcore"

	"github.com/labstack/echo/v4"
)

func scriptsRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	scripts := rootPath.Group("/scripts/")

	scripts.GET("", func(c echo.Context) error {
		return features.FindAllCategories(c, tapp)
	})

	scripts.GET(":id", func(c echo.Context) error {
		return features.FindOneCategory(c, tapp)
	})

	scripts.GET("new", func(c echo.Context) error {
		return features.ShowFormCategory(c, tapp, true)
	})

	scripts.GET("edit/:id", func(c echo.Context) error {
		return features.ShowFormCategory(c, tapp, false)
	})

	scripts.POST("", func(c echo.Context) error {
		return features.CreateCategory(c, tapp)
	})

	scripts.PUT(":id", func(c echo.Context) error {
		return features.UpdateCategory(c, tapp)
	})

	scripts.DELETE(":id", func(c echo.Context) error {
		return features.DeleteCategory(c, tapp)
	})
}
