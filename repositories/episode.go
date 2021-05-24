package repositories

import (
	"context"

	"github.com/waliqueiroz/podcastr-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type EpisodeRepository struct {
	episodeCollection *mongo.Collection
	ctx               context.Context
}

func NewEpisodeRepository(ctx context.Context, db *mongo.Database) *EpisodeRepository {

	return &EpisodeRepository{
		episodeCollection: db.Collection("episodes"),
		ctx:               ctx,
	}
}

func (repository *EpisodeRepository) Create(episode models.Episode) error {
	_, err := repository.episodeCollection.InsertOne(repository.ctx, episode)
	return err
}
