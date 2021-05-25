package services

import (
	"github.com/waliqueiroz/podcastr-api/models"
	"github.com/waliqueiroz/podcastr-api/repositories"
)

type EpisodeService struct {
	episodeRepository *repositories.EpisodeRepository
}

func NewEpisodeService(episodeRepository *repositories.EpisodeRepository) *EpisodeService {
	return &EpisodeService{
		episodeRepository: episodeRepository,
	}
}

func (service *EpisodeService) Create(episode models.Episode) (models.Episode, error) {
	return service.episodeRepository.Create(episode)
}
