package routes

import (
	"computer_shop/config"
	"computer_shop/controllers"
	"github.com/labstack/echo/v4"
	"os"
)

var (
	ProductController controllers.ProductController
	UserController    controllers.UserController
	CartController    controllers.CartController
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

		users := api.Group("/users")
		{
			users.POST("/register", UserController.Register)
			users.POST("/login", UserController.Login)
			users.GET("/logout", UserController.Logout)
		}

		carts := api.Group("/carts")
		{
			carts.GET("/:id", CartController.AddToCart)
			carts.GET("/update/:id", CartController.UpdateInCart)
			carts.DELETE("/:id", CartController.DeleteInCart)
		}

	}

	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
