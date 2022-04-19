package main

import (
	"de.stuttgart.hft/DBS2-Backend/handler"
	"github.com/gin-gonic/gin"
	//"database/sql"
	//_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.Run(":8080")
	//db, err := sql.Open("sqlite3", "./database.db")
}
