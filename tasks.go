package main

import (
	"sync"
    "fmt"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks = make([]Task, 0)
var currentID = 1
var mu sync.Mutex

func GetAllTasks() []Task {
	mu.Lock()
	defer mu.Unlock()
	return tasks
}

func AddTask(task Task) {
    
    mu.Lock()
    defer mu.Unlock()

    task.ID = currentID
    currentID++
    tasks = append(tasks, task)
    
    fmt.Println("Tasks before saving:", tasks) 
    err := SaveTasksToFile("tasks.json")
    if err != nil {
        fmt.Println("Error saving Tasks to file:", err)
    }
}

func DeleteTask(taskID int) bool {
	mu.Lock()
	defer mu.Unlock()

	for index, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:index], tasks[index+1:]...)
            SaveTasksToFile("tasks.json")
			return true
		}
	}
	return false
}

func GetTaskByID(id int) (Task, bool) {
	mu.Lock()
	defer mu.Unlock()

	for _, t := range tasks {
		if t.ID == id {
			return t, true
		}
	}
	return Task{}, false
}

func UpdateTask(updatedTask Task) {
	mu.Lock()
	defer mu.Unlock()

	for index, t := range tasks {
		if t.ID == updatedTask.ID {
			tasks[index] = updatedTask
            SaveTasksToFile("tasks.json")
			return
		}
	}
}

func ModifyTask(taskID int, name string) {
	mu.Lock()
	defer mu.Unlock()

	for index, t := range tasks {
		if t.ID == taskID && name != "" {
			tasks[index].Name = name
            SaveTasksToFile("tasks.json")
			return
		}
	}
}
