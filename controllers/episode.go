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
	episodes, err := controller.episodeService.FindAll()

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(episodes)
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

func (controller *EpisodeController) FindBySlug(c *fiber.Ctx) error {
	episodeSlug := c.Params("episodeSlug")

	episode, err := controller.episodeService.FindBySlug(episodeSlug)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(episode)
}

func (controller *EpisodeController) Delete(c *fiber.Ctx) error {
	episodeID := c.Params("episodeID")

	err := controller.episodeService.Delete(episodeID)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
