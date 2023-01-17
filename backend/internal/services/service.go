package services

import (
	"context"

	"github.com/fcsgehrke/finchat/internal/db"
	"github.com/fcsgehrke/finchat/pkg/models"
)

type encryptService interface {
	EncryptPassword(password string) (string, error)
	ValidatePassword(password string, savedPassword string) bool
	GenerateToken(user string, id int) (string, error)
}

type roomManager interface {
	NewRoom(ctx context.Context, roomId string) error
	Message(ctx context.Context, roomId string, event *models.NewMessageEvent) error
	ReceiveFromRoom(ctx context.Context, roomId string) (chan models.NewMessageEvent, error)
}

type Service struct {
	repo        db.Repository
	crypt       encryptService
	roomManager roomManager
}

func NewService(repo db.Repository, crypt encryptService, roomManager roomManager) (*Service, error) {
	return &Service{
		repo:        repo,
		crypt:       crypt,
		roomManager: roomManager,
	}, nil
}
