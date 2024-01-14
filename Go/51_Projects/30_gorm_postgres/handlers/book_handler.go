package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joeytatu/gorm-postgres/models"
	"gorm.io/gorm"
)

type Book struct {
	Barcode   string `json:"barcode"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	PrintDate string `json:"print_date"`
	Rerelease bool   `json:"rerelease"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetAllBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "books fetched successfully",
			"data":    bookModels,
		},
	)
	return nil
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := models.Books{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty"},
		)
		return nil
	}

	fmt.Println("The ID is", id)

	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "unable to find book"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "found book id",
			"data":    bookModel,
		},
	)
	return nil
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"},
		)
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "book created"},
	)
	return nil
}

func (r *Repository) UpdateBook(context *fiber.Ctx) error {
	id := context.Params("id")
	book := models.Books{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty"},
		)
		return nil
	}

	err := r.DB.Where("id = ?", id).First(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "unable to find book"},
		)
		return err
	}

	if err := context.BodyParser(&book); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"},
		)
		return err
	}

	if err := r.DB.Save(&book).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "unable to update book"},
		)
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "book updated successfully"},
	)
	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Books{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id cannot be empty"},
		)
		return nil
	}

	err := r.DB.Delete(bookModel, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "unable to delete book"},
		)
		return err.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "successfully deleted book"},
	)
	return nil
}

func (r *Repository) SetUpRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/books", r.GetAllBooks)
	api.Get("/books/book/:id", r.GetBookByID)
	api.Post("/books/new", r.CreateBook)
	api.Put("/books/book/:id", r.UpdateBook)
	api.Delete("/books/delete/:id", r.DeleteBook)
}
