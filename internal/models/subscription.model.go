package models

import (
	"time"
)

type Subs struct {
	ID        uint      `db:"id" json:"id"`
	UserID    string	`db:"user_id" json:"user_id" valid:"-"`
	StatusID  int       `db:"status_id" json:"status_id" valid:"required~Status ID tidak boleh kosong"`
	SubOptID  int       `db:"sub_opt_id" json:"SubOptID" valid:"required~Subscription Option ID tidak boleh kosong"`
	StartDate time.Time `db:"start_date" json:"start_date" valid:"required~Tanggal mulai tidak boleh kosong"`
	EndDate   time.Time `db:"end_date" json:"end_date" valid:"required~Tanggal akhir tidak boleh kosong"`
}