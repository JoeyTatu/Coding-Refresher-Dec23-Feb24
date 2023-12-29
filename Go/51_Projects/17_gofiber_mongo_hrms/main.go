package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client   *mongo.Client
	Database *mongo.Database
}

type Employee struct {
	Id              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName       string             `json:"first_name"`
	LastName        string             `json:"last_name"`
	HourlyRate      float32            `json:"hourly_rate"`
	FullPartTime    string             `json:"full_part_time"`
	DateOfBirth     time.Time          `json:"date_of_birth"`
	StartDate       time.Time          `json:"start_date"`
	TerminationDate *time.Time         `json:"termination_date,omitempty"`
	WorkEmail       string             `json:"work_email"`
	Department      string             `json:"department"`
	PpsNumber       string             `json:"pps_number"`
	DocumentType    string             `json:"document_type"`
	DocumentNumber  string             `json:"document_number"`
}

var mongoIns MongoInstance

const (
	dbName   = "fiber-hrms"
	mongoUrl = "mongodb://localhost:27017/" + dbName
	timeout  = 30 * time.Second
)

func Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoUrl)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")
	mongoIns.Client = client
	mongoIns.Database = client.Database(dbName)

	return nil
}

func GetAllEmployees(c *fiber.Ctx) error {
	var employees []Employee

	query := bson.D{}
	cursor, err := mongoIns.Database.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer cursor.Close(c.Context())

	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(employees)
}

func CreateEmployee(c *fiber.Ctx) error {
	var employee Employee

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := mongoIns.Database.Collection("employees").InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"id": result.InsertedID})
}

func GetEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var employee Employee
	err = mongoIns.Database.Collection("employees").FindOne(c.Context(), bson.M{"_id": objID}).Decode(&employee)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}

	return c.JSON(employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	var updatedEmployee Employee
	if err := c.BodyParser(&updatedEmployee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := mongoIns.Database.Collection("employees").UpdateOne(
		c.Context(),
		bson.M{"_id": objID},
		bson.M{"$set": updatedEmployee},
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}

	return c.JSON(fiber.Map{"message": "Employee updated successfully"})
}

func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	result, err := mongoIns.Database.Collection("employees").DeleteOne(c.Context(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}

	return c.JSON(fiber.Map{"message": "Employee deleted successfully"})
}

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// Routes
	app.Get("/api/employees", GetAllEmployees)
	app.Post("/api/employees/employee", CreateEmployee)
	app.Get("/api/employees/employee/:id", GetEmployeeById)
	app.Put("/api/employees/employee/:id", UpdateEmployee)
	app.Delete("/api/employees/employee/:id", DeleteEmployee)

	log.Fatal(app.Listen(":3000"))
}
