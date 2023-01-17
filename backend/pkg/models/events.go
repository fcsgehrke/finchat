package models

import "time"

type NewMessageEvent struct {
	Username string    `json:"username"`
	At       time.Time `json:"at"`
	Entry    bool      `json:"entry"`
	Content  string    `json:"content"`
}
