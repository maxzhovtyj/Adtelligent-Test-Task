package delivery

import (
	"net/http"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Init() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/sign-in", h.signIn)
	mux.HandleFunc("/auth/sign-up", h.signUp)

	return mux
}
