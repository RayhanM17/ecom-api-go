package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// Routes defines and returns the user routes as a chi.Router
func (h *Handler) UserRoutes(r *chi.Mux) {
	r.Get("/login", h.Login)
	r.Get("/logout", h.Register)
}

// Login handler
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

}

// Logout handler
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

}
