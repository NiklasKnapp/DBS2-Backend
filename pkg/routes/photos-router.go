package routes

import (
	"de.stuttgart.hft/DBS2-Backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterPhotoRoutes = func(router *gin.Engine) {
	router.POST("/photo/", controllers.CreatePhoto)
	router.GET("/photo/", controllers.GetPhoto)
	router.GET("/photo/:photo_id", controllers.GetPhotoById)
	router.GET("/photo/roll/:roll_id", controllers.GetPhotosByRollId)
	router.PUT("/photo/:photo_id", controllers.UpdatePhoto)
	router.DELETE("/photo/:photo_id", controllers.DeletePhoto)
}
