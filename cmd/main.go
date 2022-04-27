package main

import (
	"net/http"

	"de.stuttgart.hft/DBS2-Backend/pkg/routes"
	"github.com/gin-gonic/gin"
	//"database/sql"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()
	routes.RegisterManufacturerRoutes(r)
	http.Handle("/", r)
	r.Run(":8080")
}
