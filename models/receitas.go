package models

type Receita struct {
	Id           int    `json:"id"`
	Nome         string `json:"nome"`
	Ingredientes string `json:"ingredientes"`
}

var Receitas []Receita
