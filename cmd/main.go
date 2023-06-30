package main

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/rrgaya/ceifador/internal/usecase"
	"github.com/rrgaya/ceifador/pkg/zeus"
)

func main() {
	log.Println("### CEIFADOR ### >>> MAIN STARTED")

	// TODO: Implementar
	// service.Ceifador()

	projectID := "conversion-toolkit"
	subscriptionName := "MySub"

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Falha ao criar o cliente do Pub/Sub: %v", err)
	}

	subscription := client.Subscription(subscriptionName)

	// FIXME: Remover esse receive para uma gorountine e se comunicar via channel para chamar o zeus.Process
	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {

		URI_PROCESS := string(msg.Data)
		urlLanding, transactionID := usecase.GetURLCampaign(URI_PROCESS)

		zeus.Process(urlLanding, transactionID)
		// log.Printf("### CEIFADOR ### >>> PROCESS STATUS: %v \n", processStatus)

		// Essa verificação é uma garantia quer Ack só sera feito se de fato o process retornou true
		// if processStatus { //AQUI DO ENTRA SE processStatus for true
		msg.Ack()
		if err != nil {
			log.Printf("Erro ao confirmar o recebimento da mensagem: %v", err)
			return
		}
		// }

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
}

func init() {
	log.Println("### CEIFADOR ### >>> INICIALIZANDO...")
	http.HandleFunc("/", handler)
}
