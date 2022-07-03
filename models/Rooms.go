package models

import "gorm.io/gorm"

type Rooms struct {
	gorm.Model
	Price    int
	Reserved bool
}
