package handlers

import (
	"net/http"

	"github.com/ferdhika31/go-books/db"
	"github.com/ferdhika31/go-books/helpers"
	"github.com/ferdhika31/go-books/models"
	"github.com/ferdhika31/go-books/params"
	"github.com/ferdhika31/go-books/response"
	"github.com/labstack/echo"
)

func GetCategories(c echo.Context) error {
	db := db.DbManager()
	categories := []models.Category{}
	db.Find(&categories)
	return c.JSON(http.StatusOK, response.OKWithData(categories, "Data retrieved."))
}

func GetCategory(c echo.Context) error {
	id := helpers.StringToInt(c.Param("id"))
	db := db.DbManager()
	category := models.Category{}

	res := db.Find(&category, &models.Category{ID: id})
	if res.RecordNotFound() || (id == 0) {
		return c.JSON(http.StatusNotFound, response.FailedMessage("Data not found."))
	}

	return c.JSON(http.StatusOK, response.OKWithData(category, "Data retrieved."))
}

func StoreCategory(c echo.Context) error {
	errParams := map[string]string{}
	params := new(params.CategoryRequest)
	if err := c.Bind(params); err != nil {
		// fmt.Println("err", err)
		return c.JSON(http.StatusUnprocessableEntity, response.FailedMessage("invalid request parameters"))
	}

	if len(params.Name) == 0 {
		errParams["name"] = "cannot be empty"
	}

	if len(errParams) > 0 {
		return c.JSON(http.StatusUnprocessableEntity, response.FailedWithData(errParams, "invalid request parameters"))
	}

	db := db.DbManager()

	category := models.Category{
		Name: params.Name,
	}
	res := db.Create(&category)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.FailedWithData(res.Error, "Error inserting new data."))
	}

	return c.JSON(http.StatusCreated, response.OKWithData(res.Value, "Data created."))
}

func UpdateCategory(c echo.Context) error {
	id := helpers.StringToInt(c.Param("id"))
	errParams := map[string]string{}
	params := new(params.CategoryRequest)
	if err := c.Bind(params); err != nil {
		// fmt.Println("err", err)
		return c.JSON(http.StatusUnprocessableEntity, response.FailedMessage("invalid request parameters"))
	}

	if len(params.Name) == 0 {
		errParams["name"] = "cannot be empty"
	}

	if len(errParams) > 0 {
		return c.JSON(http.StatusUnprocessableEntity, response.FailedWithData(errParams, "invalid request parameters"))
	}

	db := db.DbManager()

	category := models.Category{ID: id}

	findCat := db.First(&category)

	if findCat.Error != nil {
		return c.JSON(http.StatusNotFound, response.FailedMessage("Data not found."))
	}

	res := db.Model(&category).Updates(models.Category{Name: params.Name})

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.FailedWithData(res.Error, "Error updating data."))
	}

	return c.JSON(http.StatusOK, response.OKWithData(res.Value, "Data updated."))
}

func DestroyCategory(c echo.Context) error {
	id := helpers.StringToInt(c.Param("id"))

	db := db.DbManager()

	category := models.Category{ID: id}

	findCat := db.First(&category)

	if findCat.Error != nil {
		return c.JSON(http.StatusNotFound, response.FailedMessage("Data not found."))
	}

	res := db.Delete(&category)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.FailedWithData(res.Error, "Error deleting data."))
	}

	return c.JSON(http.StatusOK, response.OKWithData(res.Value, "Data deleted."))
}
