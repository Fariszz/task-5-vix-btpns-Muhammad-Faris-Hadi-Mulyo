package models

import "time"

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
