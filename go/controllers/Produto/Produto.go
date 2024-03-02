package Produto

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/William-Libero/go-web-app/models/produto"
)

var templates = template.Must(template.ParseGlob("../front/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	var produtos = produto.BuscarTodosOsProdutos()
	exibeProdutos(w, produtos)
}

func exibeProdutos(w http.ResponseWriter, produtos []produto.Produto) {
	templates.ExecuteTemplate(w, "Index", produtos)
}

func NovoProduto(w http.ResponseWriter, r *http.Request) {
	novoProduto(w)
}
func novoProduto(w http.ResponseWriter) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		produto.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := produto.ObtemProduto(idDoProduto)
	templates.ExecuteTemplate(w, "Editar", produto)
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		produto.EditaProduto(nome, descricao, precoConvertido, quantidadeConvertidaParaInt, id)
	}
	http.Redirect(w, r, "/", 301)
}
