package service

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/rrgaya/ceifador/internal/usecase"
	"github.com/rrgaya/ceifador/pkg/zeus"
)

func Ceifador() {

	projectID := "conversion-toolkit"
	topicName := "nova-conversao"
	subscriptionName := "MySub"

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Falha ao criar o cliente do Pub/Sub: %v", err)
	}

	topic := client.Topic(topicName)

	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Falha ao verificar se o tópico existe: %v", err)
	}
	if !exists {
		log.Fatalf("O tópico %s não existe", topicName)
	}

	subscription, err := client.CreateSubscription(ctx, subscriptionName, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 10 * time.Second, // Prazo para confirmação da mensagem
	})
	if err != nil {
		log.Fatalf("Falha ao criar a assinatura: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				err := subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
					URI_PROCESS := string(msg.Data)
					urlLanding, transactionID := usecase.GetURLCampaign(URI_PROCESS)
					zeus.Process(urlLanding, transactionID)
					msg.Ack()
				})
				if err != nil {
					log.Printf("Erro ao receber mensagens: %v", err)
				}
			}
		}
	}()
	<-stop
	client.Close()

}
