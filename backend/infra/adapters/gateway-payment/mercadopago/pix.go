package mercadopago

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PayerIdentification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type PayerAddress struct {
	ZipCode      string `json:"zip_code"`
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	FederalUnit  string `json:"federal_unit"`
}

type Payer struct {
	FirstName      string              `json:"first_name"`
	LastName       string              `json:"last_name"`
	Email          string              `json:"email"`
	Identification PayerIdentification `json:"identification"`
	Address        PayerAddress        `json:"address"`
}

type createPixPaymentReq struct {
	TransactionAmount float64 `json:"transaction_amount"`
	Description       string  `json:"description"`
	PaymentMethodID   string  `json:"payment_method_id"`
	ExternalReference string  `json:"external_reference"`
	NotificationURL   string  `json:"notification_url,omitempty"`
	Payer             Payer   `json:"payer"`
}

type createPixPaymentResp struct {
	ID                 int64  `json:"id"`
	Status             string `json:"status"`
	PointOfInteraction struct {
		TransactionData struct {
			QRCode       string `json:"qr_code"`
			QRCodeBase64 string `json:"qr_code_base64"`
		} `json:"transaction_data"`
	} `json:"point_of_interaction"`
}

type PixCreateResult struct {
	PaymentID    int64
	Status       string
	QRCode       string
	QRCodeBase64 string
}

type PixPaymentInput struct {
	Amount          float64
	Description     string
	ExternalRef     string
	NotificationURL string
	IdempotencyKey  string
	Payer           Payer
}

func (c *MPClient) CreatePixPayment(input PixPaymentInput) (*PixCreateResult, error) {
	body := createPixPaymentReq{
		TransactionAmount: input.Amount,
		Description:       input.Description,
		PaymentMethodID:   "pix",
		ExternalReference: input.ExternalRef,
		NotificationURL:   input.NotificationURL,
		Payer:             input.Payer,
	}
	b, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", c.BaseURL+"/v1/payments", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	// c.AccessToken = os.Getenv("MERCADOPAGO_ACCESS_TOKEN")

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	if input.IdempotencyKey != "" {
		req.Header.Set("X-Idempotency-Key", input.IdempotencyKey)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)

	if resp.StatusCode >= 300 {
		var dbg map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&dbg)
		fmt.Println("mercadopago error: status=%d body=%v", resp.StatusCode, dbg)
		return nil, fmt.Errorf("Erro ao processar pagamento. Por favor, tente novamente mais tarde.")
	}

	var out createPixPaymentResp
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &PixCreateResult{
		PaymentID:    out.ID,
		Status:       out.Status,
		QRCode:       out.PointOfInteraction.TransactionData.QRCode,
		QRCodeBase64: out.PointOfInteraction.TransactionData.QRCodeBase64,
	}, nil
}
