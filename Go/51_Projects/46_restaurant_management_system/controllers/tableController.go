package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joeytatu/restaurant-management-system/database"
	"github.com/joeytatu/restaurant-management-system/models"
	"go.mongodb.org/mongo-driver/bson"
)

var tableCollection = database.TableCollection

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		tableId := c.Param("table_id")
		var table models.Table

		err := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching table"})
			return
		}
		c.JSON(http.StatusOK, table)
	}
}

func GetTableById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		tableId := c.Param("table_id")
		var table models.Table

		err := tableCollection.FindOne(ctx, bson.M{"table_id": tableId}).Decode(&table)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching table by ID"})
			return
		}
		c.JSON(http.StatusOK, table)
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var table models.Table

		if err := c.ShouldBindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
			return
		}

		_, err := tableCollection.InsertOne(ctx, table)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating table"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Table created successfully"})
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		tableId := c.Param("table_id")
		var updatedTable models.Table

		if err := c.ShouldBindJSON(&updatedTable); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
			return
		}

		filter := bson.M{"table_id": tableId}
		update := bson.M{"$set": updatedTable}

		_, err := tableCollection.UpdateOne(ctx, filter, update)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating table"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Table updated successfully"})
	}
}
