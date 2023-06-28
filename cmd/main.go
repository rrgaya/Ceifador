package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/rrgaya/jesuita/pkg/zeus"
)

func main() {
	// Defina o ID do projeto do Google Cloud e o nome da assinatura do Pub/Sub
	projectID := "conversion-toolkit"
	subscriptionName := "MySub"

	// Crie um cliente do Pub/Sub
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Falha ao criar o cliente do Pub/Sub: %v", err)
	}

	// Crie uma referência para a assinatura do Pub/Sub
	subscription := client.Subscription(subscriptionName)

	// Configure a função de manipulação de mensagens do Pub/Sub
	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		// Processar a mensagem recebida
		fmt.Printf("Mensagem recebida: %s\n", string(msg.Data))
		zeus.Process(string(msg.Data))
		// Confirmar o recebimento da mensagem
		msg.Ack()
		if err != nil {
			log.Printf("Erro ao confirmar o recebimento da mensagem: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Erro ao receber mensagens do Pub/Sub: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Serviço Cloud Run em execução!")
}

func init() {
	http.HandleFunc("/", handler)
}
