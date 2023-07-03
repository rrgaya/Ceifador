package usecase

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/rrgaya/ceifador/pkg/agentfy"
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

func GetAffliateClickID(urlGo2Cloud string) string {
	uri, err := url.Parse(urlGo2Cloud)
	if err != nil {
		log.Println("### CEIFADOR ERROR ### >>> Erro pegar Affliate Click ID:", err)
	}
	affliate_clickID := uri.Query().Get("aff_click_id")
	return affliate_clickID
}

func GetURLCampaign(urlGo2Cloud string) (urlCampaign string, transactionID string, affID string) {
	URL, _ := url.Parse(urlGo2Cloud)

	client := &http.Client{}

	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		log.Println("### CEIFADOR ERROR ### >>> Erro ao criar req:", err)
	}

	userAgent := agentfy.GenerateAgentfy()
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("### CEIFADOR ERROR ### >>> Erro ao fazer DO da REQUEST:", err)
	}
	defer resp.Body.Close()

	transactionID = GetTransactionID(resp.Request.URL.String())
	affID = GetAffliateClickID(urlGo2Cloud)
	return resp.Request.URL.String(), transactionID, affID
}
