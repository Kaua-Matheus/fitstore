package database

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


func GetAllProduct(db *gorm.DB) ([]Product, error) {

	var alldata []Product;
	result := db.Find(&alldata);
	return alldata, result.Error;

}

func AddProduct(db *gorm.DB, product Product) (error) {
	
	result := db.Create(&product); if result.Error != nil {
		return fmt.Errorf("error trying to add the register: %w", result.Error);
	} else {
		return nil;
	}

}

func UpdateProduct(db *gorm.DB, id uuid.UUID, product Product) (error) {

	result := db.Model(&Product{}).Where("id = ?", id).Updates(product);
	if result.Error != nil {
		return fmt.Errorf("error trying to update the data: %s", result.Error);
	}

	return nil;
}

func GetAllImage(db *gorm.DB) ([]ProductImage, error) {

	allImages := []ProductImage{};
	result := db.Find(&allImages);

	return allImages, result.Error;
}