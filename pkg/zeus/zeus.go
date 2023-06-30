package zeus

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func Process(ladingpage string, IDTransação string) {
	log.Printf("### CEIFADOR ### >>> INIT PROCESS")
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelTimeout()

	ctx, cancel := chromedp.NewContext(ctxTimeout)
	defer cancel()

	log.Printf("### CEIFADOR ### >>> CHROME RUN PARA: %s \n", IDTransação)
	err := chromedp.Run(ctx,
		chromedp.Navigate(ladingpage),
		chromedp.WaitVisible(".cta__if.button.expanded", chromedp.ByQuery),
		chromedp.Click(".cta__if.button.expanded", chromedp.ByQuery),
	)
	if err != nil {
		log.Printf("### CEIFADOR ERROR ### >>> %q", err)
		log.Fatalln(err)
	}

	log.Printf("### CEIFADOR ### >>> ESPERANDO 5 SEG DEPOIS DO RUN \n")
	time.Sleep(time.Second * 5)

	log.Printf("### CEIFADOR ### >>> LANDINGPAGE: %s \n", ladingpage)
	log.Printf("### CEIFADOR ### >>> TRANSACTION_ID: %s \n", IDTransação)
	log.Printf("### CEIFADOR ### >>> NEW CONVERSION: %s", IDTransação)

}
