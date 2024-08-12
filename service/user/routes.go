package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ebosetalee/password-service.git/types"
	"github.com/ebosetalee/password-service.git/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	repo types.UserRepo
}

func NewHandler(repo types.UserRepo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.login).Methods("POST")
	router.HandleFunc("/register", h.register).Methods("POST")
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating a new user...")
	// get json payload
	var payload types.RegisterPayload
	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// check if user exists
	_, err := h.repo.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// create user or return error
	err = h.repo.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response := types.Response{
		Message: "user created successfully",
	}

	utils.WriteJSON(w, http.StatusCreated, response)
}
