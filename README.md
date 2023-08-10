# Go JSON Task Manager API

This is a simple Task Manager API written in Go. The API provides basic CRUD (Create, Read, Update, Delete) functionality to manage tasks. Tasks are saved in a tasks.json file for persistence.

## Endpoints
```GET /tasks: Retrieve a list of all Tasks.```

Response:
```json
[
  { "id": 1, "name": "Sample Task" },
  { "id": 2, "name": "Modified Task" }
]

```

```POST /task/create: Create a new Task.```

Body:
```json
{
  "name": "Your Task Name"
}
```

```DELETE /task/delete: Delete a Task by its ID.```

Body:
```json
{
  "id": 1
}
```

```PUT /task/update: Update an existing Task by its ID or create a new one if it doesn't exist.```

Body:
```json
{
  "id": 1,
  "name": "Updated Task Name"
}
```

```PATCH /task/modify: Update specific fields of an existing Task by its ID.```

Body:
```json
{
  "id": 1,
  "name": "Modified Task Name"
}
```

## Running the Project

### Clone the Repository:

```bash
git clone https://github.com/scottgriv/go-json-task-manager-api
```

### Change to the Directory:

```bash
cd go-json-task-manager-api
```

### Run the Server:
```bash
go run .
```

The server should now be running on http://localhost:8080.

## Testing
Use a tool like [Postman](https://www.postman.com/) or [curl](https://curl.se/) to send requests to the server and test each endpoint. You can also import the incliuded Postman collection to get started.

## Credit
**Author:** Scott Grivner <br>
**Email:** scott.grivner@gmail.com <br>
**Website:** [scottgrivner.dev](https://www.scottgriv.dev) <br>
**Reference:** [Main Branch](https://github.com/scottgriv/go-json-task-manager-api)