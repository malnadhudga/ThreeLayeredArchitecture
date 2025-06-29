package taskhandler

import (
	"ThreeLayeredArchitecture/models/task"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

//type Taskservice interface {
//	GetPendingTasks() ([]models.Task, error)
//	AddTask(desc string) (models.Task, error)
//	DeleteTask(id int) error
//	CompleteTask(id int) (string, error)
//	GetTaskByID(id int) (models.Task, error)
//}

type TaskHandler struct {
	Service Taskservice
}

func NewTaskHandler(service Taskservice) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		task, err := h.Service.GetPendingTasks()
		if err != nil {
			http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(task)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		_, err = w.Write(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case "POST":
		var input models.Task
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read body", http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(body, &input)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		task, err := h.Service.AddTask(input.Description)
		if err != nil {
			http.Error(w, "Failed to add", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		data, err := json.Marshal(task)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		_, err = w.Write(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case "DELETE":
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		err := h.Service.DeleteTask(id)
		if err != nil {
			http.Error(w, "Delete failed", http.StatusNotFound)
			return
		}
		_, err = w.Write([]byte("Deleted successfully"))
		if err != nil {
			http.Error(w, "Delete failed", http.StatusInternalServerError)
			return
		}

	case "PATCH":
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		msg, err := h.Service.CompleteTask(id)
		if err != nil {
			http.Error(w, "Update failed", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write([]byte(msg))
		if err != nil {
			http.Error(w, "update failed", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) HandleTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		task, err := h.Service.GetTaskByID(id)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		responseData, err := json.Marshal(task)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(responseData)
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
