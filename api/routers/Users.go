package routers

import (
	"github.com/gorilla/mux"
	"github.com/jeffercbs/api-architecture/controllers"
)

func Set_Client(router *mux.Router) {

	router.HandleFunc("/client", controllers.Get_All_Clients)
	router.HandleFunc("/client", controllers.Get_Client).Methods("GET")
	router.HandleFunc("/client/{id}", controllers.Post_Client).Methods("POST")
	router.HandleFunc("/client/{id}", controllers.Delete_Client).Methods("DELETE")
}
