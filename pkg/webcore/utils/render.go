package utils

import (
	"github.com/k23dev/natalianatalia/app/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func RenderNotFound(c echo.Context) error {
	notfound := views.NotFound()
	return notfound.Render(c.Request().Context(), c.Response())
}
