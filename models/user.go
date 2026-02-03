package models

import "time"

type User struct {
	Username  string    `json:"username"`
	Id        int       `json:"id"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}