package repository

import (
	models "gilangaryap/gym-buddy/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryInterface interface {
	CreateData(body *models.User) (string, error)
}

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateData(body *models.User) (string, error) {
	query := `INSERT INTO public.users (username, email , password_hash) VALUES ( :username ,:email , :password_hash)`

	_, err := r.NamedExec(query, body)
	if err != nil {
		return "", err
	}

	return "Create data success", nil
}