package provider

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/rrgaya/ceifador/internal/usecase"
	"github.com/rrgaya/ceifador/pkg/zeus"
)

const (
	projectID      = "conversion-toolkit"
	subscriptionID = "MySub"
)

func GetMessages() {
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Erro ao criar o cliente pubsub: %v", err)
	}

	sub := client.Subscription(subscriptionID)

	msgCh := make(chan *pubsub.Message, 1)

	go func() {
		for {
			err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
				msgCh <- msg
			})
			if err != nil {
				log.Printf("Erro ao receber a mensagem: %v", err)
				time.Sleep(2 * time.Second) // Espera 2 segundos antes de tentar novamente
				continue
			}
		}
	}()

	for {
		select {
		case msg := <-msgCh:

			// Processa a mensagem recebida
			// log.Println("### CEIFADOR ### >>> SLEEP DE 3 SEG")
			time.Sleep(3 * time.Second)

			URI_PROCESS := string(msg.Data)
			urlLanding, transactionID, affid := usecase.GetURLCampaign(URI_PROCESS)
			zeus.Process(urlLanding, transactionID, affid)

			msg.Ack()

		case <-ctx.Done():
			// Contexto cancelado, encerra o loop
			return
		}
	}
}
