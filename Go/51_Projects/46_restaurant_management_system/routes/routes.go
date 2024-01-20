package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joeytatu/restaurant-management-system/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controllers.GetFoodById())
	incomingRoutes.POST("/foods", controllers.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controllers.UpdateFood())
	incomingRoutes.PUT("/foods/:food_id", controllers.UpdateFood())
}

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoiceById())
	incomingRoutes.POST("/invoices", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoices/:invoice_id", controllers.UpdateInvoice())
}

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controllers.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controllers.GetMenuById())
	incomingRoutes.POST("/menus", controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
	incomingRoutes.PUT("/menus/:menu_id", controllers.UpdateMenu())
}

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order-items", controllers.GetOrderItems())
	incomingRoutes.GET("/order-items/:order-item_id", controllers.GetOrderItemById())
	incomingRoutes.POST("/order-items", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/order-items/:order-item_id", controllers.UpdateOrderItem())
	incomingRoutes.PUT("/order-items/:order-item_id", controllers.UpdateOrderItem())

	incomingRoutes.GET("/order-items_order/:order_id", controllers.GetOrderItemByOrder())
}

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controllers.GetOrders())
	incomingRoutes.GET("/orders/:order_id", controllers.GetOrderById())
	incomingRoutes.POST("/orders", controllers.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controllers.UpdateOrder())
	incomingRoutes.PUT("/orders/:order_id", controllers.UpdateOrder())
}

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTableById())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())
	incomingRoutes.PUT("/tables/:table_id", controllers.UpdateTable())
}

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUserById())
	incomingRoutes.POST("/users/signup", controllers.CreateUser())
	incomingRoutes.POST("/users/login", controllers.LoginUser())
	incomingRoutes.PATCH("/users/:user_id", controllers.UpdateUser())
	incomingRoutes.PUT("/users/:user_id", controllers.UpdateUser())
}
