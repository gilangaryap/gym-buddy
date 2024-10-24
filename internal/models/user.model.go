package models

import "time"

type User struct {
	ID           string     `db:"id" json:"id"`
	Username     string     `db:"username" json:"username" valid:"-"`
	Email        string     `db:"email" json:"email" valid:"email~format email tidak valid"`
	PasswordHash string     `db:"password_hash" json:"password_hash"  valid:"stringlength(5|256)~Password minimal 5 karakter"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at" default:"now()" valid:"-"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at,omitempty" valid:"-"`
	IsActive     bool       `db:"is_active" json:"is_active" default:"true" valid:"-"`
	Role         string     `db:"role" json:"role" default:"user" validate:"oneof=user admin"`
}

type Auth struct {
	Id       	 string `db:"id" json:"id" form:"id" valid:"-"`
	Email        string     `db:"email" json:"email" valid:"email~format email tidak valid"`
	PasswordHash string     `db:"password_hash" json:"password_hash"  valid:"stringlength(5|256)~Password minimal 5 karakter"`
	Role     	 string `db:"role" json:"role" form:"role" valid:"-"`
}