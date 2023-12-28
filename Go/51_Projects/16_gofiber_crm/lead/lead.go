package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joeytatu/go-fiber-crm-basic/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetAllLeads(c *fiber.Ctx) {
	var leads []Lead

	db := database.DBConn
	db.Find(&leads)
	c.JSON(leads)
}

func GetLeadById(c *fiber.Ctx) {
	var lead Lead

	id := c.Params("id")
	db := database.DBConn
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	var lead Lead

	id := c.Params("id")
	db := database.DBConn
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(503).Send("No lead found with that ID")
	}
	db.Delete(&lead)
	c.Send("Lead sucessfully deleted")
}
