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

func GetBooks(c echo.Context) error {
	db := db.DbManager()
	books := []models.Book{}
	db.Find(&books)
	for i, _ := range books {
		db.Model(books[i]).Related(&books[i].Category)
	}
	return c.JSON(http.StatusOK, response.OKWithData(books, "Data retrieved."))
}

func GetBook(c echo.Context) error {
	id := helpers.StringToInt(c.Param("id"))
	db := db.DbManager()
	book := models.Book{}

	res := db.Find(&book, &models.Book{ID: id}).Related(&book.Category)
	if res.RecordNotFound() || (id == 0) {
		return c.JSON(http.StatusNotFound, response.FailedMessage("Data not found."))
	}

	return c.JSON(http.StatusOK, response.OKWithData(book, "Data retrieved."))
}

func StoreBook(c echo.Context) error {
	errParams := map[string]string{}
	params := new(params.BookRequest)
	if err := c.Bind(params); err != nil {
		// fmt.Println("err", err)
		return c.JSON(http.StatusUnprocessableEntity, response.FailedMessage("invalid request parameters"))
	}

	if len(params.ISBN) == 0 {
		errParams["isbn"] = "cannot be empty"
	}

	if len(params.Publisher) == 0 {
		errParams["publisher"] = "cannot be empty"
	}

	if params.Price == 0 {
		errParams["price"] = "cannot be empty"
	}

	if len(params.Title) == 0 {
		errParams["title"] = "cannot be empty"
	}

	if params.CategoryID == 0 {
		errParams["category_id"] = "cannot be empty"
	}

	if len(errParams) > 0 {
		return c.JSON(http.StatusUnprocessableEntity, response.FailedWithData(errParams, "invalid request parameters"))
	}

	db := db.DbManager()

	var catId = params.CategoryID

	findCat := db.First(&models.Category{ID: catId})
	if findCat.Error != nil {
		return c.JSON(http.StatusNotFound, response.FailedMessage("Category not found."))
	}

	book := models.Book{
		ISBN:          params.ISBN,
		Publisher:     params.Publisher,
		Price:         params.Price,
		Title:         params.Title,
		Year:          params.Year,
		Author:        params.Author,
		CoverImage:    params.CoverImage,
		Description:   params.Description,
		PublishedDate: params.PublishedDate,
		CategoryID:    catId,
	}
	res := db.Create(&book)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.FailedWithData(res.Error, "Error inserting new data."))
	}

	return c.JSON(http.StatusCreated, response.OKWithData(res.Value, "Data created."))
}

func UpdateBook(c echo.Context) error {
	id := helpers.StringToInt(c.Param("id"))
	errParams := map[string]string{}
	params := new(params.BookRequest)
	if err := c.Bind(params); err != nil {
		// fmt.Println("err", err)
		return c.JSON(http.StatusUnprocessableEntity, response.FailedMessage("invalid request parameters"))
	}

	if len(params.ISBN) == 0 {
		errParams["isbn"] = "cannot be empty"
	}

	if len(params.Publisher) == 0 {
		errParams["publisher"] = "cannot be empty"
	}

	if params.Price == 0 {
		errParams["price"] = "cannot be empty"
	}

	if len(params.Title) == 0 {
		errParams["title"] = "cannot be empty"
	}

	if params.CategoryID == 0 {
		errParams["category_id"] = "cannot be empty"
	}

	if len(errParams) > 0 {
		return c.JSON(http.StatusUnprocessableEntity, response.FailedWithData(errParams, "invalid request parameters"))
	}

	db := db.DbManager()

	book := models.Book{ID: id}

	findBook := db.First(&book)

	if findBook.Error != nil {
		return c.JSON(http.StatusNotFound, response.FailedMessage("Data not found."))
	}

	var catId = params.CategoryID

	findCat := db.First(&models.Category{ID: catId})
	if findCat.Error != nil {
		return c.JSON(http.StatusNotFound, response.FailedMessage("Category not found."))
	}

	res := db.Model(&book).Updates(models.Book{
		ISBN:          params.ISBN,
		Publisher:     params.Publisher,
		Price:         params.Price,
		Title:         params.Title,
		Year:          params.Year,
		Author:        params.Author,
		CoverImage:    params.CoverImage,
		Description:   params.Description,
		PublishedDate: params.PublishedDate,
		CategoryID:    catId,
	})

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.FailedWithData(res.Error, "Error updating data."))
	}

	return c.JSON(http.StatusOK, response.OKWithData(res.Value, "Data updated."))
}

func DestroyBook(c echo.Context) error {
	id := helpers.StringToInt(c.Param("id"))

	db := db.DbManager()

	book := models.Book{ID: id}

	findCat := db.First(&book)

	if findCat.Error != nil {
		return c.JSON(http.StatusNotFound, response.FailedMessage("Data not found."))
	}

	res := db.Delete(&book)

	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.FailedWithData(res.Error, "Error deleting data."))
	}

	return c.JSON(http.StatusOK, response.OKWithData(res.Value, "Data deleted."))
}
