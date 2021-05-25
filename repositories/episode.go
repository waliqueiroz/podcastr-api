package repositories

import (
	"context"

	"github.com/waliqueiroz/podcastr-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (repository *EpisodeRepository) Create(episode models.Episode) (models.Episode, error) {
	result, err := repository.episodeCollection.InsertOne(repository.ctx, episode)
	if err != nil {
		return models.Episode{}, err
	}

	objectID := result.InsertedID.(primitive.ObjectID)

	return repository.FindById(objectID.Hex())
}

func (repository *EpisodeRepository) FindById(id string) (models.Episode, error) {
	episodeID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": episodeID}

	result := repository.episodeCollection.FindOne(repository.ctx, filter)

	var episode models.Episode

	err := result.Decode(&episode)
	if err != nil {
		return models.Episode{}, err
	}

	return episode, nil
}
