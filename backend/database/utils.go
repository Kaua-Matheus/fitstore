package database

import "gorm.io/gorm"


func GetAllData(db *gorm.DB) ([]Product, error) {

	var alldata []Product;
	result := db.Find(&alldata);
	return alldata, result.Error;
	
}