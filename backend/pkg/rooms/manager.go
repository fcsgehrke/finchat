package rooms

import (
	"context"
	"encoding/json"

	"github.com/fcsgehrke/finchat/internal/db/entities"
	"github.com/fcsgehrke/finchat/pkg/models"
	"github.com/gofrs/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type roomsDb interface {
	GetRooms(ctx context.Context) ([]*entities.Room, error)
}

type RoomManager struct {
	connChannel *amqp.Channel
}

func NewRoomsManager(amqpAddr string, roomsDb roomsDb) (*RoomManager, error) {
	conn, err := amqp.Dial(amqpAddr)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	rooms, err := roomsDb.GetRooms(context.Background())
	if err != nil {
		return nil, err
	}

	for _, room := range rooms {
		err := ch.ExchangeDeclare(room.RoomID.String(), "fanout", false, false, false, false, nil)
		if err != nil {
			return nil, err
		}
	}

	return &RoomManager{
		connChannel: ch,
	}, nil
}

func (r *RoomManager) NewRoom(ctx context.Context, roomId string) error {
	err := r.connChannel.ExchangeDeclare(roomId, "fanout", false, true, false, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomManager) Message(ctx context.Context, roomId string, event *models.NewMessageEvent) error {
	return publishMessage(ctx, r.connChannel, roomId, event)
}

func (r *RoomManager) ReceiveFromRoom(ctx context.Context, roomId string) (chan models.NewMessageEvent, error) {
	output := make(chan models.NewMessageEvent)

	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	q, err := r.connChannel.QueueDeclare(uid.String(), false, true, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = r.connChannel.QueueBind(q.Name, "", roomId, false, nil)
	if err != nil {
		return nil, err
	}

	msgs, err := r.connChannel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		for msg := range msgs {
			var m models.NewMessageEvent
			err := json.Unmarshal(msg.Body, &m)
			if err == nil {
				output <- m
			}
		}
	}()
	return output, nil
}

func publishMessage(ctx context.Context, ch *amqp.Channel, queue string, data interface{}) error {
	datum, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ch.PublishWithContext(ctx, queue, "", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(datum),
	})
}
