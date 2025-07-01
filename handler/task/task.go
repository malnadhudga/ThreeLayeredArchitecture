package task

import (
	"ThreeLayeredArchitecture/models/task"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	Service Taskservice
}

func NewTaskHandler(service Taskservice) *TaskHandler {
	return &TaskHandler{Service: service}
}

// HandleTasks godoc
// @Summary Handle task operations
// @Description Supports GET, POST, DELETE, and PATCH on /task
// @Tags tasks
// @Accept json
// @Produce json
// @Router /task [get]
// @Router /task [post]
// @Router /task [delete]
// @Router /task [patch]

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		task, err := h.Service.GetPending()
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
	case http.MethodPost:
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

		task, err := h.Service.Add(input.Description)
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

	case http.MethodDelete:
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		err := h.Service.Delete(id)
		if err != nil {
			http.Error(w, "Delete failed", http.StatusNotFound)
			return
		}

		_, err = w.Write([]byte("Deleted successfully"))
		if err != nil {
			http.Error(w, "Delete failed", http.StatusInternalServerError)
			return
		}

	case http.MethodPatch:
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		msg, err := h.Service.MarkComplete(id)
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

// HandleTaskByID godoc
// @Summary      Get task by ID
// @Description  Retrieve task details by its ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id path int true "Task ID"
// @Success      200 {object} models.Task
// @Failure      400 {string} string "Invalid ID"
// @Failure      404 {string} string "Task not found"
// @Router       /tasks/{id} [get]
func (h *TaskHandler) HandleTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		task, err := h.Service.GetByID(id)
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
