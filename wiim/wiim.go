package wiim

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

type API struct {
	IP      string
	Client  *http.Client
	baseURL string
}

const (
	commandEndpoint = "https://%s/httpapi.asp?command="
	timeout         = 5 * time.Second
	minVolume       = 0
	maxVolume       = 100
	firstPreset     = 1
	lastPreset      = 12
)

func New(ip string) *API {
	client := getClient()
	baseURL := fmt.Sprintf(commandEndpoint, ip)

	return &API{
		Client:  client,
		IP:      ip,
		baseURL: baseURL,
	}
}

func getClient() *http.Client {
	transport := getTransport()

	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
}

func getTransport() *http.Transport {
	tlsConfig := getTLSConfig()

	return &http.Transport{
		TLSClientConfig: tlsConfig,
	}
}

func getTLSConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}
