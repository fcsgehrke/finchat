package models

type CreateRoomRequest struct {
	Name string `json:"name"`
}

type CreateRoomResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type SendMessageRequest struct {
	Content string `json:"content"`
}

type RoomListItemResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RoomListResponse struct {
	Rooms []*RoomListItemResponse `json:"rooms"`
}
