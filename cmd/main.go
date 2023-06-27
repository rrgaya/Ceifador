package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/pubsub"
	"github.com/rrgaya/jesuita/pkg/sender"
)

func main() {
	// Defina o ID do projeto do Google Cloud e o nome do tópico
	projectID := "conversion-toolkit"
	topicName := "nova-conversao"
	subscriptionName := "MySub"

	// Crie um contexto
	ctx := context.Background()

	// Crie um cliente do Pub/Sub
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Falha ao criar o cliente do Pub/Sub: %v", err)
	}

	// Crie uma referência para o tópico
	topic := client.Topic(topicName)

	// Verifique se o tópico existe
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
	// 	Topic: topic,
	// })
	// if err != nil {
	// 	log.Fatalf("Falha ao criar a assinatura: %v", err)
	// }

	// Capture os sinais do sistema para interromper a execução do programa

	urlMsg := make(chan string)

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
					// Processar a mensagem recebida
					UrlToConvert := string(msg.Data)

					urlMsg <- UrlToConvert

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

	// Inicie o servidor HTTP para ouvir na porta 8080
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, Cloud Run!")
		})
		fmt.Println("<<< CEIFADOR: Started :8080 >>>")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Falha ao iniciar o servidor HTTP: %v", err)
		}
	}()

	for msg := range urlMsg {
		fmt.Println("<< CEIFADOR: Verificando mensagens no channel >>>")
		sender.SendMessageToAPI(msg)
	}

	fmt.Println("<< CEIFADOR: Aguardando mensagens... >>>")
	<-stop

	// Feche o cliente do Pub/Sub ao final da execução
	if err := client.Close(); err != nil {
		log.Printf("Erro ao fechar o cliente do Pub/Sub: %v", err)
	}
}
