package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	var receita models.Receita
	database.DB.First(&receita, id)
	json.NewEncoder(w).Encode(receita)
}

func CriaUmaNovaReceita(w http.ResponseWriter, r *http.Request) {
	var novaReceita models.Receita
	json.NewDecoder(r.Body).Decode(&novaReceita)
	database.DB.Create(&novaReceita)
	json.NewEncoder(w).Encode(novaReceita)
}

func DeletaUmaReceita(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var receita models.Receita
	database.DB.Delete(&receita, id)
	json.NewEncoder(w).Encode(receita)
}

func EditaReceita(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var receita models.Receita
	database.DB.First(&receita, id)
	json.NewDecoder(r.Body).Decode(&receita)
	database.DB.Save(&receita)
	json.NewEncoder(w).Encode(receita)
}
