package delivery

import (
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/sign-in", h.signIn)
	mux.HandleFunc("/auth/sign-up", h.signUp)

	return mux
}
