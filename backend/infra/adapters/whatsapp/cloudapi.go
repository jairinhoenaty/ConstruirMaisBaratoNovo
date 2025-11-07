package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type CloudAPI struct {
	Token         string
	PhoneNumberID string
	BaseURL       string
	Client        *http.Client
}

func NewCloudAPI() *CloudAPI {
	return &CloudAPI{
		Token:         os.Getenv("WHATSAPP_TOKEN"),
		PhoneNumberID: os.Getenv("WHATSAPP_PHONE_NUMBER_ID"), // ex: 123456789012345
		BaseURL:       "https://graph.facebook.com/v20.0",
		Client:        &http.Client{Timeout: 15 * time.Second},
	}
}

type textMessageReq struct {
	MessagingProduct string      `json:"messaging_product"`
	To               string      `json:"to"`
	Type             string      `json:"type"`
	Text             textMessage `json:"text"`
}
type textMessage struct {
	Body string `json:"body"`
}

func (w *CloudAPI) SendText(toE164, body string) error {
	url := fmt.Sprintf("%s/%s/messages", w.BaseURL, w.PhoneNumberID)
	payload := textMessageReq{
		MessagingProduct: "whatsapp",
		To:               toE164,
		Type:             "text",
		Text:             textMessage{Body: body},
	}
	b, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+w.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := w.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("whatsapp send error: status=%d", resp.StatusCode)
	}
	return nil
}
