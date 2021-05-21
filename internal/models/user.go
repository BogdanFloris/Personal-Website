package models

import "time"

type User struct {
	Uuid         string     `json:"uuid" sql:"user_id"`
	Username     string     `json:"username" sql:"username"`
	PasswordHash string     `json:"password_hash" sql:"password_hash"`
	CreatedOn    time.Time  `json:"created_on" sql:"created_on"`
	LastLogin    *time.Time `json:"last_login" sql:"last_login"`
}
