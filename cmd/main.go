package main

import (
	"net/http"

	"de.stuttgart.hft/DBS2-Backend/pkg/config"
	"de.stuttgart.hft/DBS2-Backend/pkg/routes"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	//"database/sql"
	// _ "github.com/mattn/go-sqlite3"
)

//Main
func main() {

	//Initialize Gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(utils.RequestRecovery))

	config.Connect()

	//Register Routes
	routes.RegisterManufacturerRoutes(r)
	routes.RegisterRollTypeRoutes(r)
	routes.RegisterFilmRollRoutes(r)
	routes.RegisterPhotoRoutes(r)
	routes.RegisterAlbumRoutes(r)
	routes.RegisterRatingRoutes(r)

	//Run
	http.Handle("/", r)
	r.Run(":8080")
}
