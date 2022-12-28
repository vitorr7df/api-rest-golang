package main

import (
	"fmt"

	"github.com/vitorr7df/api-rest-golang/models"
	"github.com/vitorr7df/api-rest-golang/routes"
)

func main() {
	models.Receitas = []models.Receita{
		{Nome: "Soja", Ingredientes: "Soja, trigo, temperos"},
		{Nome: "Grão de bico", Ingredientes: "Grão de bico, farinha de trigo, temperos"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
