package main

import (
	"flag"

	"github.com/lne-io/feedback/config"
	"github.com/lne-io/feedback/router"
	"github.com/lne-io/feedback/database"
	"github.com/lne-io/feedback/taskQueue"
	mdl "github.com/lne-io/feedback/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	//"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)



func main() {

	// Load configuration
	config.LoadEnv()
	// Database
	database.ConnectDB()

	isWorker := flag.Bool("worker", false, "If true, run worker instead of web server")
	flag.Parse()
	if (*isWorker) {
		// Task queue server
		taskQueue.ConnectTaskQueueWorker()
	} else {
		// Task queue client
		taskQueue.ConnectTaskQueueClient()
		

		app := fiber.New()
		app.Use(recover.New())
		app.Use(logger.New())
		//app.Use(pprof.New())
		app.Use(mdl.UnslashURL())

		router.SetupRoutes(app)
		

		app.Use(func(c *fiber.Ctx) error {
			return c.SendStatus(404)
		})


		app.Listen(":3000")
	}
	
}