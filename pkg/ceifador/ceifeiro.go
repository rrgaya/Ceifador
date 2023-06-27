package ceifador

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/pubsub"
)

func Ceifeiro(MessageCh chan string) {
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

	// Crie uma assinatura durável para o tópico
	subscription := client.Subscription(subscriptionName)

	// subscription, err := client.CreateSubscription(ctx, subscriptionName, pubsub.SubscriptionConfig{
	// 	Topic:       topic,
	// 	AckDeadline: 10 * time.Second, // Prazo para confirmação da mensagem
	// })
	// if err != nil {
	// 	log.Fatalf("Falha ao criar a assinatura: %v", err)
	// }

	// Capture os sinais do sistema para interromper a execução do programa
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Inicie a goroutine para receber mensagens
	go func() {
		for {
			select {
			case <-stop:
				// Encerrar a goroutine ao receber um sinal de interrupção
				return
			default:
				// Receber mensagens do Pub/Sub
				err := subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
					UrlToConvert := string(msg.Data)
					MessageCh <- UrlToConvert
					// Processar a mensagem recebida
					fmt.Printf("Mensagem recebida: %s\n", string(msg.Data))

					// Confirmar o recebimento da mensagem
					msg.Ack()
				})
				if err != nil {
					log.Printf("Erro ao receber mensagens: %v", err)
				}
			}
		}
	}()

	fmt.Println("Aguardando mensagens... Pressione Ctrl+C para sair.")
	<-stop

	// Fechar o cliente do Pub/Sub ao final da execução
	client.Close()
}
