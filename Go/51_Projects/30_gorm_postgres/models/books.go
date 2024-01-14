package models

import "gorm.io/gorm"

type Books struct {
	ID        uint    `json:"id" gorm:"primary key;autoIncrement"`
	Barcode   *string `json:"barcode"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
	PrintDate *string `json:"print_date"`
	Rerelease *bool   `json:"rerelease"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	if err != nil {
		return err
	}
	return nil
}
