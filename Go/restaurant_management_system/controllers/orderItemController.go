package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joeytatu/restaurant-management-system/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var OrderItemCollection = database.OrderItemCollection

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Food_id       string             `json:"food_id" validate:"required"`
	Quantity      int                `json:"quantity" validate:"required,min=1"`
	Unit_price    float64            `json:"unit_price" validate:"required"`
	Total_price   float64            `json:"total_price"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	Order_item_id string             `json:"order_item_id"`
}

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderItemId := c.Param("order_item_id")
		var orderItem OrderItem

		err := OrderItemCollection.FindOne(ctx, bson.M{"order_item_id": orderItemId}).Decode(&orderItem)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching order item"})
			return
		}
		c.JSON(http.StatusOK, orderItem)
	}
}

func GetOrderItemById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderItemId := c.Param("order_item_id")
		var orderItem OrderItem

		err := OrderItemCollection.FindOne(ctx, bson.M{"order_item_id": orderItemId}).Decode(&orderItem)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching order item by ID"})
			return
		}
		c.JSON(http.StatusOK, orderItem)
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var orderItem OrderItem

		if err := c.ShouldBindJSON(&orderItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
			return
		}

		_, err := OrderItemCollection.InsertOne(ctx, orderItem)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating order item"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Order item created successfully"})
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderItemId := c.Param("order_item_id")
		var updatedOrderItem OrderItem

		if err := c.ShouldBindJSON(&updatedOrderItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
			return
		}

		filter := bson.M{"order_item_id": orderItemId}
		update := bson.M{"$set": updatedOrderItem}

		_, err := OrderItemCollection.UpdateOne(ctx, filter, update)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating order item"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Order item updated successfully"})
	}
}

func GetOrderItemByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("order_id")

		orderItems, err := ItemByOrder(orderId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching order items by order"})
			return
		}

		c.JSON(http.StatusOK, orderItems)
	}
}

func ItemByOrder(orderId string) (OrderItems []primitive.M, err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	cursor, err := OrderItemCollection.Find(ctx, bson.M{"order_id": orderId})
	defer cancel()
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &OrderItems); err != nil {
		return nil, err
	}

	return OrderItems, nil
}
