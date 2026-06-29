package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RAHULBONEY/clITaskTracker/internal/task"
)

type TaskRequest struct {
	Name string `json:"name"`
}

func main() {
	err := task.LoadTasks()
	if err != nil {
		fmt.Println("Could not load tasks:", err)
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/tasks", getTasks)
	http.HandleFunc("POST /tasks", createTask)
	http.HandleFunc("UPDATE /tasks", updateTask)
	http.HandleFunc("DELETE /tasks", deleteTask)
	fmt.Println("Server running fro port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error in starting the server", err)
	}

}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello! Go api successfully")
}
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(task.Tasks)
	if err != nil {
		http.Error(w, "Error in fetching tasks", http.StatusInternalServerError)
	}

}
func createTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req TaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = task.AddTask(req.Name)
	if err != nil {
		http.Error(w, "Error in creating task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Task created successfully")
}
func updateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var taskID int
	err := json.NewDecoder(r.Body).Decode(&taskID)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = task.CompleteTask(taskID)
	if err != nil {
		http.Error(w, "Error in updating task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Task updated successfully")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var taskId int
	err := json.NewDecoder(r.Body).Decode(&taskId)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = task.DeleteTask(taskId)
	if err != nil {
		http.Error(w, "Error in deleting task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Task deleted successfully")

}
