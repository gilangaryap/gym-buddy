package models

type QRCode struct {
	SubscriptionID string `json:"SubscriptionID"`
	QrCodeData     string `json:"QrCodeData"`
}
