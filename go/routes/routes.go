package routes

import (
	"net/http"

	"github.com/William-Libero/go-web-app/controllers/Produto"
)

func IniciaServer() {
	http.HandleFunc("/", Produto.Index)
	http.HandleFunc("/new", Produto.NovoProduto)
	http.HandleFunc("/insert", Produto.Insert)
	http.HandleFunc("/delete", Produto.Delete)
	http.HandleFunc("/editar", Produto.Editar)
	http.HandleFunc("/editarProduto", Produto.UpdateProduto)
	http.ListenAndServe(":8000", nil)
}
