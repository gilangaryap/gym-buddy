package repository

import (
	"context"
	"errors"
	"fmt"
	"gilangaryap/gym-buddy/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type SubRepositoryInterface interface {
	CreateDataByUser( uuid string , body *models.Subs)(string , error)
}

type SubRepository struct {
	*sqlx.DB
}

func NewSubRepository(db *sqlx.DB) *SubRepository {
	return &SubRepository{db}
}

func (r *SubRepository) CreateDataByUser(uuid string, body *models.Subs) (string, error) {
    tx, err := r.BeginTx(context.Background(), nil)
    if err != nil {
        fmt.Println("db.BeginTx", err)
        return "", err
    }
    defer tx.Rollback() 

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

    startDate := time.Now()
    endDate := startDate.AddDate(0, 0, durationDays)
    statusID := 1

    var subscriptionID string
    err = tx.QueryRowContext(context.Background(), `INSERT INTO subscription (user_id, status_id, sub_opt_id, start_date, end_date) 
        VALUES ($1, $2, $3, $4, $5) RETURNING id`, uuid, statusID, body.SubOptID, startDate, endDate).Scan(&subscriptionID)

    if err != nil {
        tx.Rollback() 
        return "", err
    }

    if err := tx.Commit(); err != nil {
        return "", err
    }

    return subscriptionID, nil
}