package bot

import (
	"context"
	"encoding/json"

	"github.com/fcsgehrke/finchat/pkg/models"
	"github.com/fcsgehrke/finchat/pkg/stooq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Bot struct {
	amqpAddr  string
	reqQueue  string
	respQueue string
	ch        *amqp.Channel
}

func NewBot(amqpAddr string, reqQueue string, respQueue string) (*Bot, error) {
	conn, err := amqp.Dial(amqpAddr)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		reqQueue, // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		respQueue, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return &Bot{
		ch:        ch,
		amqpAddr:  amqpAddr,
		reqQueue:  reqQueue,
		respQueue: respQueue,
	}, nil
}

func (b *Bot) Start() error {
	msgs, err := b.ch.Consume(
		b.reqQueue, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	// go func() {
	for msg := range msgs {
		var req models.StockPriceRequest

		err := json.Unmarshal(msg.Body, &req)
		if err != nil {
			b.sendMessage(&models.StockPriceResponse{
				Error: err.Error(),
			})
		}

		price, err := stooq.GetStockPrice(req.StockCode)
		if err != nil {
			b.sendMessage(&models.StockPriceResponse{
				Error: err.Error(),
			})
		}

		b.sendMessage(&models.StockPriceResponse{
			RoomId:    req.RoomId,
			Price:     price,
			StockCode: req.StockCode,
		})
	}
	// }()

	return nil
}

func (b *Bot) sendMessage(resp *models.StockPriceResponse) error {
	content, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	return b.ch.PublishWithContext(
		context.Background(),
		"",
		b.respQueue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        content,
		})
}
