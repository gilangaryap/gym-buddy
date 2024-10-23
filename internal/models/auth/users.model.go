package models

import (
	"time"
)

type User struct {
	ID           string  	`db:"id" json:"id"`
	Username     string     `db:"username" json:"username" valid:"-"`
	PhoneNumber  string     `db:"phone_number" json:"phone_number" valid:"stringlength(10|13)~phone number maximal 13 karakter"`
	PasswordHash string     `db:"password_hash" json:"password_hash"  valid:"stringlength(5|256)~Password minimal 5 karakter"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at" default:"now()" valid:"-"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at,omitempty" valid:"-"`
	IsActive     bool       `db:"is_active" json:"is_active" default:"true" valid:"-"`
	Role         string     `db:"role" json:"role" default:"user" validate:"oneof=user admin"`
}
