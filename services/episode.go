package services

import "github.com/waliqueiroz/podcastr-api/repositories"

type EpisodeService struct {
	episodeRepository *repositories.EpisodeRepository
}

func NewEpisodeService(episodeRepository *repositories.EpisodeRepository) *EpisodeService {
	return &EpisodeService{
		episodeRepository: episodeRepository,
	}
}