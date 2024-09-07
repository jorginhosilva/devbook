package middlewares

import (
	"api/src/autenticacao"
	respostas "api/src/resposta"
	"log"
	"net/http"
)

// Logger é uma ferramenta que registra informações sobre a execução de um programa
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// O pacote middlewares ficará entre a requisição e a resposta, absorvendo todas as informações
// func (w http.ResponseWriter, r *http.Request)
// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(w, r)
	}
}

