package routes

import (
	"log"
	"net/http"

	"github.com/vitorr7df/api-rest-golang/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/receitas", controllers.TodasReceitas)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
