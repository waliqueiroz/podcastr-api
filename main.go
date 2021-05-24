package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/waliqueiroz/podcastr-api/config"
	"github.com/waliqueiroz/podcastr-api/controllers"
	"github.com/waliqueiroz/podcastr-api/database"
	"github.com/waliqueiroz/podcastr-api/repositories"
	"github.com/waliqueiroz/podcastr-api/services"
)

func main() {
	config.Load()
	ctx := context.Background()

	db, err := database.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
		return
	}

	episodeRepository := repositories.NewEpisodeRepository(ctx, db)
	espisodeService := services.NewEpisodeService(episodeRepository)
	episodeController := controllers.NewEpisodeController(espisodeService)

	app := fiber.New()
	app.Get("/", episodeController.Index)

	app.Listen(":8080")
}
