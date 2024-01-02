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

func FindOneScapp_param(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	s := models.NewScapp_param()
	scapp_param, _ := s.FindOne(tapp.App.DB.Primary, id)
	if scapp_param != nil {
		return utils.Render(ctx, views.Scapp_paramsShowOne(tapp.GetTitleAndVersion(), *scapp_param))
	} else {
		return ctx.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAllScapp_params(ctx echo.Context, tapp *webcore.TangoApp) error {
	queryPage := ctx.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	s := models.NewScapp_param()
	counter, _ := s.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	sBuf, _ := s.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	if sBuf != nil {
		return utils.Render(ctx, views.Scapp_paramsShowList(tapp.GetTitleAndVersion(), *sBuf, *pagination))
	} else {
		return utils.Render(ctx, views.Scapp_paramsShowListEmpty(tapp.GetTitleAndVersion()))
	}

}

func ShowFormScapp_param(ctx echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	s := models.NewScapp()
	list, _ := s.FindAll(tapp.App.DB.Primary)

	sp := models.NewScapp_param()

	if is_new {
		return utils.Render(ctx, views.Scapp_paramsFormCreate(tapp.GetTitleAndVersion(), &list))
	} else {
		id, _ := strconv.Atoi(ctx.Param("id"))
		spItem, _ := sp.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(ctx, views.Scapp_paramsFormUpdate(tapp.GetTitleAndVersion(), spItem, &list))
	}
}

func CreateScapp_param(ctx echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	sDTO := models.Scapp_paramDTO{}
	if err := ctx.Bind(&sDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	s := models.NewScapp_param()
	s.Create(tapp.App.DB.Primary, sDTO)

	return ctx.Redirect(http.StatusMovedPermanently, "/scapp_params/")
}

func UpdateScapp_param(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get the incoming values
	sDTO := models.Scapp_paramDTO{}
	if err := ctx.Bind(&sDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	s := models.NewScapp_param()

	s.Update(tapp.App.DB.Primary, id, sDTO)

	return ctx.Redirect(http.StatusMovedPermanently, "/scapp_params/")
}

func DeleteScapp_param(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	s := models.NewScapp_param()
	s.Delete(tapp.App.DB.Primary, id)

	return ctx.Redirect(http.StatusMovedPermanently, "/scapp_params/")
}
