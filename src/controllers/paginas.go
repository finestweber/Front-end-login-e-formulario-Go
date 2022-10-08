package controllers

import (
	"net/http"
	"go-webapp/src/utils"
	
)
func CarregaTelaLogin(w http.ResponseWriter, r *http.Request){
	utils.ExecutarTemplate(w, "login.html", nil)
}
// carrega pagina de cadastro
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request){
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}