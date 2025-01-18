package services

import (
	"mock-crud-api/models"
	"mock-crud-api/repositories"
)

func GetProducts() ([]models.Product, error) {
	return repositories.GetAllProducts()
}

func AddProduct(product *models.Product) error {
	return repositories.CreateProduct(product)
}

func FindProductByID(id uint) (models.Product, error) {
	return repositories.GetProductByID(id)
}

func ModifyProduct(product *models.Product) error {
	return repositories.UpdateProduct(product)
}

func RemoveProduct(id uint) error {
	return repositories.DeleteProduct(id)
}
