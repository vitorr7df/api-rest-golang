package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vitorr7df/api-rest-golang/controllers"
	"github.com/vitorr7df/api-rest-golang/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/receitas", controllers.TodasReceitas).Methods("Get")
	r.HandleFunc("/api/receitas/{id}", controllers.RetornaUmaReceita).Methods("Get")
	r.HandleFunc("/api/receitas", controllers.CriaUmaNovaReceita).Methods("Post")
	r.HandleFunc("/api/receitas/{id}", controllers.DeletaUmaReceita).Methods("Delete")
	r.HandleFunc("/api/receitas/{id}", controllers.EditaReceita).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
