package db

import (
	"Rest-Api/internal/user"
	"Rest-Api/pkg/logging"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		if errors.Is(result.Err(), mongo.ErrNoDocuments){
			// TODO Error Entity not found
			return u, fmt.Errorf("error: Entity Not Found")
		}
		return u, fmt.Errorf("filed find user by ID: %s due to error: %s", id, result.Err())
	}
	if err := result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user from DB by ID: %s due to error: %s", id, err)
	}
	return u, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("filed convertion user.ID to objectID: %s", user.ID)
	}
	filter := bson.M{"_id": objectID}

	var userBytes []byte
	userBytes, err = bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("filed to marshal user due to error: %v", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("filed to unmarshal userBytes due to error: %v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{
		"$set": updateUserObj,
	}

	var result *mongo.UpdateResult
	result, err = d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("filed to Update user due to error: %v", err)
	}
	if result.MatchedCount == 0 {
		// TODO errorentity not found 404
		return fmt.Errorf("not found")
	}

	d.logger.Tracef("Matched %d docs; Modified: %d docs;", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed convertion id to objectID: %s", id)
	}
	filter := bson.M{"_id": objectID}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query due to error: %v", err)
	}

	if result.DeletedCount == 0 {
		// TODO errorentity not found 404
		return fmt.Errorf("not found")
	}

	d.logger.Tracef("Deleted: %d docs;", result.DeletedCount)

	return nil
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: 	database.Collection(collection),
		logger: 		logger,
	}
}