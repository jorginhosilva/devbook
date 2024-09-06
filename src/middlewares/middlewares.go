package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger é uma ferramenta que registra informações sobre a execução de um programa
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// O pacote middlewares ficará entre a requisição e a resposta, absorvendo todas as informações
// func (w http.ResponseWriter, r *http.Request)
// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando...")
		next(w, r)
	}
}

