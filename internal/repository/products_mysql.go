package repository

import (
	"database/sql"
	"fmt"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/internal/models"
)

type ProductsRepo struct {
	db *sql.DB
}

func NewProductsRepo(db *sql.DB) *ProductsRepo {
	return &ProductsRepo{db: db}
}

func (r *ProductsRepo) Get(id int) (models.Product, error) {
	var product models.Product

	queryGetProduct := fmt.Sprintf("SELECT id, title, price, seller_id FROM %s WHERE id = ?", productsTable)

	row := r.db.QueryRow(queryGetProduct, id)
	if err := row.Scan(&product.ID, &product.Title, &product.Price, &product.SellerID); err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *ProductsRepo) Create(product models.Product) error {
	queryInsertProduct := fmt.Sprintf(
		"INSERT INTO %s (title, price, seller_id) VALUES (?, ?, ?)",
		productsTable,
	)

	_, err := r.db.Exec(queryInsertProduct, product.Title, product.Price, product.SellerID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductsRepo) Delete(productID int) error {
	queryDeleteProduct := fmt.Sprintf("DELETE FROM %s WHERE id = ?", productsTable)

	_, err := r.db.Exec(queryDeleteProduct, productID)
	if err != nil {
		return err
	}

	return nil
}
