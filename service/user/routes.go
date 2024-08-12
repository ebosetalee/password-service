package user

import (
	"net/http"

	"github.com/ebosetalee/password-service.git/types"
	"github.com/ebosetalee/password-service.git/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.login).Methods("POST")
	router.HandleFunc("/register", h.register).Methods("POST")
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
}
func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	// get json payload
	var payload types.RegisterPayload
	if err := utils.ReadJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// check if user exists

	
	// create user or return error
}
