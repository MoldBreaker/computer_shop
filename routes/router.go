package routes

import (
	"computer_shop/config"
	"computer_shop/controllers"
	"computer_shop/middlewares"
	"os"

	"github.com/labstack/echo/v4"
)

var (
	ProductController      controllers.ProductController
	UserController         controllers.UserController
	CartController         controllers.CartController
	RoleController         controllers.RoleController
	InvoiceController      controllers.InvoiceController
	HomeController         controllers.HomeController
	NotificationController controllers.NotificationController
	CategoryController     controllers.CategoryController
	AuthMiddleware         middlewares.AuthMiddleware
	AuthRedirect           middlewares.AuthRedirect
	CheckCartMiddleware    middlewares.CheckCartMiddleware
	SuperAdminController   controllers.SuperAdminController
)

func InitWebRoutes() {
	router := echo.New()
	config.LoadENV()
	router.Static("/", "assets")

	router.Use(AuthMiddleware.Auth)

	router.GET("/", HomeController.RenderHomePage)
	router.GET("/auth", HomeController.RenderAuthPage)
	router.GET("/product/detail/:id", HomeController.RenderProductDetailPage)
	router.GET("/cart", HomeController.RenderCartPage, AuthRedirect.IsLogined)
	router.GET("/profile/:id", HomeController.RenderProfilePage)
	router.GET("/checkout", HomeController.RenderCheckoutPage, AuthMiddleware.IsLogined, CheckCartMiddleware.CheckCartEmpty)

	router.GET("/dashboard", SuperAdminController.RenderSuperAminPage, AuthMiddleware.IsLogined, AuthMiddleware.IsSuperAdmin)

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
			users.GET("/", UserController.GetAllUsers, AuthMiddleware.IsLogined, AuthMiddleware.IsSuperAdmin)
			users.GET("/:id", UserController.GetUserById)
			users.PUT("/block/:id", UserController.BlockUser, AuthMiddleware.IsLogined, AuthMiddleware.IsSuperAdmin)
		}

		carts := api.Group("/carts")
		{
			carts.GET("/", CartController.GetItemsInCart, AuthMiddleware.IsLogined)
			carts.GET("/:id", CartController.AddToCart, AuthMiddleware.IsLogined)
			carts.GET("/update/:id", CartController.UpdateInCart, AuthMiddleware.IsLogined)
			carts.DELETE("/:id", CartController.DeleteInCart, AuthMiddleware.IsLogined)
		}

		role := api.Group("/role")
		{
			role.GET("/", RoleController.GettAllRoles)
			role.POST("/", RoleController.CreateRole, AuthMiddleware.IsLogined, AuthMiddleware.IsSuperAdmin)
			role.PUT("/:id", RoleController.UpdateUserRole, AuthMiddleware.IsLogined, AuthMiddleware.IsSuperAdmin)
		}

		invoices := api.Group("/invoices")
		{
			invoices.POST("/", InvoiceController.CreateInvoice, AuthMiddleware.IsLogined)
			invoices.GET("/", InvoiceController.GetHistoryInvoices, AuthMiddleware.IsLogined)
			invoices.GET("/:id", InvoiceController.GetHistoryInvoiceDetails, AuthMiddleware.IsLogined)
		}

		notifications := api.Group("/notifications")
		{
			notifications.GET("/", NotificationController.GetAllNotifications, AuthMiddleware.IsLogined)
			notifications.DELETE("/:id", NotificationController.DelateNotification, AuthMiddleware.IsLogined)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", CategoryController.GetAllCategory)
		}
	}

	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
