package model

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	ID        uuid.UUID  `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Address   string     `json:"address"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
