package sender

import (
	"fmt"
	"log"
	"net/http"
)

func SendMessageToAPI(msg string) {

	// urlToGet := "https://toolkit-od4zxa4f4a-uw.a.run.app/engine&site=" + msg
	urlToGet := "http://localhost:8080/engine"

	client := &http.Client{}

	req, err := http.NewRequest("GET", urlToGet, nil)
	if err != nil {
		fmt.Println("<<< ERROR AO MONTAR A REQUISÇÃO >>>", err)
	}
	fmt.Printf("urlToGet: %s\n", urlToGet)
	req.Header.Set("X-SITE", urlToGet)

	resp, err := client.Do(req)
	if resp != nil {
		fmt.Println("<<< ERROR AO EXECUTAR A REQUISÇÃO/RESPONSE >>>", err)
	}

	log.Println("<<< CEIFADOR RESPONSE %d", resp.StatusCode)

}
