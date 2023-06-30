package provider

import (
	"context"
	"fmt"
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

	// Cria um cliente pubsub
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Erro ao criar o cliente pubsub: %v", err)
	}

	// Cria um subscriber
	sub := client.Subscription(subscriptionID)

	// Cria um canal com buffer de tamanho 1
	msgCh := make(chan *pubsub.Message, 1)

	// Inicia uma goroutine para receber mensagens
	go func() {
		for {
			// Recebe uma mensagem do pubsub
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
			log.Println("### CEIFADOR ### >>> SLEEP DE 3 SEG")
			time.Sleep(3 * time.Second)
			fmt.Printf("Mensagem recebida: %s\n", string(msg.Data))

			URI_PROCESS := string(msg.Data)
			urlLanding, transactionID := usecase.GetURLCampaign(URI_PROCESS)
			zeus.Process(urlLanding, transactionID)

			// Marca a mensagem como concluída
			msg.Ack()

		case <-ctx.Done():
			// Contexto cancelado, encerra o loop
			return
		}
	}
}
