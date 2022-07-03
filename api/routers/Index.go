package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Set_Root(router *mux.Router) {
	router.HandleFunc("/", Root)

}

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "http://localhost:3000/api/client")
}
