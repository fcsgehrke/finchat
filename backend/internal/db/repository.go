package db

import (
	"context"

	"github.com/fcsgehrke/finchat/internal/db/entities"
)

type Repository interface {
	// Rooms
	CreateRoom(ctx context.Context, room *entities.Room) (*entities.Room, error)
	GetRoom(ctx context.Context, id int) (*entities.Room, error)
	GetRooms(ctx context.Context) ([]*entities.Room, error)

	// Message
	CreateMessage(ctx context.Context, message *entities.Message) (*entities.Message, error)
	GetLastMessagesByRoom(ctx context.Context, roomID int, count int) ([]*entities.Message, error)

	// Users
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
}
