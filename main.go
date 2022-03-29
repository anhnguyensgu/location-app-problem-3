package main

import (
	"log"
	"os"
	"problem3/web-service/handler"
	"problem3/web-service/mgconfig"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	port := os.Getenv("PORT")
	engine := html.New("./view", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./view")

	client, ctx, eventRepository, visitStatsRepository := mgconfig.InitializeMongoConnection()
	eventHadler := handler.NewEventHandler(eventRepository, visitStatsRepository)
	statHandler := handler.NewStatHandler(visitStatsRepository)
	defer client.Disconnect(ctx)

	app.Get("/api/events", eventHadler.GetEvents)

	app.Post("/api/events", eventHadler.CreateEvent)

	app.Get("/api/stats", statHandler.GetStats)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	log.Fatal(app.Listen(":" + port))
}
