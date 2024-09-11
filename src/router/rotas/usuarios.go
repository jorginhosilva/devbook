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
		RequerAutenticacao: true,
	},
	{
	// Busca usuário por ID	
		URI: 				"/usuarios/{usuarioID}",
		Metodo: 			http.MethodGet,
		Funcao: 			controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
	// Atualiza informações do usuário	
		URI: 				"/usuarios/{usuarioId}",
		Metodo: 			http.MethodPut,
		Funcao: 			controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
	// Deleta informações do usuário	
		URI: 				"/usuarios/{usuarioId}",
		Metodo: 			http.MethodDelete,
		Funcao: 			controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{usuariosId}/seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	
}