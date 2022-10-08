package rotas

import (
	"net/http"
	"go-webapp/src/controllers"
)
var RotasUsuarios = []Rota {
	{
		URI: "/criar-usuario",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutenticacao: false,
	},
}