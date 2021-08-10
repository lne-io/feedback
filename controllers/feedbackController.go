package controllers

import (
	"log"

	"github.com/lne-io/feedback/models"
	"github.com/lne-io/feedback/taskQueue"
	"github.com/lne-io/feedback/tasks"

	"github.com/gofiber/fiber/v2"
)

// Recieves user submitted content. Parse it and validate it then enqueue it to our task queue
func CreateFeedback(c *fiber.Ctx) error {
	var feedback models.Feedback 

	err := c.BodyParser(&feedback); 
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to parse form body",
			"error": err,
		})
	}
	
	vaErr := feedback.Validate()
	if vaErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Error while trying to validate feedback",
			"error": vaErr,
		})
	} 
	
	task, err := tasks.NewRegisterFeedbackTask(&feedback)
    if err != nil {
		log.Printf("could not create task: %v\n", err)
        return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Internal Error while trying to save feedback",
		})
    }
	
	queueClient := taskQueue.QueueClient

	info, err := queueClient.Enqueue(task)
	log.Printf("I am after enqueue , %v \n", queueClient)

    if err != nil {
		log.Printf("could not enqueue task: %v\n", err)
        return c.Status(500).JSON(fiber.Map{
			"ok": false,
			"message": "Internal Error while trying to save feedback",
		})
    }
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	
	return c.JSON(fiber.Map{
		"ok": true,
		"message": "Thank you for your feedback",
	})
}
