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

func FindOneTanga_field(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	t := models.NewTanga_field()
	tanga_field, _ := t.FindOne(tapp.App.DB.Primary, id)
	if tanga_field != nil {
		return utils.Render(ctx, views.Tanga_fieldsShowOne(tapp.GetTitleAndVersion(), *tanga_field))
	} else {
		return ctx.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAllTanga_fields(ctx echo.Context, tapp *webcore.TangoApp) error {
	queryPage := ctx.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	t := models.NewTanga_field()
	counter, _ := t.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	tBuf, _ := t.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	if tBuf != nil {
		return utils.Render(ctx, views.Tanga_fieldsShowList(tapp.GetTitleAndVersion(), *tBuf, *pagination))
	} else {
		return utils.Render(ctx, views.Tanga_fieldsShowListEmpty(tapp.GetTitleAndVersion()))
	}

}

func ShowFormTanga_field(ctx echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	t := models.NewTanga_field()

	tM := models.NewTanga()
	list, _ := tM.FindAll(tapp.App.DB.Primary)

	if is_new {
		return utils.Render(ctx, views.Tanga_fieldsFormCreate(tapp.GetTitleAndVersion(), &list))
	} else {
		id, _ := strconv.Atoi(ctx.Param("id"))
		t, _ := t.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(ctx, views.Tanga_fieldsFormUpdate(tapp.GetTitleAndVersion(), t, &list))
	}
}

func CreateTanga_field(ctx echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	tDTO := models.Tanga_fieldDTO{}
	if err := ctx.Bind(&tDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	t := models.NewTanga_field()
	t.Create(tapp.App.DB.Primary, tDTO)

	return ctx.Redirect(http.StatusMovedPermanently, "/tanga_fields/")
}

func UpdateTanga_field(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get the incoming values
	tDTO := models.Tanga_fieldDTO{}
	if err := ctx.Bind(&tDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	t := models.NewTanga_field()
	t.Update(tapp.App.DB.Primary, id, tDTO)

	return ctx.Redirect(http.StatusMovedPermanently, "/tanga_fields/")
}

func DeleteTanga_field(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	t := models.NewTanga_field()
	t.Delete(tapp.App.DB.Primary, id)

	return ctx.Redirect(http.StatusMovedPermanently, "/tanga_fields/")
}
