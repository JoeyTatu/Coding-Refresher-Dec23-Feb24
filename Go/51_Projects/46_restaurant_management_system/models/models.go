package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	Price      *float64           `json:"price" validate:"required"`
	Food_image *string            `json:"food_image" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Food_id    string             `json:"food_id"`
	Menu_id    *string            `json:"menu_id" validate:"required"`
}

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       string             `json:"invoice_id"`
	Order_id         string             `json:"order_id"`
	Payment_method   *string            `json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
	Payment_status   *string            `json:"payment_status" validate:"eq=PENDING|eq=PAID"`
	Payment_due_date time.Time          `json:"payment_due_date"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
}

type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required"`
	Category   string             `json:"category" validate:"required"`
	Start_date *time.Time         `json:"start_date"`
	End_date   *time.Time         `json:"end_date"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Menu_id    string             `json:"menu_id"`
}

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

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Table_id   string             `json:"table_id" validate:"required"`
	Status     string             `json:"status" validate:"eq=OPEN|eq=CLOSED"`
	Total_cost float64            `json:"total_cost"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Order_id   string             `json:"order_id"`
}

type Table struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required"`
	Capacity   int                `json:"capacity" validate:"required,min=1"`
	Status     string             `json:"status" validate:"eq=AVAILABLE|eq=OCCUPIED"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Table_id   string             `json:"table_id"`
}

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	Username   string             `json:"username" validate:"required,min=3,max=50"`
	Email      string             `json:"email" validate:"required,email"`
	Password   string             `json:"password" validate:"required,min=6"`
	Role       string             `json:"role" validate:"eq=ADMIN|eq=STAFF"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	User_id    string             `json:"user_id"`
}
