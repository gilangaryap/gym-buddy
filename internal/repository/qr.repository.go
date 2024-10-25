package repository

import (
	"errors"
	"gilangaryap/gym-buddy/internal/models"
	"time"

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

	var durationDays int
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
	
	startSubAt := time.Now()
	expiryAt := startSubAt.AddDate(0, 0, durationDays)

	body.StartSubAt = startSubAt
	body.ExpiryAt = expiryAt
	
	query := `INSERT INTO qr_codes (url, user_id, sub_opt_id , start_sub_at , expiry_at ) VALUES (:url, :user_id, :sub_opt_id, :start_sub_at, :expiry_at)`
	_, err := r.NamedExec(query, body)
	if err != nil {
		return "", err
	}

	return body.URL, nil 
}