package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joeytatu/restaurant-management-system/middleware"
	"github.com/joeytatu/restaurant-management-system/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)

	router.Run(":" + port)
}
