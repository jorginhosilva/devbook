package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)


// JSON retorna uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	// O comando abaixo quer dizer que, o formato do conteúdo é em JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Conversão de dados para JSON
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Erro retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		// Método que retorna a mensagem de erro, o Error() é uma interface em GO
		Erro: erro.Error(),
	})
}