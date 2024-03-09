package interfaces

import "gorm.io/gorm"

type Example struct {
	gorm.Model
	Some string
}
