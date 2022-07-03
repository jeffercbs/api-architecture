package models

import "gorm.io/gorm"

type Clients struct {
	gorm.Model
	Level         int     `json:"level"`
	Avatar        string  `json:"avatar"`
	Name          string  `json:"name"`
	Username      string  `json:"username"`
	Document      int     `json:"document"`
	Email         *string `json:"email"`
	Country       string  `json:"country"`
	Cel           int     `json:"cel"`
	Millas        int     `json:"millas"`
	Ages          int     `json:"ages"`
	ReservationId int
	Reservation   Reservation `gorm:"foreingKey:ReservationId"`
}
