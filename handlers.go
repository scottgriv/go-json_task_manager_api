package main

import (
    "encoding/json"
    "net/http"
    "fmt"
)

// GET
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    tasks := GetAllTasks()

    // Check if Tasks slice is empty
    if len(tasks) == 0 {
        w.Header().Set("Content-Type", "application/json")
        http.Error(w, "No Tasks available, please create a new Task!", http.StatusNotFound)
        // json.NewEncoder(w).Encode(map[string]string{"message": "No Tasks available, please create a new Task!"})
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

// POST
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var task Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Invalid task data", http.StatusBadRequest)
        return
    }
    defer r.Body.Close() // Close the request body

    AddTask(task)

    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    _, err := w.Write([]byte("Task created successfully"))
    if err != nil {
        fmt.Println("Error writing response:", err)
    }
}

// DELETE
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var task Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Invalid Task data", http.StatusBadRequest)
        return
    }

    if DeleteTask(task.ID) {
        w.Write([]byte("Task deleted successfully"))
    } else {
        http.Error(w, "Task ID does not exist", http.StatusNotFound)
    }
}

// PUT
func PutTaskHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var task Task
    err := json.NewDecoder(r.Body).Decode(&task)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    if _, exists := GetTaskByID(task.ID); exists {
        // Replace the task
        UpdateTask(task)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Task updated successfully"))
    } else {
        // Add the new task
        AddTask(task)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Task does not exist, new Task added"))
    }

}

// PATCH
func PatchTaskHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPatch {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var updatedTask Task
    err := json.NewDecoder(r.Body).Decode(&updatedTask)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }

    task, exists := GetTaskByID(updatedTask.ID)
    if !exists {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    // Modify only provided fields
    if updatedTask.Name != "" {
        task.Name = updatedTask.Name
    }
    
    UpdateTask(task)
    
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Task modified successfully"))
}


