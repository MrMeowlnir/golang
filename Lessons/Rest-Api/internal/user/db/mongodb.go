package db

import (
	"Rest-Api/internal/user"
	"Rest-Api/pkg/logging"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type db struct {
	collection 	*mongo.Collection
	logger 		*logging.Logger
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user, error: %s", err)
	}
	d.logger.Debug("convert InsertID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	} else {
		d.logger.Trace(user)
		return "", fmt.Errorf("failed to convert ObjectID to Hex")
	}
}

func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	d.logger.Debug("find user")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to ObjectID. hex: %s", id)
	}
	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		// TODO 404
		return u, fmt.Errorf("filed find user by ID: %s due to error: %s", id, result.Err())
	}
	if err := result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user from DB by ID: %s due to error: %s", id, err)
	}
	return u, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {

	return &db{
		collection: 	database.Collection(collection),
		logger: 		logger,
	}
}