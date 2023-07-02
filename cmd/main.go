package main

import (
	"log"
	"net/http"

	"github.com/rrgaya/ceifador/provider"
)

func main() {
	log.Println("### CEIFADOR ### >>> MAIN STARTED")
	provider.GetMessages()

	log.Printf("Serviço Cloud Run em execução na porta 8001")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor HTTP: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("### CEIFADOR ### >>> HANDLER")
}

func init() {
	http.HandleFunc("/", handler)
	log.Println("### CEIFADOR ### >>> INICIALIZANDO...")
}
