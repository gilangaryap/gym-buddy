package repository

import "github.com/jmoiron/sqlx"


type StatusRepositoryInterface interface {
	GetDataByUser(uuid string)(string, error)
}

type StatusRepository struct {
	*sqlx.DB
}

func NewStatusRepository(db *sqlx.DB) *StatusRepository {
	return &StatusRepository{db}
}

func (r *StatusRepository) GetDataByUser(uuid string)(string, error){
	query := `SELECT s2.status_option 
			FROM public."subscription" s
			inner join public."status" s2 ON s.status_id = s2.id
			inner join users u on s.user_id = u.id 
			where u.id = $1`

	var statusOption string
	err := r.DB.QueryRow(query, uuid).Scan(&statusOption)
	if err != nil {
		return "", err
	}

	return statusOption, nil
}