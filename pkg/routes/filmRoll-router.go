package routes

import (
	"de.stuttgart.hft/DBS2-Backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterFilmRollRoutes = func(router *gin.Engine) {
	router.POST("/filmroll/", controllers.CreateFilmRoll)
	router.GET("/filmroll/", controllers.GetFilmRoll)
	router.GET("/filmroll/:rollId", controllers.GetFilmRollById)
	router.PUT("/filmroll/:rollId", controllers.UpdateFilmRoll)
	router.DELETE("/filmroll/:rollId", controllers.DeleteFilmRoll)
}
