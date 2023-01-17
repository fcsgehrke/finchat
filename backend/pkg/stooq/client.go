package stooq

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	url = "https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv"
)

var (
	ErrStockNotFound = errors.New("ErrStockNotFound")
)

func GetStockPrice(stockCode string) (float32, error) {
	addr := fmt.Sprintf(url, stockCode)
	httpRequest, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		return 0.0, err
	}

	client := http.Client{Timeout: time.Second * 10}
	response, err := client.Do(httpRequest)
	if err != nil {
		return 0.0, err
	}

	if response.StatusCode == http.StatusOK {
		content, err := csv.NewReader(response.Body).ReadAll()
		if err != nil {
			return 0.0, err
		}

		stockPrice := content[1][3]
		price, err := strconv.ParseFloat(stockPrice, 32)
		if err != nil {
			return 0.0, ErrStockNotFound
		}

		return float32(price), nil
	}

	return 0.0, ErrStockNotFound
}
