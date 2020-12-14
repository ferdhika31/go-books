package main

import (
	"fmt"
	"net/http"

	"github.com/ferdhika31/go-books/api"
	"github.com/ferdhika31/go-books/config"
	"github.com/ferdhika31/go-books/db"
	"github.com/ferdhika31/go-books/response"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Welcome to the webserver.")

	db.Init()

	// instance
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// routing
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, response.OKMessage("REST API Catalog Go."))
	})
	api.CategoryGroup(e.Group("/v1/categories"))
	api.BookGroup(e.Group("/v1/books"))

	configuration, _ := config.LoadConfig()
	// e.Logger.Fatal(e.Start(":8081"))
	e.Start(configuration.API.Host)
}
