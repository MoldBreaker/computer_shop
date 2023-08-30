package routes

import (
	"computer_shop/config"
	"computer_shop/controllers"
	"computer_shop/middlewares"
	"github.com/labstack/echo/v4"
	"os"
)

var (
	ProductController controllers.ProductController
	UserController    controllers.UserController
	CartController    controllers.CartController
	RoleController    controllers.RoleController
	AuthMiddleware    middlewares.AuthMiddleware
)

func InitWebRoutes() {
	router := echo.New()
	config.LoadENV()

	router.Use(AuthMiddleware.Auth)

	router.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World")
	})

	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("/", ProductController.GetListProducts)
			products.GET("/:id", ProductController.GetProductById)
			products.POST("/", ProductController.CreateProduct, AuthMiddleware.IsLogined, AuthMiddleware.IsAdminOrSuperAdmin)
			products.PUT("/:id", ProductController.UpdateProduct, AuthMiddleware.IsLogined, AuthMiddleware.IsAdminOrSuperAdmin)
			products.DELETE("/:id", ProductController.DeleteProduct, AuthMiddleware.IsLogined, AuthMiddleware.IsAdminOrSuperAdmin)
		}

		users := api.Group("/users")
		{
			users.POST("/register", UserController.Register)
			users.POST("/login", UserController.Login)
			users.GET("/logout", UserController.Logout, AuthMiddleware.IsLogined)
			users.POST("/reset-password", UserController.ResetPassword, AuthMiddleware.IsLogined)
			users.POST("/avatar", UserController.ChangeAvatar, AuthMiddleware.IsLogined)
			users.POST("/info", UserController.UpdateInformation, AuthMiddleware.IsLogined)
		}

		carts := api.Group("/carts")
		{
			carts.GET("/:id", CartController.AddToCart)
			carts.GET("/update/:id", CartController.UpdateInCart)
			carts.DELETE("/:id", CartController.DeleteInCart)
		}

		role := api.Group("/role")
		{
			role.POST("/", RoleController.CreateRole, AuthMiddleware.IsLogined, AuthMiddleware.IsSuperAdmin)
		}

	}

	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
