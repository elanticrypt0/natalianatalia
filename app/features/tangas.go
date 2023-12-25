package features

import (
	"net/http"
	"strconv"

	"github.com/k23dev/natalianatalia/app/models"
	"github.com/k23dev/natalianatalia/app/views"
	"github.com/k23dev/natalianatalia/pkg/webcore"
	"github.com/k23dev/natalianatalia/pkg/webcore/utils"

	"github.com/labstack/echo/v4"
)

func FindOneTanga(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tan := models.NewTanga()
	tanga := tan.FindOne(tapp.App.DB.Primary, id)
	if tanga.ID != 0 {
		return utils.Render(c, views.TangasShowOne(tapp.GetTitleAndVersion(), *tanga))
	} else {
		return utils.RenderNotFound(c)
	}
}

func FindAllTangas(c echo.Context, tapp *webcore.TangoApp) error {
	tan := models.NewTanga()
	Tangas := tan.FindAll(tapp.App.DB.Primary)
	return utils.Render(c, views.TangasShowList(tapp.GetTitleAndVersion(), *Tangas))
}

func ShowFormTanga(c echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	tan := models.NewTanga()

	if is_new {
		return utils.Render(c, views.TangasFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(c.Param("id"))
		tan := tan.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(c, views.TangasFormUpdate(tapp.GetTitleAndVersion(), tan))
	}
}

func CreateTanga(c echo.Context, tapp *webcore.TangoApp) error {
	tanDTO := models.TangaDTO{}
	if err := c.Bind(&tanDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	tan := models.NewTanga()
	tanga := tan.Create(tapp.App.DB.Primary, tanDTO.Codename)

	return c.JSON(http.StatusOK, tanga)
}

func UpdateTanga(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tan := models.NewTanga()
	tanga := tan.FindOne(tapp.App.DB.Primary, id)
	c.Bind(&tan)
	tanga = tan.Update(tapp.App.DB.Primary, tanga)
	return c.JSON(http.StatusOK, tanga)
}

func DeleteTanga(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cat := models.NewTanga()
	tanga := cat.Delete(tapp.App.DB.Primary, id)
	return c.JSON(http.StatusOK, tanga)
}
