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

func (repository *EpisodeRepository) FindAll() ([]models.Episode, error) {
	result, err := repository.episodeCollection.Find(repository.ctx, bson.M{})

	if err != nil {
		return []models.Episode{}, err
	}

	defer result.Close(repository.ctx)

	var episodes []models.Episode

	for result.Next(repository.ctx) {
		var episode models.Episode

		err := result.Decode(&episode)
		if err != nil {
			return []models.Episode{}, err
		}

		episodes = append(episodes, episode)
	}

	return episodes, nil
}

func (repository *EpisodeRepository) Create(episode models.Episode) (models.Episode, error) {
	result, err := repository.episodeCollection.InsertOne(repository.ctx, episode)
	if err != nil {
		return models.Episode{}, err
	}

	objectID := result.InsertedID.(primitive.ObjectID)

	return repository.FindById(objectID.Hex())
}

func (repository *EpisodeRepository) FindById(episodeID string) (models.Episode, error) {
	id, _ := primitive.ObjectIDFromHex(episodeID)
	filter := bson.M{"_id": id}

	result := repository.episodeCollection.FindOne(repository.ctx, filter)

	var episode models.Episode

	err := result.Decode(&episode)
	if err != nil {
		return models.Episode{}, err
	}

	return episode, nil
}
func (repository *EpisodeRepository) FindBySlug(episodeSlug string) (models.Episode, error) {
	filter := bson.M{"slug": episodeSlug}

	result := repository.episodeCollection.FindOne(repository.ctx, filter)

	var episode models.Episode

	err := result.Decode(&episode)
	if err != nil {
		return models.Episode{}, err
	}

	return episode, nil
}

func (repository *EpisodeRepository) Delete(episodeID string) error {
	id, _ := primitive.ObjectIDFromHex(episodeID)
	filter := bson.M{"_id": id}

	_, err := repository.episodeCollection.DeleteOne(repository.ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
