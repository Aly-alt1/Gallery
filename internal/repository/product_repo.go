package repository

import (
	"basics/Gallery/internal/domain"
	"database/sql"
	// Importa tu paquete domain según la ruta de tu proyecto, ej: "Gallery/internal/domain"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Crear un producto
func (r *ProductRepository) Create(p *domain.Product) error {
	_, err := r.db.Exec("INSERT INTO PRODUCT (NAME, PRICE) VALUES (?, ?)", p.Name, p.Price)
	return err
}

// Leer todos los productos
func (r *ProductRepository) GetAll() ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT ID, NAME, PRICE FROM PRODUCT")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err == nil {
			products = append(products, p)
		}
	}
	return products, nil
}
