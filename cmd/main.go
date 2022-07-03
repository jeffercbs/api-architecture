package main

import (
	"github.com/jeffercbs/api-architecture/api"
	"github.com/jeffercbs/api-architecture/common"
	"github.com/jeffercbs/api-architecture/models"
)

func main() {
	db := common.ConnectionDB()

	db.AutoMigrate(&models.Clients{})
	db.AutoMigrate(&models.Reservation{})
	db.AutoMigrate(&models.Rooms{})

	api.RunServer()
}
