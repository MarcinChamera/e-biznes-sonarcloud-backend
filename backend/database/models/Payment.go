package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	AmountToPay   int `json:"amounttopay"`
	PaymentNumber int `json:"paymentnumber"`
}
