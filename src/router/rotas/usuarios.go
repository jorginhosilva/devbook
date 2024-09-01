package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	// Cadastra todos os usuários!
	{
		URI: 					"/usuarios",
		Metodo: 				http.MethodPost,
		Funcao: 			controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
	// Busca todos os usuários	
		URI: 				"/usuarios",
		Metodo: 			http.MethodGet,
		Funcao: 			controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
	// Busca usuário por ID	
		URI: 				"/usuarios/{usuarioID}",
		Metodo: 			http.MethodGet,
		Funcao: 			controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
	// Atualiza informações do usuário	
		URI: 				"/usuarios/{usuarioId}",
		Metodo: 			http.MethodPut,
		Funcao: 			controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
	// Deleta informações do usuário	
		URI: 				"/usuarios/{usuarioId}",
		Metodo: 			http.MethodDelete,
		Funcao: 			controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
	
}