package services

import (
	"context"
	"time"

	"github.com/fcsgehrke/finchat/internal/db/entities"
	"github.com/fcsgehrke/finchat/pkg/models"
)

func (s *Service) CreateRoom(ctx context.Context, req *models.CreateRoomRequest) (*models.CreateRoomResponse, error) {
	room, err := s.repo.CreateRoom(ctx, &entities.Room{
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	s.roomManager.NewRoom(ctx, room.RoomID.String())

	return &models.CreateRoomResponse{
		ID:   int(room.ID),
		Name: room.Name,
	}, nil
}

func (s *Service) ListRooms(ctx context.Context) (*models.RoomListResponse, error) {
	rooms, err := s.repo.GetRooms(ctx)
	if err != nil {
		return nil, err
	}

	roomItems := []*models.RoomListItemResponse{}
	for _, room := range rooms {
		roomItems = append(roomItems, &models.RoomListItemResponse{
			ID:   int(room.ID),
			Name: room.Name,
		})
	}

	return &models.RoomListResponse{
		Rooms: roomItems,
	}, nil
}

func (s *Service) GetMessages(ctx context.Context, roomId int) (chan models.NewMessageEvent, error) {
	room, err := s.repo.GetRoom(ctx, roomId)
	if err != nil {
		return nil, err
	}

	return s.roomManager.ReceiveFromRoom(ctx, room.RoomID.String())
}

func (s *Service) SendMessage(ctx context.Context, roomId int, userEmail string, content string, entry bool) error {
	room, err := s.repo.GetRoom(ctx, roomId)
	if err != nil {
		return err
	}

	user, err := s.repo.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return err
	}

	return s.roomManager.Message(ctx, room.RoomID.String(), &models.NewMessageEvent{
		Username: user.Name,
		At:       time.Now(),
		Entry:    entry,
		Content:  content,
	})
}
