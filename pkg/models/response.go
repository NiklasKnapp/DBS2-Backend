package models

import (
	"database/sql"

	"de.stuttgart.hft/DBS2-Backend/pkg/config"
)

var db *sql.DB

func init() {
	db = config.GetDB()
}

type Message struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

type Response[T any] struct {
	Success  bool      `json:"success"`
	Errors   []Message `json:"errors"`
	Messages []Message `json:"messages"`
	Result   T         `json:"result"`
}
