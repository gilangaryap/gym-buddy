package repository

import (
	models "gilangaryap/gym-buddy/internal/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepositoryInterface interface {
	GetByEmail(email string) (*models.Auth, error)
}

type AuthRepository struct {
	*sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func(r *AuthRepository)GetByEmail(email string)(*models.Auth , error) {
	result := models.Auth{}
	query := `select id , email , password_hash , "role" from users where email = $1`
	
	err := r.Get(&result, query, email)
	if err != nil {
		return nil, err
	}

	return &result , nil
}