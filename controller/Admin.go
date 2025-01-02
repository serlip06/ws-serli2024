package controller

import(
	"github.com/gofiber/fiber/v2"

    cek "github.com/serlip06/ujicobapackage/module"
	inimodel "github.com/serlip06/ujicobapackage/model"
)

func LoginAdmin(c *fiber.Ctx) error {
	var input inimodel.Admin

	// Parse JSON body to Admin struct
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	// Validate credentials
	if cek.ValidateAdmin(input) {
		return c.JSON(fiber.Map{
			"message": "Login successful",
			"status":  "success",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Invalid username or password",
	})
}