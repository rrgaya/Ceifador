package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func Process(urlToProcess, transactionIDToPrint string) {
	fmt.Println("<<< PROCESS INITIATED >>>")

	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelTimeout()

	ctx, cancel := chromedp.NewContext(ctxTimeout)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(urlToProcess),
		chromedp.WaitVisible(".cta__if.button.expanded", chromedp.ByQuery),
		chromedp.Click(".cta__if.button.expanded", chromedp.ByQuery),
	)
	if err != nil {
		fmt.Println("<<< ERROR: Ao tentar clicar no CTA. CHROMEDP >>>")
		fmt.Printf("<<< ERROR: URLCAMPAING %s >>>\n", urlToProcess)
		fmt.Println("ERROR CTA: ", err)
		return
	}

	fmt.Println("<<< PROCESS: PASSOU DO CHROMEDP RUN >>>")
	time.Sleep(time.Millisecond * 500)

	var response = map[string]string{}
	response["TRANSACTION_ID"] = transactionIDToPrint
	fmt.Printf("Conversão %s gerada.\n", transactionIDToPrint)
}
