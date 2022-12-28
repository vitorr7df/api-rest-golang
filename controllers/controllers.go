package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitorr7df/api-rest-golang/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasReceitas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Receitas)
}
