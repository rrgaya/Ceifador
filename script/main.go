package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	uri := "https://apiconversion-od4zxa4f4a-rj.a.run.app/santander?aff_click_id=KWAI-07-L"

	for i := 0; i < 100; i++ {
		response, err := http.Get(uri)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Millisecond * 500)

		log.Printf("Status Code: %d - Conversion: %d", response.StatusCode, i)
	}

}
