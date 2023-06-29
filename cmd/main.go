package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/rrgaya/ceifador/internal/usecase"
	"github.com/rrgaya/ceifador/pkg/zeus"
)

func main() {
	log.Println("### CEIFADOR ### >>> MAIN")
	projectID := "conversion-toolkit"
	subscriptionName := "MySub"

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Falha ao criar o cliente do Pub/Sub: %v", err)
	}

	subscription := client.Subscription(subscriptionName)

	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {

		fmt.Printf("Mensagem recebida: %s\n", string(msg.Data))

		URI_PROCESS := string(msg.Data)
		log.Printf("### CEIFADOR ### >>> URI_PROCESS: %s \n", URI_PROCESS)
		urlLanding, transactionID := usecase.GetURLCampaign(URI_PROCESS)
		log.Printf("### CEIFADOR ### >>> URI_LANDING: %s \n", urlLanding)
		log.Printf("### CEIFADOR ### >>> TRANSACTIONID: %s \n", transactionID)
		zeus.Process(urlLanding, transactionID)

		msg.Ack()
		if err != nil {
			log.Printf("Erro ao confirmar o recebimento da mensagem: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Erro ao receber mensagens do Pub/Sub: %v", err)
	}

	log.Printf("Serviço Cloud Run em execução na porta 8001")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor HTTP: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("### CEIFADOR ### >>> HANDLER")
	fmt.Fprint(w, "Serviço Cloud Run em execução!")
}

func init() {
	log.Println("### CEIFADOR ### >>> FUNC INIT")
	http.HandleFunc("/", handler)
}
