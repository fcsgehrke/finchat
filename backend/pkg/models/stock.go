package models

type StockPriceRequest struct {
	StockCode string `json:"stock_code"`
	RoomId    string `json:"room_id"`
}

type StockPriceResponse struct {
	StockCode string  `json:"stock_code"`
	RoomId    string  `json:"room_id"`
	Price     float32 `json:"price"`
	Error     string  `json:"error"`
}
