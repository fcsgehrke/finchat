package entities

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID  int
	RoomID  int
	Content string
}
