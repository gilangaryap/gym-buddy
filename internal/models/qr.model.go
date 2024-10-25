package models

type QRCode struct {
    SubscriptionID int    `json:"SubscriptionID"` // Ensure this matches the JSON key
    QrCodeData     string `json:"QrCodeData"`
}
