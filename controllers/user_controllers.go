package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffercbs/api-architecture/common"
	"github.com/jeffercbs/api-architecture/models"
)

func Get_All_Clients(w http.ResponseWriter, r *http.Request) {
	clients := []models.Clients{}
	db := common.ConnectionDB()

	db.Find(clients)
	json, _ := json.Marshal(clients)
	common.SendResponse(w, http.StatusOK, json)
}

func Get_Client(w http.ResponseWriter, r *http.Request) {
	client := models.Clients{}
	db := common.ConnectionDB()

	id := mux.Vars(r)["id"]

	db.Find(client, id)
	if client.ID > 0 {
		json, _ := json.Marshal(client)
		common.SendResponse(w, http.StatusOK, json)
	} else {
		common.SendError(w, http.StatusNotFound)
	}
}

func Post_Client(w http.ResponseWriter, r *http.Request) {
	client := models.Clients{}
	db := common.ConnectionDB()
	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		log.Fatal(err)

		common.SendError(w, http.StatusBadRequest)
	}

	err = db.Save(&client).Error

	if err != nil {
		log.Fatal(err)
		common.SendError(w, http.StatusInternalServerError)
	}

	json, _ := json.Marshal(client)
	common.SendResponse(w, http.StatusCreated, json)
}

func Delete_Client(w http.ResponseWriter, r *http.Request) {
	client := models.Clients{}
	db := common.ConnectionDB()

	id := mux.Vars(r)["id"]
	db.Find(&client, id)

	if client.ID == 0 {
		db.Delete(client)
		common.SendResponse(w, http.StatusOK, []byte("{}"))
	} else {
		common.SendError(w, http.StatusNotFound)
	}
}
