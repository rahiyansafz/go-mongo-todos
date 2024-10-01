package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rahiyansafz/go-mongo-todos/pb"
)

type Todo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Completed bool               `bson:"completed"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (t *Todo) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	if len(t.Title) > 100 {
		return errors.New("title must be 100 characters or less")
	}
	return nil
}

func (t *Todo) ToProto() *pb.Todo {
	return &pb.Todo{
		Id:        t.ID.Hex(),
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}

func TodoFromProto(protoTodo *pb.Todo) (*Todo, error) {
	id, err := primitive.ObjectIDFromHex(protoTodo.Id)
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:        id,
		Title:     protoTodo.Title,
		Completed: protoTodo.Completed,
		CreatedAt: protoTodo.CreatedAt.AsTime(),
		UpdatedAt: protoTodo.UpdatedAt.AsTime(),
	}, nil
}
