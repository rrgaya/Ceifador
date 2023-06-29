package usecase

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

func GetCustomIP() int {
	randomIP := rand.Intn(100000000)
	return randomIP
}

func GetTransactionID(urlString string) string {
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err)
	}

	queryParam := u.Query()
	transactionID := queryParam.Get("transaction_id")
	fmt.Println("TransactionID: ", transactionID)

	return transactionID

}
func UserAgentGenerator() string {
	user_agents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.81 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.81 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.96 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.96 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	}

	randomIndex := rand.Intn(len(user_agents))

	return user_agents[randomIndex]
}

func GetURLCampaign(urlGo2Cloud string) (urlCampaign string) {
	URL, _ := url.Parse(urlGo2Cloud)

	client := &http.Client{}

	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		fmt.Println("Erro ao criar a solicitação:", err)
	}

	my_uuid := uuid.New()
	userAgent := my_uuid.String()
	customIP := GetCustomIP()

	fmt.Println("User-Agent:", userAgent)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("X-Custom-IP", fmt.Sprintf("%d", customIP))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a solicitação:", err)
	}
	defer resp.Body.Close()

	return resp.Request.URL.String()
}
