package postgres

import (
	"context"

	"github.com/fcsgehrke/finchat/internal/db/entities"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*PostgresRepository, error) {
	return &PostgresRepository{
		db: db,
	}, nil
}

// Rooms
func (r *PostgresRepository) CreateRoom(ctx context.Context, room *entities.Room) (*entities.Room, error) {
	if err := r.db.WithContext(ctx).Create(&room).Error; err != nil {
		return nil, err
	}
	return room, nil
}

func (r *PostgresRepository) GetRoom(ctx context.Context, id int) (*entities.Room, error) {
	var room entities.Room
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *PostgresRepository) GetRooms(ctx context.Context) ([]*entities.Room, error) {
	var rooms []*entities.Room
	err := r.db.WithContext(ctx).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

// Users
func (r *PostgresRepository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresRepository) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := r.db.WithContext(ctx).Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Message
func (r *PostgresRepository) CreateMessage(ctx context.Context, message *entities.Message) (*entities.Message, error) {
	if err := r.db.WithContext(ctx).Create(message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (r *PostgresRepository) GetLastMessagesByRoom(ctx context.Context, roomID int, count int) ([]*entities.Message, error) {
	var messages []*entities.Message
	err := r.db.WithContext(ctx).Where("room_id = ?", roomID).Limit(count).Order("created_at DESC").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
