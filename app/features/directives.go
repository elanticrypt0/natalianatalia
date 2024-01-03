package features

import (
	"net/http"
	"strconv"

	"github.com/k23dev/natalianatalia/app/models"
	"github.com/k23dev/natalianatalia/app/views"
	"github.com/k23dev/natalianatalia/pkg/pagination"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/k23dev/natalianatalia/pkg/webcore/utils"
	"github.com/labstack/echo/v4"
)

func FindOneDirective(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	d := models.NewDirective()
	directive, _ := d.FindOne(tapp.App.DB.Primary, id)
	if directive != nil {
		return utils.Render(ctx, views.DirectivesShowOne(tapp.GetTitleAndVersion(), *directive))
	} else {
		return ctx.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAllDirectives(ctx echo.Context, tapp *webcore.TangoApp) error {
	queryPage := ctx.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	d := models.NewDirective()
	counter, _ := d.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	dBuf, _ := d.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	if dBuf != nil {
		return utils.Render(ctx, views.DirectivesShowList(tapp.GetTitleAndVersion(), *dBuf, *pagination))
	} else {
		return utils.Render(ctx, views.DirectivesShowListEmpty(tapp.GetTitleAndVersion()))
	}

}

func ShowFormDirective(ctx echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	d := models.NewDirective()

	list, _ := d.FindAll(tapp.App.DB.Primary)

	if is_new {
		return utils.Render(ctx, views.DirectivesFormCreate(tapp.GetTitleAndVersion(), &list))
	} else {
		id, _ := strconv.Atoi(ctx.Param("id"))
		d, _ := d.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(ctx, views.DirectivesFormUpdate(tapp.GetTitleAndVersion(), d, &list))
	}
}

func CreateDirective(ctx echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	dDTO := models.DirectiveDTO{}
	if err := ctx.Bind(&dDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	d := models.NewDirective()
	d.Create(tapp.App.DB.Primary, dDTO)

	return ctx.Redirect(http.StatusMovedPermanently, "/directives/")
}

func UpdateDirective(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get the incoming values
	dDTO := models.DirectiveDTO{}
	if err := ctx.Bind(&dDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	d := models.NewDirective()
	d.Update(tapp.App.DB.Primary, id, dDTO)

	return ctx.Redirect(http.StatusMovedPermanently, "/directives/")
}

func DeleteDirective(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	d := models.NewDirective()
	d.Delete(tapp.App.DB.Primary, id)

	return ctx.Redirect(http.StatusMovedPermanently, "/directives/")
}
