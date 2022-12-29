package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vitorr7df/api-rest-golang/database"
	"github.com/vitorr7df/api-rest-golang/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasReceitas(w http.ResponseWriter, r *http.Request) {
	var p []models.Receita // p é um array de Receita
	database.DB.Find(&p)   // buscar todas as listas de receita
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaReceita(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, receita := range models.Receitas {
		if strconv.Itoa(receita.Id) == id {
			json.NewEncoder(w).Encode(receita)
		}
	}
}
