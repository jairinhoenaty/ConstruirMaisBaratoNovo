package mercadopago

import (
	"net/http"
	"os"
	"time"
)

type MPClient struct {
	AccessToken string
	BaseURL     string
	HttpClient  *http.Client
}

func NewMPClient(accessToken string) *MPClient {
	return &MPClient{
		AccessToken: accessToken,
		BaseURL:     os.Getenv("MERCADOPAGO_BASE_URL_API"),
		HttpClient:  &http.Client{Timeout: 15 * time.Second},
	}
}
