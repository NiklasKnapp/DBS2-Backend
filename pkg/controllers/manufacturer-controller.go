package controllers

import (
	"log"
	"net/http"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"github.com/gin-gonic/gin"
)

var NewManufacturer models.Manufacturer

func CreateManufacturer(c *gin.Context) {
	CreateManufacturer := &models.Manufacturer{}
	if err := c.ShouldBindJSON(CreateManufacturer); err != nil {
		c.AbortWithError(http.StatusBadRequest, err) //TODO: Change err response to json body
		return
	}
	m, err := CreateManufacturer.CreateManufacturer()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": nil, //response is null
			"error":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, m)
}

func GetManufacturer(c *gin.Context) {
	newManufacturers, err := models.GetManufacturer()
	if err != nil {
		log.Printf(">>ERROR<< %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"response": newManufacturers, //response is null
			"error":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": newManufacturers, //response is empty struct []
		"error":    nil,
	})
}

func GetManufacturerById(c *gin.Context) {
	mIdParam := c.Params.ByName("mId")
	MId, err := strconv.ParseInt(mIdParam, 0, 0)
	if err != nil {
		log.Printf(">>ERROR<< %v\n", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	manufacturer, err := models.GetManufacturerById(MId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": nil, //response is null
			"error":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": manufacturer, //response is null
		"error":    nil,
	})
}

func UpdateManufacturer(c *gin.Context) {
	UpdateManufacturer := &models.Manufacturer{}
	if err := c.ShouldBindJSON(UpdateManufacturer); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	mIdParam := c.Params.ByName("mId")
	MId, err := strconv.ParseInt(mIdParam, 0, 0)
	if err != nil {
		log.Printf(">>ERROR<< %v\n", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	currentManufacturer, err := models.GetManufacturerById(MId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": nil, //response is null
			"error":    err.Error(),
		})
		return
	}
	if UpdateManufacturer.Name != "" {
		currentManufacturer.Name = UpdateManufacturer.Name
	}
	//m, err := UpdateManufacturer.UpdateManufacturer()
	m, _ := currentManufacturer.UpdateManufacturer()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": nil, //response is null
			"error":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": m, //response is null
		"error":    nil,
	})
}

func DeleteManufacturer(c *gin.Context) {
	mIdParam := c.Params.ByName("mId")
	MId, err := strconv.ParseInt(mIdParam, 0, 0)
	if err != nil {
		log.Printf(">>ERROR<< %v\n", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	manufacturer, err := models.DeleteManufacturer(MId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"response": nil, //response is null
			"error":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": manufacturer, //response is null
		"error":    nil,
	})
}
