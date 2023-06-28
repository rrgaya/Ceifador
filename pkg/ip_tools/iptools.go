package iptools

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func MedusaIP() {
	// Lista de IPs
	ips := []string{
		"192.168.0.1",
		"192.168.0.2",
		"192.168.0.3",
	}

	// URL de destino
	url := "http://example.com"

	// Cliente HTTP
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Loop pelas requisições
	for _, ip := range ips {
		// Criar um novo transporte com o endereço IP atual
		transport := &http.Transport{
			Proxy: http.ProxyURL(nil),
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				LocalAddr: &net.TCPAddr{
					IP: net.ParseIP(ip),
				},
			}).DialContext,
		}

		// Atribuir o transporte ao cliente HTTP
		client.Transport = transport

		// Fazer a requisição GET
		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("Erro na requisição: %s\n", err.Error())
			continue
		}

		// Ler o corpo da resposta
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("Erro ao ler o corpo da resposta: %s\n", err.Error())
			continue
		}

		// Exibir o resultado
		fmt.Printf("Resposta do IP %s: %s\n", ip, string(body))
	}
}
