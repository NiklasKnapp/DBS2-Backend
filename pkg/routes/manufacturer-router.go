package routes

import (
	"de.stuttgart.hft/DBS2-Backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterManufacturerRoutes = func(router *gin.Engine) {
	router.POST("/manufacturer/", controllers.CreateManufacturer)
	router.GET("/manufacturer/", controllers.GetManufacturer)
	router.GET("/manufacturer/:mId", controllers.GetManufacturerById)
	router.PUT("/manufacturer/:mId", controllers.UpdateManufacturer)
	router.DELETE("/manufacturer/:mId", controllers.DeleteManufacturer)
}
