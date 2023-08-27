package routes

import (
	"computer_shop/config"
	"computer_shop/controllers"
	"github.com/labstack/echo/v4"
	"os"
)

var (
	ProductController controllers.ProductController
)

func InitWebRoutes() {
	router := echo.New()
	config.LoadENV()

	router.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World")
	})

	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("/", ProductController.GetListProducts)
			products.GET("/:id", ProductController.GetProductById)
			products.POST("/", ProductController.CreateProduct)
			products.PUT("/:id", ProductController.UpdateProduct)
			products.DELETE("/:id", ProductController.DeleteProduct)
		}
	}

	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
