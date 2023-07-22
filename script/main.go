package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rrgaya/ceifador/pkg/agentfy"
)

func MakeRequest(urlSrc string) int {
	client := &http.Client{}

	req, err := http.NewRequest("GET", urlSrc, nil)
	if err != nil {
		log.Println("### SCRIPT  ### >>> Erro ao criar req:", err)
	}

	userAgent := agentfy.GenerateAgentfy()
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("### SCRIPT ### >>> Erro ao fazer DO da REQUEST:", err)
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
func main() {
	uri := "https://apiconversion-od4zxa4f4a-rj.a.run.app/santander?aff_click_id=KWAI-MOBILE-G-1"
	// uri := "https://grupodenegocios.com.br/versao6"
	for i := 0; i < 10; i++ {
		statusCode := MakeRequest(uri)

		time.Sleep(time.Millisecond * 200)

		log.Printf("Status Code: %d - Conversion: %d", statusCode, i)
	}

}
