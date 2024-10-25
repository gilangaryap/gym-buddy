package models

import "time"

type QRCode struct {
	ID         int       `db:"id"`
	URL        string    `db:"url"`
	UserID     string    `db:"user_id"`
	SubOptID   int       `db:"sub_opt_id"`
	StartSubAt time.Time  `db:"start_sub_at"` 
	ExpiryAt   time.Time  `db:"expiry_at"`     
}