package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
    "os"
)

// Task struct
func main() {

    err := LoadTasksFromFile("tasks.json")
    if err != nil {
        fmt.Println("Error loading Tasks:", err)
        return
    }

    router := NewRouter()
    fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", router)

}

// Save Tasks to tasks.json file
func SaveTasksToFile(filename string) error {
    fmt.Println("Saving tasks to file...")
    
    data, err := json.Marshal(tasks)
    if err != nil {
        fmt.Println("Error marshalling Tasks:", err)
        return err
    }

    err = ioutil.WriteFile(filename, data, 0644)
    if err != nil {
        fmt.Println("Error writing Tasks to file:", err)
    }
    
    fmt.Println("Tasks saved successfully.")
    return err
}

// Load Tasks from tasks.json file
func LoadTasksFromFile(filename string) error {
	mu.Lock()
	defer mu.Unlock()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) { // If the file does not exist
			return ioutil.WriteFile(filename, []byte("[]"), 0644) // Create an empty JSON array file
		}
		return err
	}

	return json.Unmarshal(data, &tasks)
}
