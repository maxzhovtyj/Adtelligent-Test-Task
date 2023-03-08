package delivery

import (
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/service"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/pkg/auth"
	"net/http"
)

type Handler struct {
	services     *service.Service
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Service, manager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: manager,
	}
}

func (h *Handler) Init() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/sign-in", h.signIn)
	mux.HandleFunc("/auth/sign-up", h.signUp)

	mux.HandleFunc("/seller", h.userIdentity(h.newSeller))

	mux.HandleFunc("/product", h.userIdentity(h.newProduct))

	return mux
}
