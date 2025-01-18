package repositories

import (
	"mock-crud-api/config"
	"mock-crud-api/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := config.DB.Find(&products)
	return products, result.Error
}

func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}

func GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	result := config.DB.First(&product, id)
	return product, result.Error
}

func UpdateProduct(product *models.Product) error {
	return config.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return config.DB.Delete(&models.Product{}, id).Error
}
