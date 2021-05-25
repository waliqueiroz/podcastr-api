package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/waliqueiroz/podcastr-api/models"
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

func (controller *EpisodeController) Create(c *fiber.Ctx) error {
	var episode models.Episode

	err := json.Unmarshal(c.Body(), &episode)

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	createdEpisode, err := controller.episodeService.Create(episode)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(createdEpisode)
}
