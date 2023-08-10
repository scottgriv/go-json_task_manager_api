package main

import (
    "net/http"
)

func NewRouter() *http.ServeMux {
    router := http.NewServeMux()

    router.HandleFunc("/tasks", GetTasksHandler)
    router.HandleFunc("/task/create", CreateTaskHandler)
    router.HandleFunc("/task/delete", DeleteTaskHandler)
    router.HandleFunc("/task/update", PutTaskHandler)
    router.HandleFunc("/task/modify", PatchTaskHandler)
    
    return router
}
