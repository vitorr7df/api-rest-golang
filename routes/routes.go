package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorr7df/api-rest-golang/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/receitas", controllers.TodasReceitas).Methods("Get")
	r.HandleFunc("/api/receitas/{id}", controllers.RetornaUmaReceita).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
