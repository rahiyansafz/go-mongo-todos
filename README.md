# Go MongoDB Todo API

This project is a RESTful API for a Todo application built with Go and MongoDB.

## Features

- CRUD operations for Todo items
- Search functionality
- Pagination support
- MongoDB integration
- Environment variable configuration

## Prerequisites

- Go 1.22.5 or later
- MongoDB 4.2.16 or later
- Docker and Docker Compose (for running MongoDB)

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

4. Build and run the application:
   ```
   make restart
   ```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | /todos   | Get all todos |
| GET    | /todos/{id} | Get a specific todo |
| POST   | /todos   | Create a new todo |
| PUT    | /todos/{id} | Update a todo |
| DELETE | /todos/{id} | Delete a todo |
| GET    | /search?q={query} | Search todos |

For detailed API documentation, refer to the `api.http` file in the project root.

## Usage

You can use curl, Postman, or any API client to interact with the endpoints. Here are some examples using curl:

1. Create a new todo:
   ```
   curl -X POST http://localhost:8080/todos \
        -H "Content-Type: application/json" \
        -d '{"title": "Learn Go", "completed": false}'
   ```

2. Get all todos:
   ```
   curl http://localhost:8080/todos
   ```

3. Get a specific todo (replace {id} with an actual todo ID):
   ```
   curl http://localhost:8080/todos/{id}
   ```

4. Update a todo:
   ```
   curl -X PUT http://localhost:8080/todos/{id} \
        -H "Content-Type: application/json" \
        -d '{"title": "Learn Go Advanced", "completed": true}'
   ```

5. Delete a todo:
   ```
   curl -X DELETE http://localhost:8080/todos/{id}
   ```

6. Search todos:
   ```
   curl http://localhost:8080/search?q=Go
   ```

## Project Structure

- `cmd/api/main.go`: Main application entry point.
- `db/db.go`: Database connection setup.
- `models/todo.go`: Todo model and validation.
- `services/todo.go`: Service functions for CRUD operations.
- `handlers/handlers.go`: HTTP handlers for API endpoints.
- `api.http`: Postman collection for API testing.
- `Makefile`: Convenient commands for Docker and application management.


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.