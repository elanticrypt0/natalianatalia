package routes

import (
	"github.com/k23dev/natalianatalia/app/views"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/k23dev/natalianatalia/pkg/webcore/utils"
	"github.com/labstack/echo/v4"
)

func IndexRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {

	rootPath.GET("/", func(c echo.Context) error {
		return utils.Render(c, views.Index(tapp.GetTitleAndVersion()))
	})

	rootPath.GET("404", func(c echo.Context) error {
		return utils.RenderNotFound(c, tapp.GetTitleAndVersion())
	})
}
