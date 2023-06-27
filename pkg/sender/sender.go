package sender

import (
	"fmt"
	"log"
	"net/http"
)

func SendMessageToAPI(msg string) {
	urlToGet := "https://toolkit-od4zxa4f4a-uw.a.run.app&site=" + msg
	resp, err := http.Get(urlToGet)
	if err != nil {
		fmt.Printf("<<< ERRROR AO ENVIAR GET REQUEST PARA O CEIFADOR: %s\n", err)
	}

	if resp.StatusCode == 200 {
		return
	}
	log.Println("<<< ERRROR AO ENVIAR GET REQUEST PARA O CEIFADOR APOS 200: ", err)
}
