package models

import "time"

type Token struct {
	ID        int       `json:"id"`
	Token     string    `json:"token"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
