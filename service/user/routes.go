package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ebosetalee/password-service.git/service/auth"
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
	log.Println("user logging in...")
	var payload types.LoginPayload
	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusPreconditionFailed, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusPreconditionFailed, fmt.Errorf(utils.ValidateError(err)))
		return
	}

	// get user account or throw error
	user, err := h.repo.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("invalid credentials"))
		return
	}

	// verify hashed password
	if !auth.VerifyHash(payload.Password, user.Password) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials, try again"))
		return
	}

	if user.DeletedAt != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("you have previously deleted your account, please contact support"))
	}

	// generate jwt token
	token, err := auth.GenerateJWT(user.ID, user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	user.Password = ""
	response := types.Response{
		Message: "user log in successful",
		Data: types.LoginResponse{
			Token: token,
			User:  user,
		},
	}

	utils.WriteJSON(w, http.StatusCreated, response)
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating a new user...")
	// get json payload
	var payload types.RegisterPayload
	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusPreconditionFailed, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusPreconditionFailed, fmt.Errorf(utils.ValidateError(err)))
		return
	}

	// check if user exists
	u, err := h.repo.GetUserByEmail(payload.Email)
	if u.Email == payload.Email {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	hashedPassword, err := auth.Hash(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// create user or return error
	err = h.repo.CreateUser(types.User{
		Username:  &payload.Username,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
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
