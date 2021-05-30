package user

import "time"

type User struct {
	Uuid         string     `json:"uuid" sql:"user_id"`
	Username     string     `json:"username" sql:"username"`
	PasswordHash string     `json:"password_hash" sql:"password_hash"`
	UserRole     string     `json:"user_role" sql:"user_role"`
	CreatedOn    time.Time  `json:"created_on" sql:"created_on"`
	LastLogin    *time.Time `json:"last_login" sql:"last_login"`
}
