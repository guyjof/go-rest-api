# Golang REST API

This is a simple REST API written in Golang.

## Endpoints

- GET /todos
- POST /todos
- GET /todos/{id}
- PATCH /todos/{id}
- DELETE /todos/{id}

## How to run

```bash
go run main.go
```

## How to use

You can use tools like Postman or curl to interact with the API.

### Get all todos

```bash
curl http://localhost:8080/todos
```

### Create a todo

```bash
curl -X POST http://localhost:8080/todos -d '{"item": "New todo"}'
```
or:

```json
{
  "item": "New todo"
}
```

### Get a todo

```bash
curl http://localhost:8080/todos/1
```

### Update a todo

```bash
curl -X PATCH http://localhost:8080/todos/1 -d '{"item": "Updated todo"}'
```
or:

```json
{
  "item": "Updated todo"
}
```

### Delete a todo

```bash
curl -X DELETE http://localhost:8080/todos/1
```
