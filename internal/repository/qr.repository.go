package repository

import (
	"context"
	"fmt"
	"gilangaryap/gym-buddy/internal/models"

	"github.com/jmoiron/sqlx"
)

type QrRepositoryInterface interface {
	CreateQRCode( body *models.QRCode) (string, error)
}

type QrRepository struct {
	*sqlx.DB
}

func NewQrRepository(db *sqlx.DB) *QrRepository {
	return &QrRepository{db}
}

func (r *QrRepository) CreateQRCode( body *models.QRCode) (string, error) {
	tx, err := r.BeginTx(context.Background(),nil)
	if err != nil {
		fmt.Println("db.BeginTx", err)
	} 
	
	var qrCodeID string
	err = tx.QueryRowContext(context.Background(), 
	`INSERT INTO qr_codes (subscription_id, qr_code_data) VALUES ($1, $2) RETURNING id`, 
	body.SubscriptionID, body.QrCodeData).Scan(&qrCodeID)

	 if err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

    return qrCodeID, nil
}