# Go JSON Task Manager API

This is a simple Task Manager API written in Go. The API provides basic CRUD (Create, Read, Update, Delete) functionality to manage tasks. Tasks are saved in a tasks.json file for persistence.

## API Endpoints
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

## Getting Started

### Clone the Repository:

```bash
git clone https://github.com/scottgriv/go-json_task_manager_api
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

## What's Inside?
go-json_task_manager_api/ 
├── main.go (entry point)
├── router.go (routes)
├── handlers.go (request handlers)
├── tasks.go (task model)
├── postman/ (Postman collection)
│   └── go-json_task_manager_api.postman_collection.json (import this file into Postman)
├── sample/ (sample data)
│   └── tasks_sample_output.json (sample output from GET /tasks)
├── .gitignore (git ignore file)
├── LICENSE (license file)
└── README.md (this file)

## Credit
**Author:** Scott Grivner <br>
**Email:** scott.grivner@gmail.com <br>
**Website:** [scottgrivner.dev](https://www.scottgriv.dev) <br>
**Reference:** [Main Branch](https://github.com/scottgriv/go-json_task_manager_api)