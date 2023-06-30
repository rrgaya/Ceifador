package zeus

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func Process(ladingpage string, IDTransação string) {
	log.Printf("### CEIFADOR ### >>> INIT PROCESS")
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelTimeout()

	ctx, cancel := chromedp.NewContext(ctxTimeout)
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(ladingpage),
		chromedp.WaitVisible(".cta__if.button.expanded", chromedp.ByQuery),
		chromedp.Click(".cta__if.button.expanded", chromedp.ByQuery),
	)
	if err != nil {
		log.Printf("### CEIFADOR ### >>> %q", err)
	}
	time.Sleep(time.Second * 5)
	log.Printf("### CEIFADOR ### >>> CONVERSÂO FINALIZADA COM SUCESSO: %s", IDTransação)
}
