package utils

import (
	"github.com/a-h/templ"
	"github.com/k23dev/tango/app/views"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func RenderNotFound(c echo.Context) error {
	notfound := views.NotFound()
	return notfound.Render(c.Request().Context(), c.Response())
}
