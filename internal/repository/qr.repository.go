package repository

import (
	"fmt"
	"gilangaryap/gym-buddy/internal/models"

	"github.com/jmoiron/sqlx"
)

type QrRepositoryInterface interface {
	CreateQRCode(body *models.QRCode) (string, error)
}

type QrRepository struct {
	*sqlx.DB
}

func NewQrRepository(db *sqlx.DB) *QrRepository {
	return &QrRepository{db}
}

func (r *QrRepository) CreateQRCode(body *models.QRCode) (string, error) {
    fmt.Printf("Inserting QR code with SubscriptionID: %d, QrCodeData: %s\n", body.SubscriptionID, body.QrCodeData)

    query := `INSERT INTO qr_codes (subscription_id, qr_code_data) VALUES ($1, $2)`
    _, err := r.Exec(query, body.SubscriptionID, body.QrCodeData)
    if err != nil {
        // Log kesalahan jika terjadi
        fmt.Printf("Error inserting QR code: %v\n", err)
        return "", err
    }
    return body.QrCodeData, nil
}


/* var durationDays int
	switch body.SubOptID {
	case 1:
		durationDays = 30
	case 2:
		durationDays = 180
	case 3:
		durationDays = 360
	default:
		return "", errors.New("invalid sub_opt_id")
	}
	 */