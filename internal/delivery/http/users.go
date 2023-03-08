package delivery

import (
	"encoding/json"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
	"net/http"
)

type SignInInput struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func (h *handler) signIn(writer http.ResponseWriter, request *http.Request) {
	var input SignInInput

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.services.SignInSeller(models.Seller{
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	return
}

type SignUpInput struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func (h *handler) signUp(writer http.ResponseWriter, request *http.Request) {
	var input SignUpInput

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.services.SignUpSeller(models.Seller{
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	return
}
