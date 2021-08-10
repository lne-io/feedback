package controllers

import (
	"github.com/lne-io/feedback/models"
	"github.com/lne-io/feedback/database"
	"github.com/lne-io/feedback/utils"

	"github.com/gofiber/fiber/v2"
)

// Get all websites with a limit of 20 per page
func GetWebsites(c *fiber.Ctx) error {
	db := database.DB
	var website models.Website
	var websites []models.Website
	var totalWebsites int64

	page := c.Params("page")

	countResult := db.Model(&website).Count(&totalWebsites) 
	if countResult.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to count total websites",
			"error": countResult.Error,
		})
	}

	pagination := utils.NewPagination(totalWebsites, 20, page) 
	if pagination.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok": false,
			"message": "Pagination error",
			"error": pagination.Error,
		})
	}

	if result := db.Scopes(pagination.Paginate()).Find(&websites); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to get websites",
			"error": result.Error,
		})
	}
	
	return c.JSON(fiber.Map{
		"ok": true,
		"websites": websites,
		"pagination": pagination,
	})
}

// Get one website by ID
func GetWebsiteById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var website models.Website
	var feedback models.Feedback
	var totalFeedback int64

	if result := db.First(&website, "id = ?", id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to get website",
			"error": result.Error,
		})
	}

	countResult := db.Model(&feedback).Where("website_id = ?", website.ID).Count(&totalFeedback) 
	if countResult.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to count website's total Feedback messages",
			"error": countResult.Error,
		})
	}

	return c.JSON(fiber.Map{
		"ok": true,
		"website": website,
		"totalFeedback": totalFeedback,
	})
}

// Get website's feedback message with a limit of 20 per page
func GetWebsiteFeedback(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var website models.Website
	var feedback []models.Feedback 
	var totalFeedback int64

	page := c.Params("page")

	if result := db.First(&website, "id = ?", id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to get website",
			"error": result.Error,
		})
	}

	totalFeedback = db.Model(&website).Association("Feedback").Count()

	pagination := utils.NewPagination(totalFeedback, 20, page) 
	if pagination.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Pagination error",
			"error": pagination.Error,
		})
	}
	
	if res := db.Where("website_id = ?", website.ID).Scopes(pagination.Paginate()).Find(&feedback); res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to get website feedback",
			"error": res.Error,
		})
	}

	return c.JSON(fiber.Map{
		"ok": true,
		"website": website,
		"feedback": feedback,
		"pagination": pagination,
	})
}

// Getting a single website's feedback 
func GetWebsiteSingleFeedback(c *fiber.Ctx) error {
	id := c.Params("id")
	feedbackID := c.Params("feedback_id")
	db := database.DB
	var website models.Website
	var feedback models.Feedback 

	if result := db.First(&website, "id = ?", id); result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to get website",
			"error": result.Error,
		})
	}
	
	if res := db.Where("website_id = ?", website.ID).Where("id = ?", feedbackID).Find(&feedback); res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to get website feedback",
			"error": res.Error,
		})
	}

	return c.JSON(fiber.Map{
		"ok": true,
		"website": website,
		"feedback": feedback,
	})
}

// Create a new website 
func CreateWebsite(c *fiber.Ctx) error {
	var website models.Website
	db := database.DB

	if err := c.BodyParser(&website); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to parse form body",
			"error": err,
		})
	} else if vaErr := website.Validate(); vaErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to validate website",
			"error": vaErr,
		})
	} else if result := db.Create(&website); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to create website",
			"error": result.Error,
		})
	} else {
		return c.JSON(fiber.Map{
			"ok": true,
			"website": website,
			"message": "New website has been successfully created",
		})
	}	
}

// Update website
func UpdateWebsite(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var website models.Website

	result := db.First(&website, "id = ?", id)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok": 		false,
			"message": "Error while trying to get website",
			"error":	result.Error,
		})
	}
	
	if err := c.BodyParser(&website); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to parse form body",
			"error": err,
		})
	}

	db.Model(&website).Updates(website)
	return c.JSON(fiber.Map{
		"ok":	true,
		"message": "website has been successfully updated",
		"website": website,
	})
}

// Delete website
func DeleteWebsite(c *fiber.Ctx) error {
	id := c.Params("id")

	var website models.Website
	db := database.DB
	db.First(&website, "id = ?", id)

	if website.ID == "" {
		return c.Status(404).JSON(fiber.Map{
			"ok": 		false,
			"message": "No website found with provided ID",
		})
	}

	db.Delete(&website)

	return c.JSON(fiber.Map{
		"ok":			true,
		"message": "website has been successfully deleted with all its feedback messages",
		"website":		website,
	})	
}