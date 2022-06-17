package routes

import (
	"de.stuttgart.hft/DBS2-Backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterAlbumRoutes = func(router *gin.Engine) {
	router.POST("/album/", controllers.CreateAlbum)
	router.GET("/album/", controllers.GetAlbum)
	router.GET("/album/:albumId", controllers.GetAlbumById)
	router.PUT("/album/:albumId", controllers.UpdateAlbum)
	router.DELETE("/album/:albId", controllers.DeleteAlbum)
	router.POST("/photos_album/", controllers.Photos_Album)
}
