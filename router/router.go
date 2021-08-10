package router

import (
	"github.com/lne-io/feedback/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Private endpoints
	api := app.Group("/api")
	v1 := api.Group("/v1")

	website := v1.Group("/websites")
	website.Get("/", controllers.GetWebsites)
	website.Get("/page/:page", controllers.GetWebsites)
	website.Get("/:id", controllers.GetWebsiteById)
	website.Get("/:id/feedback", controllers.GetWebsiteFeedback)
	website.Get("/:id/feedback/page/:page", controllers.GetWebsiteFeedback)
	website.Get("/:id/feedback/:id", controllers.GetWebsiteSingleFeedback)
	website.Post("/", controllers.CreateWebsite)
	website.Patch("/:id", controllers.UpdateWebsite)
	website.Delete("/:id", controllers.DeleteWebsite)
	
	// Public endpoints
	public := app.Group("/public")
	apiPublic := public.Group("/api")
	v1Public := apiPublic.Group("/v1")

	feedback := v1Public.Group("/feedback")
	feedback.Post("/", controllers.CreateFeedback)

}