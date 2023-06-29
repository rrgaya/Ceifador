package usecase

import (
	"fmt"
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
	fmt.Println("### CEIFADOR ### >>> TransactionID: ", transactionID)

	return transactionID
}

func GetURLCampaign(urlGo2Cloud string) (urlCampaign string, transactionID string) {
	URL, _ := url.Parse(urlGo2Cloud)

	client := &http.Client{}

	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		fmt.Println("### CEIFADOR ### >>> Erro ao criar req:", err)
	}

	userAgent := GenerateRandomUserAgentAndroid()

	fmt.Println("User-Agent:", userAgent)
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("### CEIFADOR ### >>> Erro ao fazer DO da REQUEST:", err)
	}
	defer resp.Body.Close()

	transactionID = GetTransactionID(URL.String())

	return resp.Request.URL.String(), transactionID
}
