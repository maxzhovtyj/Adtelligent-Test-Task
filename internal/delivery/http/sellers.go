package delivery

import (
	"encoding/json"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
	"net/http"
)

type NewSellerInput struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

func (h *Handler) newSeller(writer http.ResponseWriter, request *http.Request) {
	var input NewSellerInput

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(writer, models.ErrInvalidInputBody.Error(), http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		http.Error(writer, "invalid seller name", http.StatusBadRequest)
		return
	}

	if input.PhoneNumber == "" {
		http.Error(writer, "invalid seller phone number", http.StatusBadRequest)
		return
	}

	err = h.services.Sellers.Create(models.Seller{
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = writer.Write([]byte("Seller successfully created"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
