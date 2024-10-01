package services

import (
	"context"
	"errors"
	"time"

	"github.com/rahiyansafz/go-mongo-todos/db"
	"github.com/rahiyansafz/go-mongo-todos/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTodo(todo models.Todo) (*models.Todo, error) {
	if err := todo.Validate(); err != nil {
		return nil, err
	}

	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	result, err := db.GetCollection().InsertOne(context.Background(), todo)
	if err != nil {
		return nil, err
	}
	todo.ID = result.InsertedID.(primitive.ObjectID)
	return &todo, nil
}

func GetAllTodos(limit, page int64) ([]models.Todo, error) {
	var todos []models.Todo

	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip((page - 1) * limit)
	options.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := db.GetCollection().Find(context.Background(), bson.M{}, options)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func GetTodoByID(id string) (*models.Todo, error) {
	var todo models.Todo
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = db.GetCollection().FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&todo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}
	return &todo, nil
}

func UpdateTodo(id string, updatedTodo models.Todo) (*models.Todo, error) {
	if err := updatedTodo.Validate(); err != nil {
		return nil, err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	updatedTodo.UpdatedAt = time.Now()
	_, err = db.GetCollection().UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		bson.M{"$set": updatedTodo},
	)
	if err != nil {
		return nil, err
	}
	return GetTodoByID(id)
}

func DeleteTodo(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := db.GetCollection().DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("todo not found")
	}
	return nil
}

func SearchTodos(query string) ([]models.Todo, error) {
	filter := bson.M{
		"title": bson.M{
			"$regex":   query,
			"$options": "i",
		},
	}
	cursor, err := db.GetCollection().Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var todos []models.Todo
	if err = cursor.All(context.Background(), &todos); err != nil {
		return nil, err
	}
	return todos, nil
}
