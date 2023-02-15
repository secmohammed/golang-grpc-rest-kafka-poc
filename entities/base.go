package entities

import (
    "github.com/google/uuid"
    "time"
)

type BaseModel struct {
    ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
    CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
    UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`
}
