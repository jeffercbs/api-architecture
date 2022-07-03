package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffercbs/api-architecture/api/routers"
)

func RunServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/api")

	routers.Set_Client(router)
	routers.Set_Root(router)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 3000")
	log.Println("http://localhost:3000")
	log.Println(server.ListenAndServe())

}
