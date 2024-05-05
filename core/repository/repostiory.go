package repository

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID        uuid.UUID `gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
