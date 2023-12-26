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

func FindOneTanga(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tan := models.NewTanga()
	tanga, _ := tan.FindOne(tapp.App.DB.Primary, id)
	if tanga.ID != 0 {
		return utils.Render(c, views.TangasShowOne(*tanga))
	} else {
		return utils.RenderNotFound(c, tapp.GetTitleAndVersion())
	}
}

func FindAllTangas(c echo.Context, tapp *webcore.TangoApp) error {
	// tan := models.NewTanga()
	// Tangas, _ := tan.FindAll(tapp.App.DB.Primary)

	// return utils.Render(c, views.TangasShowList(tapp.GetTitleAndVersion(), *Tangas))

	queryPage := c.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	tan := models.NewTanga()
	counter, _ := tan.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	tangas, _ := tan.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	return utils.Render(c, views.TangasShowList(tapp.GetTitleAndVersion(), *tangas, *pagination))

}

func ShowFormTanga(c echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	tan := models.NewTanga()

	if is_new {
		return utils.Render(c, views.TangasFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(c.Param("id"))
		tan, _ := tan.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(c, views.TangasFormUpdate(tapp.GetTitleAndVersion(), tan))
	}
}

func CreateTanga(c echo.Context, tapp *webcore.TangoApp) error {
	tanDTO := models.TangaDTO{}
	if err := c.Bind(&tanDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	tan := models.NewTanga()
	tan.Create(tapp.App.DB.Primary, tanDTO.Codename)

	return c.Redirect(http.StatusMovedPermanently, "/tangas/")
}

func UpdateTanga(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// get the incoming values
	tanDTO := models.TangaDTO{}
	if err := c.Bind(&tanDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	tan := models.NewTanga()
	tan.Update(tapp.App.DB.Primary, id, tanDTO.Codename)
	return c.Redirect(http.StatusMovedPermanently, "/tangas/")
}

func DeleteTanga(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cat := models.NewTanga()
	cat.Delete(tapp.App.DB.Primary, id)
	return c.Redirect(http.StatusMovedPermanently, "/tangas/")
}
