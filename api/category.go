package api

import (
	"github.com/ferdhika31/go-books/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CategoryGroup(g *echo.Group) {
	g.Use(middleware.Logger())

	g.GET("", handlers.GetCategories)
	g.GET("/:id", handlers.GetCategory)
	g.POST("", handlers.StoreCategory)
	g.PATCH("/:id", handlers.UpdateCategory)
	g.DELETE("/:id", handlers.DestroyCategory)
}
