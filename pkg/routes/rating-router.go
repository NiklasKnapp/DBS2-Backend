package routes

import (
	"de.stuttgart.hft/DBS2-Backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterFilmRollRoutes = func(router *gin.Engine) {
	router.POST("/rating/", controllers.CreateRating)
	router.GET("/rating/", controllers.GetRating)
	router.GET("/rating/:ratingId", controllers.GetRatingById)
	router.PUT("/rating/:ratingId", controllers.UpdateRating)
	router.DELETE("/rating/:ratingId", controllers.DeleteRating)
}
