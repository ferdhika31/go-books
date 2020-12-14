package api

import (
	"github.com/ferdhika31/go-books/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func BookGroup(g *echo.Group) {
	g.Use(middleware.Logger())

	g.GET("", handlers.GetBooks)
	g.GET("/:id", handlers.GetBook)
	g.POST("", handlers.StoreBook)
	g.PATCH("/:id", handlers.UpdateBook)
	g.DELETE("/:id", handlers.DestroyBook)
}
