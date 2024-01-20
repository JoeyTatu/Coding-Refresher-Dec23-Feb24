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

var invoiceCollection = database.InvoiceCollection

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Param("invoice_id")
		var invoice models.Invoice

		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching invoice"})
			return
		}
		c.JSON(http.StatusOK, invoice)
	}
}

func GetInvoiceById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Param("invoice_id")
		var invoice models.Invoice

		err := invoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching invoice by ID"})
			return
		}
		c.JSON(http.StatusOK, invoice)
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var invoice models.Invoice

		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
			return
		}

		_, err := invoiceCollection.InsertOne(ctx, invoice)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while creating invoice"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Invoice created successfully"})
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		invoiceId := c.Param("invoice_id")
		var updatedInvoice models.Invoice

		if err := c.ShouldBindJSON(&updatedInvoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
			return
		}

		filter := bson.M{"invoice_id": invoiceId}
		update := bson.M{"$set": updatedInvoice}

		_, err := invoiceCollection.UpdateOne(ctx, filter, update)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while updating invoice"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Invoice updated successfully"})
	}
}
