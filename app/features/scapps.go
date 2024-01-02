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

func FindOneScapp(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	s := models.NewScapp()
	scapp, _ := s.FindOne(tapp.App.DB.Primary, id)
	if scapp != nil {
		return utils.Render(c, views.ScappsShowOne(tapp.GetTitleAndVersion(), *scapp))
	} else {
		return c.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAllScapps(c echo.Context, tapp *webcore.TangoApp) error {
	queryPage := c.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	s := models.NewScapp()
	counter, _ := s.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	sBuf, _ := s.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	if sBuf != nil {
		return utils.Render(c, views.ScappsShowList(tapp.GetTitleAndVersion(), *sBuf, *pagination))
	} else {
		return utils.Render(c, views.ScappsShowListEmpty(tapp.GetTitleAndVersion()))
	}
}

func ShowFormScapp(c echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	s := models.NewScapp()

	if is_new {
		return utils.Render(c, views.ScappsFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(c.Param("id"))
		s, _ := s.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(c, views.ScappsFormUpdate(tapp.GetTitleAndVersion(), s))
	}
}

func CreateScapp(c echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	sDTO := models.ScappDTO{}
	if err := c.Bind(&sDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	s := models.NewScapp()
	s.Create(tapp.App.DB.Primary, sDTO)

	return c.Redirect(http.StatusMovedPermanently, "/scapps/")
}

func UpdateScapp(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// get the incoming values
	sDTO := models.ScappDTO{}
	if err := c.Bind(&sDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	s := models.NewScapp()
	s.Name = strings.ToLower(sDTO.Name)
	s.Path = sDTO.Path
	s.IsSudo = sDTO.IsSudo
	s.Comment = sDTO.Comment

	s.Update(tapp.App.DB.Primary, id, s.Name, s.Path, s.IsSudo, s.Comment)

	return c.Redirect(http.StatusMovedPermanently, "/scapps/")
}

func DeleteScapp(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	s := models.NewScapp()
	s.Delete(tapp.App.DB.Primary, id)

	return c.Redirect(http.StatusMovedPermanently, "/scapps/")
}
