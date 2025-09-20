<!-- Begin README -->

<div align="center">
    <a href="https://github.com/scottgriv/go-json_task_manager_api" target="_blank">
        <img src="./docs/images/icon.png" width="150" height="150"/>
    </a>
</div>
<br>
<p align="center">
    <a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.21.4-00ADD8?style=for-the-badge&logo=go" alt="Go Badge" /></a>
    <br>
    <a href="https://github.com/scottgriv"><img src="https://img.shields.io/badge/github-follow_me-181717?style=for-the-badge&logo=github&color=181717" alt="GitHub Badge" /></a>
    <a href="mailto:scott.grivner@gmail.com"><img src="https://img.shields.io/badge/gmail-contact_me-EA4335?style=for-the-badge&logo=gmail" alt="Email Badge" /></a>
    <a href="https://www.buymeacoffee.com/scottgriv"><img src="https://img.shields.io/badge/buy_me_a_coffee-support_me-FFDD00?style=for-the-badge&logo=buymeacoffee&color=FFDD00" alt="BuyMeACoffee Badge" /></a>
    <br>
    <a href="https://prgportfolio.com" target="_blank"><img src="https://img.shields.io/badge/PRG-Bronze Project-CD7F32?style=for-the-badge&logo=data:image/svg%2bxml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBzdGFuZGFsb25lPSJubyI/Pgo8IURPQ1RZUEUgc3ZnIFBVQkxJQyAiLS8vVzNDLy9EVEQgU1ZHIDIwMDEwOTA0Ly9FTiIKICJodHRwOi8vd3d3LnczLm9yZy9UUi8yMDAxL1JFQy1TVkctMjAwMTA5MDQvRFREL3N2ZzEwLmR0ZCI+CjxzdmcgdmVyc2lvbj0iMS4wIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciCiB3aWR0aD0iMjYuMDAwMDAwcHQiIGhlaWdodD0iMzQuMDAwMDAwcHQiIHZpZXdCb3g9IjAgMCAyNi4wMDAwMDAgMzQuMDAwMDAwIgogcHJlc2VydmVBc3BlY3RSYXRpbz0ieE1pZFlNaWQgbWVldCI+Cgo8ZyB0cmFuc2Zvcm09InRyYW5zbGF0ZSgwLjAwMDAwMCwzNC4wMDAwMDApIHNjYWxlKDAuMTAwMDAwLC0wLjEwMDAwMCkiCmZpbGw9IiNDRDdGMzIiIHN0cm9rZT0ibm9uZSI+CjxwYXRoIGQ9Ik0xMiAzMjggYy04IC04IC0xMiAtNTEgLTEyIC0xMzUgMCAtMTA5IDIgLTEyNSAxOSAtMTQwIDQyIC0zOCA0OAotNDIgNTkgLTMxIDcgNyAxNyA2IDMxIC0xIDEzIC03IDIxIC04IDIxIC0yIDAgNiAyOCAxMSA2MyAxMyBsNjIgMyAwIDE1MCAwCjE1MCAtMTE1IDMgYy04MSAyIC0xMTkgLTEgLTEyOCAtMTB6IG0xMDIgLTc0IGMtNiAtMzMgLTUgLTM2IDE3IC0zMiAxOCAyIDIzCjggMjEgMjUgLTMgMjQgMTUgNDAgMzAgMjUgMTQgLTE0IC0xNyAtNTkgLTQ4IC02NiAtMjAgLTUgLTIzIC0xMSAtMTggLTMyIDYKLTIxIDMgLTI1IC0xMSAtMjIgLTE2IDIgLTE4IDEzIC0xOCA2NiAxIDc3IDAgNzIgMTggNzIgMTMgMCAxNSAtNyA5IC0zNnoKbTExNiAtMTY5IGMwIC0yMyAtMyAtMjUgLTQ5IC0yNSAtNDAgMCAtNTAgMyAtNTQgMjAgLTMgMTQgLTE0IDIwIC0zMiAyMCAtMTgKMCAtMjkgLTYgLTMyIC0yMCAtNyAtMjUgLTIzIC0yNiAtMjMgLTIgMCAyOSA4IDMyIDEwMiAzMiA4NyAwIDg4IDAgODggLTI1eiIvPgo8L2c+Cjwvc3ZnPgo=" alt="Bronze" /></a>
</p>

---------------

<h1 align="center">Go JSON Task Manager API</h1>

This is a simple Task Manager API written in Go. The API provides basic CRUD (Create, Read, Update, Delete) functionality to manage tasks. Tasks are saved in a tasks.json file for persistence.

---------------

## Table of Contents

- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Testing](#testing)
- [What's Inside?](#whats-inside)
- [Resources](#resources)
- [License](#license)
- [Credits](#credits)

## Getting Started

### Installation

1. Clone the Repository:

```bash
git clone https://github.com/scottgriv/go-json_task_manager_api
```

2. Change to the Directory:

```bash
cd go-json-task-manager-api
```

3. Run the Server:
```bash
go run .
```
4. The server should now be running on http://localhost:8080.

### Usage

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

### Testing

Use a tool like [Postman](https://www.postman.com/) or [curl](https://curl.se/) to send requests to the server and test each endpoint. You can also import the included Postman collection in the [api](api) directory to get started.

## What's Inside?

```bash
go-json_task_manager_api/ # Root directory
├── main.go # Entry point
├── router.go # Routes
├── handlers.go # Request handlers
├── tasks.go # Task model
├── postman/ # Postman collection
│   └── go-json_task_manager_api.postman_collection.json # Import this file into Postman
├── sample/ # Sample data
│   └── tasks_sample_output.json # Sample output from GET /tasks
├── tasks.json # Tasks data file
├── .gitignore # git ignore file
├── LICENSE # License file
└── README.md # This file
```

## Resources

- [Go](https://golang.org/)
- [Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Go Tutorial](https://www.golangtutorial.dev/)

## License

This project is released under the terms of **The Unlicense**, which allows you to use, modify, and distribute the code as you see fit. 
- [The Unlicense](https://choosealicense.com/licenses/unlicense/) removes traditional copyright restrictions, giving you the freedom to use the code in any way you choose.
- For more details, see the [LICENSE](LICENSE) file in this repository.

## Credits

**Author:** [Scott Grivner](https://github.com/scottgriv) <br>
**Email:** [scott.grivner@gmail.com](mailto:scott.grivner@gmail.com) <br>
**Website:** [linktr.ee/scottgriv](https://www.linktr.ee/scottgriv) <br>
**Reference:** [Main Branch](https://github.com/scottgriv/go-json_task_manager_api) <br>

---------------

<div align="center">
    <a href="https://linktr.ee/scottgriv" target="_blank">
        <img src="./docs/images/footer.png" width="100" height="100"/>
    </a>
</div>

<!-- End README -->
