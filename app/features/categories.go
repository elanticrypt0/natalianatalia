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

func FindOneCategory(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	cat := models.NewCategory()
	category := cat.FindOne(tapp.App.DB.Primary, id)
	if category.ID != 0 {
		return utils.Render(c, views.CategoriesShowOne(tapp.GetTitleAndVersion(), *category))
	} else {
		return utils.RenderNotFound(c)
	}
}

func FindAllCategories(c echo.Context, tapp *webcore.TangoApp) error {
	cat := models.NewCategory()
	categories := cat.FindAll(tapp.App.DB.Primary)
	return utils.Render(c, views.CategoriesShowList(tapp.GetTitleAndVersion(), categories))
}

func ShowFormCategory(c echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	cat := models.NewCategory()

	if is_new {
		return utils.Render(c, views.CategoriesFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(c.Param("id"))
		cat := cat.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(c, views.CategoriesFormUpdate(tapp.GetTitleAndVersion(), cat))
	}
}

func CreateCategory(c echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	catDTO := models.CategoryDTO{}
	if err := c.Bind(&catDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	cat := models.NewCategory()
	category := cat.Create(tapp.App.DB.Primary, catDTO.Name)
	return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// get the incoming values
	catDTO := models.CategoryDTO{}
	if err := c.Bind(&catDTO); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	cat := models.NewCategory()
	category := cat.FindOne(tapp.App.DB.Primary, id)

	category.Name = catDTO.Name

	category = cat.Update(tapp.App.DB.Primary, *category)
	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cat := models.NewCategory()
	category := cat.Delete(tapp.App.DB.Primary, id)
	return c.JSON(http.StatusOK, category)
}
