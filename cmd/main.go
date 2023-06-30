package main

import (
	"log"
	"net/http"

	"github.com/rrgaya/ceifador/provider"
)

func main() {
	log.Println("### CEIFADOR ### >>> MAIN STARTED")
	provider.GetMessages()
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("### CEIFADOR ### >>> HANDLER")
}

func init() {
	http.HandleFunc("/", handler)
	log.Println("### CEIFADOR ### >>> INICIALIZANDO...")
}
