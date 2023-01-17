package api

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/websocket/v2"
)

func ConfigRoutes(app *fiber.App, handler *Handler, jwtSecret string) {
	// API Group
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Users
	users := v1.Group("/users")
	users.Post("/", handler.UserCreate)
	users.Post("/login", handler.UserLogin)

	// Rooms
	rooms := v1.Group("/rooms")
	rooms.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte(jwtSecret),
		ErrorHandler: jwtError,
	}))
	rooms.Post("/", handler.RoomCreate)
	rooms.Get("/", handler.RoomList)
	rooms.Get("/:id/ws", websocket.New(handler.RoomConnect))
	rooms.Post("/:id/message", handler.RoomSendMessage)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}
