# Go MongoDB gRPC Todo Service

This project is a gRPC-based Todo service built with Go and MongoDB.

## Features

- CRUD operations for Todo items
- Search functionality
- Pagination support
- MongoDB integration
- gRPC API

## Prerequisites

- Go 1.22.5 or later
- MongoDB 4.2.16 or later
- Docker and Docker Compose (for running MongoDB)
- Protocol Buffers compiler (protoc)

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/go-mongo-todos.git
   cd go-mongo-todos
   ```

2. Copy the example environment file and modify as needed:
   ```
   cp .env.example .env
   ```

3. Start the MongoDB container:
   ```
   make up
   ```

4. Generate gRPC code:
   ```
   make proto
   ```

5. Build and run the application:
   ```
   make build
   make run
   ```

## gRPC API

The service exposes the following gRPC methods:

- `CreateTodo`: Creates a new todo item
- `GetTodo`: Retrieves a specific todo item by ID
- `ListTodos`: Lists all todo items with pagination
- `UpdateTodo`: Updates an existing todo item
- `DeleteTodo`: Deletes a todo item
- `SearchTodos`: Searches for todo items based on a query string

## Usage

To interact with the gRPC service, you'll need a gRPC client. You can use tools like [grpcurl](https://github.com/fullstorydev/grpcurl) or create a custom client using the generated protobuf code.

Here are some example commands using grpcurl:

1. Create a new todo:
   ```
   grpcurl -plaintext -d '{"title": "Learn gRPC", "completed": false}' localhost:50051 todo.TodoService/CreateTodo
   ```

2. Get a todo by ID:
   ```
   grpcurl -plaintext -d '{"id": "5f9f1b9b9c9d440000a1b1b1"}' localhost:50051 todo.TodoService/GetTodo
   ```

3. List todos:
   ```
   grpcurl -plaintext -d '{"page": 1, "limit": 10}' localhost:50051 todo.TodoService/ListTodos
   ```

4. Update a todo:
   ```
   grpcurl -plaintext -d '{"id": "5f9f1b9b9c9d440000a1b1b1", "title": "Learn gRPC Advanced", "completed": true}' localhost:50051 todo.TodoService/UpdateTodo
   ```

5. Delete a todo:
   ```
   grpcurl -plaintext -d '{"id": "5f9f1b9b9c9d440000a1b1b1"}' localhost:50051 todo.TodoService/DeleteTodo
   ```

6. Search todos:
   ```
   grpcurl -plaintext -d '{"query": "gRPC"}' localhost:50051 todo.TodoService/SearchTodos
   ```

## Project Structure

```
.
├── cmd
│   └── server
│       └── main.go
├── pb
│   └── todo.proto
├── server
│   └── todo_server.go
├── db
│   └── db.go
├── models
│   └── todo.go
├── .env
├── .gitignore
├── docker-compose.yml
├── go.mod
├── Makefile
└── README.md
```

- `cmd/server/main.go`: Main application entry point
- `pb/todo.proto`: Protocol Buffers definition file
- `server/todo_server.go`: gRPC server implementation
- `db/db.go`: Database connection setup
- `models/todo.go`: Todo model and validation

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.