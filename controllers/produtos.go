package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/luizclaudioholanda/loja/models"
)

var htmlTemplates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := models.BuscaProdutos()
	htmlTemplates.ExecuteTemplate(w, "Index", produtos)

}

func New(w http.ResponseWriter, r *http.Request) {

	htmlTemplates.ExecuteTemplate(w, "New", nil)

}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter o preco do produto: ", err)
		}

		quantidadeConvertido, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter a quantidade do produto: ", err)
		}

		models.CriaProduto(nome, descricao, quantidadeConvertido, precoConvertido)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")
	models.DeleteProduto(idProduto)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")
	produtoDB := models.BuscaProdutoPorId(idProduto)

	htmlTemplates.ExecuteTemplate(w, "Edit", produtoDB)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro ao converter o id do produto: ", err)
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter o preco do produto: ", err)
		}

		quantidadeConvertido, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter a quantidade do produto: ", err)
		}

		models.UpdateProduto(idConvertido, nome, descricao, quantidadeConvertido, precoConvertido)
	}

	http.Redirect(w, r, "/", 301)

}
