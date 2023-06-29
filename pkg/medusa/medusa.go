package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest(goproxy.DstHostIs("example.com")).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			// Modifica o IP de origem da solicitação
			originalURL := r.URL
			proxyURL, _ := url.Parse("http://new-ip-address:port") // Substitua "new-ip-address" pelo novo IP desejado
			r.URL.Scheme = proxyURL.Scheme
			r.URL.Host = proxyURL.Host
			r.Host = proxyURL.Host

			fmt.Printf("Solicitação GET modificada: %s\n", originalURL.String())

			return r, nil
		},
	)

	log.Fatal(http.ListenAndServe(":8080", proxy))
}
