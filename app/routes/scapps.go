package routes

import (
	"github.com/k23dev/natalianatalia/app/features"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func scappsRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	scapps := rootPath.Group("/scapps/")

	scapps.GET("", func(c echo.Context) error {
		return features.FindAllScapps(c, tapp)
	})

	scapps.GET(":id", func(c echo.Context) error {
		return features.FindOneScapp(c, tapp)
	})

	scapps.GET("new", func(c echo.Context) error {
		return features.ShowFormScapp(c, tapp, true)
	})

	scapps.GET("edit/:id", func(c echo.Context) error {
		return features.ShowFormScapp(c, tapp, false)
	})

	scapps.POST("create", func(c echo.Context) error {
		return features.CreateScapp(c, tapp)
	})

	scapps.POST("update/:id", func(c echo.Context) error {
		return features.UpdateScapp(c, tapp)
	})

	scapps.GET("delete/:id", func(c echo.Context) error {
		return features.DeleteScapp(c, tapp)
	})
}
