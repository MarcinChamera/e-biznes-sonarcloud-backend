package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name     string
	Price    int
	ImageUrl string `json:"image_url"`
	// IconName string
	// IconColor string
	// Name string `json:"name"`
	// Price      int    `json:"price"`
	// Icon       string `json:"icon"`
	// Color      string `json:"color"`
	// CategoryID int
	// Category   Category
}
