package api

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/fcsgehrke/finchat/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handler) UserCreate(ctx *fiber.Ctx) error {
	var req models.UserCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	resp, err := h.usersService.CreateUser(ctx.Context(), &req)
	if err != nil {
		return err
	}

	output, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	return ctx.SendString(string(output))
}

func (h *Handler) UserLogin(ctx *fiber.Ctx) error {
	user := ctx.FormValue("user")
	pass := ctx.FormValue("pass")

	resp, err := h.usersService.Login(ctx.Context(), &models.UserLoginRequest{
		Email:    user,
		Password: pass,
	})

	if err != nil {
		return err
	}

	output, err := json.Marshal(resp)

	if err != nil {
		return err
	}

	return ctx.SendString(string(output))
}

func (h *Handler) RoomCreate(ctx *fiber.Ctx) error {
	var req models.CreateRoomRequest

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	resp, err := h.roomsService.CreateRoom(ctx.Context(), &req)
	if err != nil {
		return err
	}

	output, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	return ctx.SendString(string(output))
}

func (h *Handler) RoomList(ctx *fiber.Ctx) error {
	resp, err := h.roomsService.ListRooms(ctx.Context())
	if err != nil {
		return err
	}

	output, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	return ctx.SendString(string(output))
}

func (h *Handler) RoomConnect(conn *websocket.Conn) {
	var msgs chan models.NewMessageEvent

	defer func() {
		conn.Close()
	}()

	ctx := context.Background()

	roomId, err := strconv.Atoi(conn.Params("id"))
	if err != nil {
		return
	}

	// token := conn.Locals("user").(*jwt.Token)
	// userEmail := getUserEmailFromToken(token)

	// err = h.roomsService.SendMessage(ctx, roomId, userEmail, "", true)
	// if err != nil {
	// 	return
	// }

	msgs, err = h.roomsService.GetMessages(ctx, roomId)
	if err != nil {
		return
	}

	go func() {
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				h.quit <- true
				break
			}
		}
	}()

	for {
		select {
		case msg := <-msgs:
			conn.WriteJSON(msg)
		case <-h.quit:
			return
		}
	}
}

func (h *Handler) RoomSendMessage(ctx *fiber.Ctx) error {
	var req models.SendMessageRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	roomId, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}

	token := ctx.Locals("user").(*jwt.Token)
	userEmail := getUserEmailFromToken(token)

	// TODO: Add user
	err = h.roomsService.SendMessage(ctx.Context(), roomId, userEmail, req.Content, false)
	if err != nil {
		return err
	}
	return nil
}

func getUserEmailFromToken(token *jwt.Token) string {
	claims := token.Claims.(jwt.MapClaims)
	return claims["user"].(string)
}
