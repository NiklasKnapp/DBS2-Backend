package models

import (
	"database/sql"

	"de.stuttgart.hft/DBS2-Backend/pkg/config"
)

var db *sql.DB

// Initialize response
func init() {
	db = config.GetDB()
}

// Message
type Message struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

// Response template
type Response[T any] struct {
	Success  bool      `json:"success"`
	Errors   []Message `json:"errors"`
	Messages []Message `json:"messages"`
	Result   T         `json:"result"`
}
