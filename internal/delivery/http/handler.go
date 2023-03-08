package delivery

import (
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/service"
	"net/http"
)

type handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *handler {
	return &handler{
		services: services,
	}
}

func (h *handler) Init() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/sign-in", h.signIn)
	mux.HandleFunc("/auth/sign-up", h.signUp)

	return mux
}
