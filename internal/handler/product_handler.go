package handler

import (
	// Importa tus paquetes repository y domain
	"basics/Gallery/internal/domain"
	"basics/Gallery/internal/repository"

	"github.com/gofiber/fiber/v3"
)

type ProductHandler struct {
	repo *repository.ProductRepository
}

func NewProductHandler(repo *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) GetProducts(c fiber.Ctx) error {
	products, err := h.repo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al obtener productos"})
	}
	return c.JSON(products)
}

func (h *ProductHandler) CreateProduct(c fiber.Ctx) error {
	var p domain.Product
	if err := c.Bind().Body(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	if err := h.repo.Create(&p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al guardar en base de datos"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Producto creado exitosamente"})
}
