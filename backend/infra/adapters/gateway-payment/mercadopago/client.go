package mercadopago

import (
	"net/http"
	"time"
)

type MPClient struct {
	AccessToken string
	BaseURL     string
	HttpClient  *http.Client
}

func NewMPClient(accessToken, baseURL string) *MPClient {
	return &MPClient{
		AccessToken: accessToken,
		BaseURL:     baseURL,
		HttpClient:  &http.Client{Timeout: 15 * time.Second},
	}
}
