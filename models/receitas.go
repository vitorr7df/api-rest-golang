package models

type Receita struct {
	Nome         string `json:"nome"`
	Ingredientes string `json:"ingredientes"`
}

var Receitas []Receita
