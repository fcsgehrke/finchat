package entities

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name   string
	RoomID uuid.UUID
}

func (r *Room) BeforeCreate(tx *gorm.DB) (err error) {
	r.RoomID, err = uuid.NewV4()
	return
}
