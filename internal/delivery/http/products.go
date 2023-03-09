package delivery

import (
	"encoding/json"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) getProduct(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	productID := request.URL.Query().Get("id")

	productIDint, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(writer, "invalid product id", http.StatusBadRequest)
		return
	}

	product, err := h.services.Products.Get(productIDint)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(&product)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

type NewProductInput struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	SellerID int     `json:"sellerId"`
}

func (h *Handler) newProduct(writer http.ResponseWriter, request *http.Request) {
	var input NewProductInput

	err := json.NewDecoder(request.Body).Decode(&input)
	if err != nil {
		http.Error(writer, models.ErrInvalidInputBody.Error(), http.StatusInternalServerError)
		return
	}

	err = h.services.Products.Create(models.Product{
		Title:    input.Title,
		SellerID: input.SellerID,
		Price:    input.Price,
	})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = writer.Write([]byte("product successfully created"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) updateProduct(writer http.ResponseWriter, request *http.Request) {
	//TODO
	http.Error(writer, "not implemented", http.StatusNotImplemented)
}

func (h *Handler) deleteProduct(writer http.ResponseWriter, request *http.Request) {
	productID := request.URL.Query().Get("id")

	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(writer, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.services.Products.Delete(productIDInt)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = writer.Write([]byte("product successfully deleted"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
