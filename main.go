package main

import (
	"computer_shop/config"
	"computer_shop/routes"
	"os"
)

func main() {
	config.LoadENV()
	config.GetConnection()
	router := routes.InitWebRoutes()
	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}