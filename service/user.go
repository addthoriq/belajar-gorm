package service

import (
	"belajar-gorm/common"
	"belajar-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	db := models.DB
	var users []models.User

	limit, offset, page := common.Paginate(c)

	var total int64
	num_count := db.Limit(limit).Offset(offset).Find(&users).Count(&total)

	if num_count.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": num_count.Error.Error()})
	}

	result := db.Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.JSON(fiber.Map{
		"data":       users,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": int((total + int64(limit) - 1) / int64(limit)),
	})
}
