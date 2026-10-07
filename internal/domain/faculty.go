package domain

import (
	"time"

	"github.com/google/uuid"
)



type Faculty struct {
	ID   int    `json:"id,omitempty"`
	UUID  uuid.UUID `json:"uuid"`
	Name string `json:"name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}


