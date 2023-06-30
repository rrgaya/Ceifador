package usecase

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func GetTransactionID(urlString string) string {
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err)
	}

	queryParam := u.Query()
	transactionID := queryParam.Get("transaction_id")
	return transactionID
}

func GetURLCampaign(urlGo2Cloud string) (urlCampaign string, transactionID string) {
	URL, _ := url.Parse(urlGo2Cloud)

	client := &http.Client{}

	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		log.Println("### CEIFADOR ERROR ### >>> Erro ao criar req:", err)
	}

	userAgent := GenerateRandomUserAgentAndroid()
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("### CEIFADOR ERROR ### >>> Erro ao fazer DO da REQUEST:", err)
	}
	defer resp.Body.Close()

	transactionID = GetTransactionID(resp.Request.URL.String())
	return resp.Request.URL.String(), transactionID
}
