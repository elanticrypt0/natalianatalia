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

func FindOneTangas_field(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	t := models.NewTangas_field()
	tangas_field, _ := t.FindOne(tapp.App.DB.Primary, id)
	if tangas_field != nil {
		return utils.Render(ctx, views.Tangas_fieldsShowOne(tapp.GetTitleAndVersion(), *tangas_field))
	} else {
		return ctx.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAllTangas_fields(ctx echo.Context, tapp *webcore.TangoApp) error {
	queryPage := ctx.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	t := models.NewTangas_field()
	counter, _ := t.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	tBuf, _ := t.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	if tBuf != nil {
		return utils.Render(ctx, views.Tangas_fieldsShowList(tapp.GetTitleAndVersion(), *tBuf, *pagination))
	} else {
		return utils.Render(ctx, views.Tangas_fieldsShowListEmpty(tapp.GetTitleAndVersion()))
	}

}

func ShowFormTangas_field(ctx echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	t := models.NewTangas_field()

	// prepare tangas selector
	tangaM := models.NewTanga()
	tangasList, _ := tangaM.FindAll(tapp.App.DB.Primary)

	if is_new {
		return utils.Render(ctx, views.Tangas_fieldsFormCreate(tapp.GetTitleAndVersion(), &tangasList))
	} else {
		id, _ := strconv.Atoi(ctx.Param("id"))
		t, _ := t.FindOne(tapp.App.DB.Primary, id)

		return utils.Render(ctx, views.Tangas_fieldsFormUpdate(tapp.GetTitleAndVersion(), t, &tangasList))
	}
}

func CreateTangas_field(ctx echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	tDTO := models.Tangas_fieldDTO{}
	if err := ctx.Bind(&tDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	t := models.NewTangas_field()
	t.Create(tapp.App.DB.Primary, tDTO.Name, tDTO.FValue, tDTO.TangaID)

	return ctx.Redirect(http.StatusMovedPermanently, "/tangas_fields/")
}

func UpdateTangas_field(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get the incoming values
	tDTO := models.Tangas_fieldDTO{}
	if err := ctx.Bind(&tDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	t := models.NewTangas_field()
	t.Name = strings.ToLower(tDTO.Name)
	t.FValue = tDTO.FValue
	t.TangaID = tDTO.TangaID

	t.Update(tapp.App.DB.Primary, id, t.Name, t.FValue, t.TangaID)

	return ctx.Redirect(http.StatusMovedPermanently, "/tangas_fields/")
}

func DeleteTangas_field(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	t := models.NewTangas_field()
	t.Delete(tapp.App.DB.Primary, id)

	return ctx.Redirect(http.StatusMovedPermanently, "/tangas_fields/")
}
