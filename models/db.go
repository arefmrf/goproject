package models

import "time"

type DBProduct struct {
	ID              uint
	ProductID       string
	Title           string
	Discount        int
	Price           int
	DiscountedPrice int
	EndAt           time.Time
}
