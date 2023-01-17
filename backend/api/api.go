package api

import (
	"context"

	"github.com/fcsgehrke/finchat/pkg/models"
)

type usersService interface {
	CreateUser(ctx context.Context, req *models.UserCreateRequest) (*models.UserCreateResponse, error)
	Login(ctx context.Context, req *models.UserLoginRequest) (*models.UserLoginResponse, error)
}

type roomsService interface {
	CreateRoom(ctx context.Context, req *models.CreateRoomRequest) (*models.CreateRoomResponse, error)
	ListRooms(ctx context.Context) (*models.RoomListResponse, error)
	GetMessages(ctx context.Context, roomId int) (chan models.NewMessageEvent, error)
	SendMessage(ctx context.Context, roomId int, userEmail string, content string, entry bool) error
}

type Handler struct {
	usersService usersService
	roomsService roomsService
	quit         chan bool
}

func NewAPIHandler(usersService usersService, roomsService roomsService, quit chan bool) (*Handler, error) {
	return &Handler{
		usersService: usersService,
		roomsService: roomsService,
		quit:         quit,
	}, nil
}
