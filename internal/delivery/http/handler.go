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

	mux.HandleFunc("/seller/add", h.userIdentity(h.newSeller))

	mux.HandleFunc("/product/add", h.userIdentity(h.newProduct))
	mux.HandleFunc("/product/get", h.userIdentity(h.getProduct))
	mux.HandleFunc("/product/update", h.userIdentity(h.updateProduct))
	mux.HandleFunc("/product/delete", h.userIdentity(h.deleteProduct))

	return mux
}
