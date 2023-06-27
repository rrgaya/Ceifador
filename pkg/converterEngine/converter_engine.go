package converterengine

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func Process(urlToProcess string) {
	fmt.Println("<<< PROCESS INITIATED BY CEIFADOR >>>")

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
		fmt.Println("<<<  CEIFADOR ERROR: Ao tentar clicar no CTA. CHROMEDP >>>")
		fmt.Printf("<<< CEIFADOR ERROR: URLCAMPAING %s >>>\n", urlToProcess)
		fmt.Println("ERROR CTA: ", err)
		return
	}
	fmt.Println("<<< CEIFADOR PROCESS: PASSOU DO CHROMEDP RUN >>>")
	time.Sleep(time.Millisecond * 500)
	fmt.Println("<<<  CEIFADOR LOG: New Conversion >>> ")

}
