package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waliqueiroz/podcastr-api/controllers"
)

func SetupEpisodeRoutes(router fiber.Router, episodeController *controllers.EpisodeController) {
	router.Get("/episodes", episodeController.Index)
	router.Post("/episodes", episodeController.Create)
}
