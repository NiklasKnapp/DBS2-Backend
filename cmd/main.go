package main

import (
	"net/http"

	"de.stuttgart.hft/DBS2-Backend/pkg/routes"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	//"database/sql"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(utils.RequestRecovery))

	routes.RegisterManufacturerRoutes(r)
	http.Handle("/", r)
	r.Run(":8080")
}
