package models

import (
	"time"
)

type User struct {
	ID        uint64    `json:"id" faker:"-"`
	Name      string    `json:"name" faker:"name"`
	Email     string    `json:"email" faker:"email"`
	CreatedAt time.Time `json:"created_id" faker:"time"`
	UpdatedAt time.Time `json:"updated_id" faker:"time"`
}
