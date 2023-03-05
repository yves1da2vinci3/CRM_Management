package lead

import (
	"gofibercrm/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn()
	var leads []Lead
	db.find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn()
	var lead Lead
	db.First(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn()
	lead := new(Lead)
	if err := c.BodyParser(&lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.Status(201).Send(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn()
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(400).Send("No lead found")
	}
	db.Delete(Lead{}, "id =?", id)
	c.Status(204).Send("Lead deleted")

}
