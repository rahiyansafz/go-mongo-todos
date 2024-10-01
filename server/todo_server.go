package server

import (
	"context"
	"errors"
	"time"

	"github.com/rahiyansafz/go-mongo-todos/db"
	"github.com/rahiyansafz/go-mongo-todos/models"
	"github.com/rahiyansafz/go-mongo-todos/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.Todo, error) {
	todo := &models.Todo{
		Title:     req.Title,
		Completed: req.Completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := todo.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := db.GetCollection().InsertOne(ctx, todo)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create todo")
	}

	todo.ID = result.InsertedID.(primitive.ObjectID)
	return todo.ToProto(), nil
}

func (s *TodoServer) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.Todo, error) {
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ID")
	}

	var todo models.Todo
	err = db.GetCollection().FindOne(ctx, bson.M{"_id": objectID}).Decode(&todo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, status.Error(codes.NotFound, "Todo not found")
		}
		return nil, status.Error(codes.Internal, "Failed to get todo")
	}

	return todo.ToProto(), nil
}

func (s *TodoServer) ListTodos(ctx context.Context, req *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	limit := int64(req.Limit)
	if limit == 0 {
		limit = 10
	}
	skip := int64(req.Page-1) * limit
	if skip < 0 {
		skip = 0
	}

	opts := options.Find().SetLimit(limit).SetSkip(skip).SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := db.GetCollection().Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to list todos")
	}

	var todos []*models.Todo
	if err = cursor.All(ctx, &todos); err != nil {
		return nil, status.Error(codes.Internal, "Failed to decode todos")
	}

	protoTodos := make([]*pb.Todo, len(todos))
	for i, todo := range todos {
		protoTodos[i] = todo.ToProto()
	}

	return &pb.ListTodosResponse{Todos: protoTodos}, nil
}

func (s *TodoServer) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.Todo, error) {
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ID")
	}

	update := bson.M{
		"$set": bson.M{
			"title":      req.Title,
			"completed":  req.Completed,
			"updated_at": time.Now(),
		},
	}

	result, err := db.GetCollection().UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update todo")
	}

	if result.MatchedCount == 0 {
		return nil, status.Error(codes.NotFound, "Todo not found")
	}

	var updatedTodo models.Todo
	err = db.GetCollection().FindOne(ctx, bson.M{"_id": objectID}).Decode(&updatedTodo)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get updated todo")
	}

	return updatedTodo.ToProto(), nil
}

func (s *TodoServer) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*emptypb.Empty, error) {
	objectID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ID")
	}

	result, err := db.GetCollection().DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to delete todo")
	}

	if result.DeletedCount == 0 {
		return nil, status.Error(codes.NotFound, "Todo not found")
	}

	return &emptypb.Empty{}, nil
}

func (s *TodoServer) SearchTodos(ctx context.Context, req *pb.SearchTodosRequest) (*pb.ListTodosResponse, error) {
	filter := bson.M{
		"title": bson.M{
			"$regex":   req.Query,
			"$options": "i",
		},
	}

	cursor, err := db.GetCollection().Find(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to search todos")
	}

	var todos []*models.Todo
	if err = cursor.All(ctx, &todos); err != nil {
		return nil, status.Error(codes.Internal, "Failed to decode todos")
	}

	protoTodos := make([]*pb.Todo, len(todos))
	for i, todo := range todos {
		protoTodos[i] = todo.ToProto()
	}

	return &pb.ListTodosResponse{Todos: protoTodos}, nil
}
