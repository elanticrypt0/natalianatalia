package features

import (
	"net/http"
	"strconv"

	"github.com/k23dev/tango/app/models"
	"github.com/k23dev/tango/app/views"
	"github.com/k23dev/tango/pkg/webcore"
	"github.com/k23dev/tango/pkg/webcore/utils"
	"github.com/labstack/echo/v4"
)

func FindOneCategory(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))

	cat := models.NewCategory()
	category := cat.FindOne(tapp.App.DB.Primary, id)
	if category.ID != 0 {
		return utils.Render(c, views.CategoriesShowOne(*category))
	} else {
		return utils.RenderNotFound(c)
	}
}

func FindAllCategories(c echo.Context, tapp *webcore.TangoApp) error {
	cat := models.NewCategory()
	categories := cat.FindAll(tapp.App.DB.Primary)
	return utils.Render(c, views.CategoriesShowList(categories))
}

func ShowFormCategory(c echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	cat := models.NewCategory()

	if is_new {
		return utils.Render(c, views.CategoriesFormCreate())
	} else {
		id, _ := strconv.Atoi(c.Param("id"))
		cat := cat.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(c, views.CategoriesFormUpdate(cat))
	}
}

func CreateCategory(c echo.Context, tapp *webcore.TangoApp) error {
	cat := models.NewCategory()
	c.Bind(&cat)
	category := cat.Create(tapp.App.DB.Primary, cat.Name)
	return c.JSON(http.StatusOK, category)
}

func UpdateCategory(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cat := models.NewCategory()
	category := cat.FindOne(tapp.App.DB.Primary, id)
	c.Bind(&cat)
	category = cat.Update(tapp.App.DB.Primary, *category)
	return c.JSON(http.StatusOK, category)
}

func DeleteCategory(c echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cat := models.NewCategory()
	category := cat.Delete(tapp.App.DB.Primary, id)
	return c.JSON(http.StatusOK, category)
}
