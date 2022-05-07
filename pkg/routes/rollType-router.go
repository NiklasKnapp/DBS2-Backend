package routes

import (
	"de.stuttgart.hft/DBS2-Backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterRollTypeRoutes = func(router *gin.Engine) {
	router.POST("/rolltype/", controllers.CreateRollType)
	router.GET("/rolltype/", controllers.GetRollType)
	router.GET("/rolltype/:typeId", controllers.GetRollTypeById)
	router.PUT("/rolltype/:typeId", controllers.UpdateRollType)
	router.DELETE("/rolltype/:typeId", controllers.DeleteRollType)
}
