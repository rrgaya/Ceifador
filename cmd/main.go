package main

import (
	"fmt"

	"github.com/rrgaya/jesuita/pkg/ceifador"
	converterengine "github.com/rrgaya/jesuita/pkg/converterEngine"
)

func main() {
	chMsg := make(chan string)
	ceifador.Ceifeiro(chMsg)

	for urlToProcess := range chMsg {
		converterengine.Process(urlToProcess)
	}
	fmt.Println("<<< Ceifeiro FINALIZADO >>>")
}
