package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/waliqueiroz/podcastr-api/services"
)

type EpisodeController struct {
	episodeService *services.EpisodeService
}

func NewEpisodeController(episodeService *services.EpisodeService) *EpisodeController {
	return &EpisodeController{
		episodeService: episodeService,
	}
}

func (controller *EpisodeController) Index(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
