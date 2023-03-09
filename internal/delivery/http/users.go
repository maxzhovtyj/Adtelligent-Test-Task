package delivery

import (
	"encoding/json"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
	"net/http"
	"net/mail"
)

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) signIn(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var input SignInInput

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = mail.ParseAddress(input.Email)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if len(input.Password) < 4 {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := h.services.Users.SignIn(models.User{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(writer).Encode(map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

type SignUpInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) signUp(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var input SignUpInput

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = mail.ParseAddress(input.Email)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if len(input.Password) < 4 {
		http.Error(writer, "invalid password length", http.StatusBadRequest)
		return
	}

	err = h.services.SignUp(models.User{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
