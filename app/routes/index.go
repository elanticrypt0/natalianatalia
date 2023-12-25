package routes

import (
	"github.com/k23dev/tango/app/views"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/k23dev/tango/pkg/webcore/utils"
	"github.com/labstack/echo/v4"
)

func IndexRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {

	rootPath.GET("/", func(c echo.Context) error {
		return utils.Render(c, views.Index(tapp.App.Config.App_name+" V."+tapp.App.Config.App_version))
	})
}
