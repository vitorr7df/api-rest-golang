package main

import (
	"fmt"

	"github.com/vitorr7df/api-rest-golang/database"
	"github.com/vitorr7df/api-rest-golang/models"
	"github.com/vitorr7df/api-rest-golang/routes"
)

func main() {
	models.Receitas = []models.Receita{
		{Id: 1, Nome: "Soja", Ingredientes: "Soja, trigo, temperos"},
		{Id: 2, Nome: "Grão de bico", Ingredientes: "Grão de bico, farinha de trigo, temperos"},
	}

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
