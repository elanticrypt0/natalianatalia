package features

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/k23dev/natalianatalia/app/models"
	"github.com/k23dev/natalianatalia/app/views"
	"github.com/k23dev/natalianatalia/pkg/pagination"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/k23dev/natalianatalia/pkg/webcore/utils"
	"github.com/labstack/echo/v4"
)

func FindOneTanga(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	t := models.NewTanga()
	tanga, _ := t.FindOne(tapp.App.DB.Primary, id)
	if tanga != nil {
		return utils.Render(ctx, views.TangasShowOne(tapp.GetTitleAndVersion(), *tanga))
	} else {
		return ctx.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAllTangas(ctx echo.Context, tapp *webcore.TangoApp) error {
	queryPage := ctx.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	t := models.NewTanga()
	counter, _ := t.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	tBuf, _ := t.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	if tBuf != nil {
		return utils.Render(ctx, views.TangasShowList(tapp.GetTitleAndVersion(), *tBuf, *pagination))
	} else {
		return utils.Render(ctx, views.TangasShowListEmpty(tapp.GetTitleAndVersion()))
	}

}

func ShowFormTanga(ctx echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	t := models.NewTanga()

	if is_new {
		return utils.Render(ctx, views.TangasFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(ctx.Param("id"))
		t, _ := t.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(ctx, views.TangasFormUpdate(tapp.GetTitleAndVersion(), t))
	}
}

func CreateTanga(ctx echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	tDTO := models.TangaDTO{}
	if err := ctx.Bind(&tDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	t := models.NewTanga()
	t.Create(tapp.App.DB.Primary, tDTO.Name, tDTO.Comment)

	return ctx.Redirect(http.StatusMovedPermanently, "/tangas/")
}

func UpdateTanga(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get the incoming values
	tDTO := models.TangaDTO{}
	if err := ctx.Bind(&tDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	t := models.NewTanga()
	t.Name = strings.ToLower(tDTO.Name)
	t.Comment = tDTO.Comment

	t.Update(tapp.App.DB.Primary, id, t.Name, t.Comment)

	return ctx.Redirect(http.StatusMovedPermanently, "/tangas/")
}

func DeleteTanga(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	t := models.NewTanga()
	t.Delete(tapp.App.DB.Primary, id)

	return ctx.Redirect(http.StatusMovedPermanently, "/tangas/")
}
