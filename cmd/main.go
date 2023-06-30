package main

import (
	"log"
	"net/http"

	"github.com/rrgaya/ceifador/provider"
)

func main() {
	log.Println("### CEIFADOR ### >>> MAIN STARTED")

	provider.GetMessages()

	// TODO: Implementar
	// service.Ceifador()

	// projectID := "conversion-toolkit"
	// subscriptionName := "MySub"

	// ctx := context.Background()
	// client, err := pubsub.NewClient(ctx, projectID)
	// if err != nil {
	// 	log.Fatalf("Falha ao criar o cliente do Pub/Sub: %v", err)
	// }

	// subscription := client.Subscription(subscriptionName)

	// for {
	// 	log.Println("### CEIFADOR ### >>> Sleep de 10 seg para puxar msgs")
	// 	time.Sleep(time.Second * 1)
	// 	// FIXME: Remover esse receive para uma gorountine e se comunicar via channel para chamar o zeus.Process
	// 	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
	// 		log.Println("### CEIFADOR ### >>> RECEIVING: ", string(msg.Data))

	// 		URI_PROCESS := string(msg.Data)
	// 		urlLanding, transactionID := usecase.GetURLCampaign(URI_PROCESS)

	// 		go zeus.Process(urlLanding, transactionID)
	// 		// log.Printf("### CEIFADOR ### >>> PROCESS STATUS: %v \n", processStatus)

	// 		// Essa verificação é uma garantia quer Ack só sera feito se de fato o process retornou true
	// 		msg.Ack()

	// 	})
	// 	if err != nil {
	// 		log.Fatalf("Erro ao receber mensagens do Pub/Sub: %v", err)
	// 	}

	// }

	// // log.Printf("Serviço Cloud Run em execução na porta 8001")
	// // if err := http.ListenAndServe(":8080", nil); err != nil {
	// // 	log.Fatalf("Erro ao iniciar o servidor HTTP: %v", err)
	// // }
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("### CEIFADOR ### >>> HANDLER")
}

func init() {
	log.Println("### CEIFADOR ### >>> INICIALIZANDO...")
	http.HandleFunc("/", handler)
}
