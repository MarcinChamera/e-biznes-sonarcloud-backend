package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	PaymentNumber int `json:"paymentnumber"`
}
