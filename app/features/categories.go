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

func FindOneCategory(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	c := models.NewCategory()
	category, _ := c.FindOne(tapp.App.DB.Primary, id)
	if category != nil {
		return utils.Render(ctx, views.CategoriesShowOne(tapp.GetTitleAndVersion(), *category))
	} else {
		return ctx.Redirect(http.StatusMovedPermanently, "/404")
	}
}

func FindAllCategories(ctx echo.Context, tapp *webcore.TangoApp) error {
	queryPage := ctx.Param("page")
	var currentPage = 0
	if queryPage != "" {
		currentPage, _ = strconv.Atoi(queryPage)
	}

	c := models.NewCategory()
	counter, _ := c.Count(tapp.App.DB.Primary)
	pagination := pagination.NewPagination(currentPage, itemsPerPage, counter)
	cBuf, _ := c.FindAllPagination(tapp.App.DB.Primary, itemsPerPage, currentPage)

	if cBuf != nil {
		return utils.Render(ctx, views.CategoriesShowList(tapp.GetTitleAndVersion(), *cBuf, *pagination))
	} else {
		return utils.Render(ctx, views.CategoriesShowListEmpty(tapp.GetTitleAndVersion()))
	}

}

func ShowFormCategory(ctx echo.Context, tapp *webcore.TangoApp, is_new bool) error {
	c := models.NewCategory()

	if is_new {
		return utils.Render(ctx, views.CategoriesFormCreate(tapp.GetTitleAndVersion()))
	} else {
		id, _ := strconv.Atoi(ctx.Param("id"))
		c, _ := c.FindOne(tapp.App.DB.Primary, id)
		return utils.Render(ctx, views.CategoriesFormUpdate(tapp.GetTitleAndVersion(), c))
	}
}

func CreateCategory(ctx echo.Context, tapp *webcore.TangoApp) error {
	// get the incoming values
	cDTO := models.CategoryDTO{}
	if err := ctx.Bind(&cDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	c := models.NewCategory()
	c.Create(tapp.App.DB.Primary, cDTO.Name)

	return ctx.Redirect(http.StatusMovedPermanently, "/categories/")
}

func UpdateCategory(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get the incoming values
	cDTO := models.CategoryDTO{}
	if err := ctx.Bind(&cDTO); err != nil {
		return ctx.String(http.StatusBadRequest, "Bad request")
	}

	c := models.NewCategory()
	c.Name = strings.ToLower(cDTO.Name)

	c.Update(tapp.App.DB.Primary, id, c.Name)

	return ctx.Redirect(http.StatusMovedPermanently, "/categories/")
}

func DeleteCategory(ctx echo.Context, tapp *webcore.TangoApp) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	c := models.NewCategory()
	c.Delete(tapp.App.DB.Primary, id)

	return ctx.Redirect(http.StatusMovedPermanently, "/categories/")
}
