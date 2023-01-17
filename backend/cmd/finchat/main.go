package main

import (
	"log"
	"time"

	"github.com/fcsgehrke/finchat/api"
	"github.com/fcsgehrke/finchat/internal/db/postgres"
	"github.com/fcsgehrke/finchat/internal/services"
	"github.com/fcsgehrke/finchat/pkg/crypt"
	"github.com/fcsgehrke/finchat/pkg/db"
	"github.com/fcsgehrke/finchat/pkg/rooms"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database, err := db.Connect("postgresql://user:pass@localhost:5432/finchat?sslmode=disable")
	if err != nil {
		log.Fatalf("[ERR] - Couldn't connect DB w/ err: %s", err.Error())
	}

	repository, err := postgres.NewRepository(database)
	if err != nil {
		log.Fatalf("[ERR] - Couldn't setup Repository w/ err: %s", err.Error())
	}

	repository.RunMigrations()

	crypt, err := crypt.NewCrypter("secret", 72*time.Hour)
	if err != nil {
		log.Fatalf("[ERR] - Couldn't start Crypt service w/ err: %s", err.Error())
	}

	roomManager, err := rooms.NewRoomsManager("amqp://guest:guest@localhost:5672/", repository)
	if err != nil {
		log.Fatalf("[ERR] - Couldn't start Rooms Manager service w/ err: %s", err.Error())
	}

	service, err := services.NewService(repository, crypt, roomManager)
	if err != nil {
		log.Fatalf("[ERR] - Couldn't start FinChat service w/ err: %s", err.Error())
	}

	quit := make(chan bool)
	handler, err := api.NewAPIHandler(service, service, quit)
	if err != nil {
		log.Fatalf("[ERR] - Couldn't start API handler w/ err: %s", err.Error())
	}

	// price, err := stooq.GetStockPrice("aapl.us")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("Price: %f\n", price)

	app := fiber.New()
	api.ConfigRoutes(app, handler, "secret")
	log.Fatal(app.Listen(":5000"))
}
