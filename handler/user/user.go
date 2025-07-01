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

// CreateUser godoc
// @Summary      Create a new user
// @Description  Add a new user to the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.User true "User object"
// @Success      201 {object} models.User
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /users [post]
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

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieve all users from the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200 {array} models.User
// @Failure      500 {string} string "Internal Server Error"
// @Router       /users [get]
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

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Retrieve a specific user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} models.User
// @Failure      400 {string} string "Invalid ID"
// @Failure      404 {string} string "User not found"
// @Router       /users/{id} [get]
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
