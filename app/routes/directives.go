package routes

import (
	"github.com/k23dev/natalianatalia/app/features"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/labstack/echo/v4"
)

func directivesRoutes(tapp *webcore.TangoApp, rootPath *echo.Group) {
	directives := rootPath.Group("/directives/")

	directives.GET("", func(ctx echo.Context) error {
		return features.FindAllDirectives(ctx, tapp)
	})

	directives.GET(":id", func(ctx echo.Context) error {
		return features.FindOneDirective(ctx, tapp)
	})

	directives.GET("new", func(ctx echo.Context) error {
		return features.ShowFormDirective(ctx, tapp, true)
	})

	directives.GET("edit/:id", func(ctx echo.Context) error {
		return features.ShowFormDirective(ctx, tapp, false)
	})

	directives.POST("create", func(ctx echo.Context) error {
		return features.CreateDirective(ctx, tapp)
	})

	directives.POST("update/:id", func(ctx echo.Context) error {
		return features.UpdateDirective(ctx, tapp)
	})

	directives.GET("delete/:id", func(ctx echo.Context) error {
		return features.DeleteDirective(ctx, tapp)
	})
}
