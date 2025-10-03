package mercadopago

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type paymentQueryResp struct {
	ID                int64   `json:"id"`
	Status            string  `json:"status"`
	StatusDetail      string  `json:"status_detail"`
	TransactionAmount float64 `json:"transaction_amount"`
	Description       string  `json:"description"`
	PaymentMethodID   string  `json:"payment_method_id"`
	ExternalReference string  `json:"external_reference"`
	DateCreated       string  `json:"date_created"`
	DateApproved      string  `json:"date_approved,omitempty"`
	DateLastUpdated   string  `json:"date_last_updated"`
}

type PaymentQueryResult struct {
	PaymentID         int64
	Status            string
	StatusDetail      string
	TransactionAmount float64
	Description       string
	PaymentMethodID   string
	ExternalReference string
	DateCreated       string
	DateApproved      string
	DateLastUpdated   string
}

func (c *MPClient) GetPayment(paymentID int64) (*PaymentQueryResult, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/payments/%d", c.BaseURL, paymentID), nil)
	if err != nil {
		return nil, err
	}

	c.AccessToken = os.Getenv("MERCADOPAGO_ACCESS_TOKEN_TESTE")

	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		var dbg map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&dbg)
		return nil, fmt.Errorf("mercadopago error: status=%d body=%v", resp.StatusCode, dbg)
	}

	var out paymentQueryResp
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return &PaymentQueryResult{
		PaymentID:         out.ID,
		Status:            out.Status,
		StatusDetail:      out.StatusDetail,
		TransactionAmount: out.TransactionAmount,
		Description:       out.Description,
		PaymentMethodID:   out.PaymentMethodID,
		ExternalReference: out.ExternalReference,
		DateCreated:       out.DateCreated,
		DateApproved:      out.DateApproved,
		DateLastUpdated:   out.DateLastUpdated,
	}, nil
}
