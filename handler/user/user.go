package user

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"ThreeLayeredArchitecture/models/user"
)

type UserHandler struct {
	Service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newUser, err := h.Service.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *UserHandler) GetUserbyID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
