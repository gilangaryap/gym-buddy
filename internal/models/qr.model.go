package models

import "time"

type QRCode struct {
	ID              int       `db:"id" json:"id"`
	SubscriptionID  int       `db:"subscription_id" json:"subscription_id"`
	QrCodeData      string    `db:"qr_code_data" json:"qr_code_data"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
}