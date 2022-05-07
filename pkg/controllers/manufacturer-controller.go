package controllers

import (
	"log"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

// var NewManufacturer models.Manufacturer

func CreateManufacturer(c *gin.Context) {
	CreateManufacturer := &models.Manufacturer{}
	if err := c.ShouldBindJSON(CreateManufacturer); err != nil {
		log.Println("[JSON PARSING]: CreateManufacturer: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	m, err := CreateManufacturer.CreateManufacturer()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, m, 200)
}

func GetManufacturer(c *gin.Context) {
	newManufacturers, err := models.GetManufacturer()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, newManufacturers, 200)
}

func GetManufacturerById(c *gin.Context) {
	mIdParam := c.Params.ByName("mId")
	MId, err := strconv.ParseInt(mIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetManufacturerById: Could not parse manufacturer id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	manufacturer, err := models.GetManufacturerById(MId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, manufacturer, 200)
}

func UpdateManufacturer(c *gin.Context) {
	UpdateManufacturer := &models.Manufacturer{}
	if err := c.ShouldBindJSON(UpdateManufacturer); err != nil {
		log.Println("[JSON PARSING]: UpdateManufacturer: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	mIdParam := c.Params.ByName("mId")
	MId, err := strconv.ParseInt(mIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: UpdateManufacturer: Could not parse manufacturer id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	currentManufacturer, err := models.GetManufacturerById(MId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	if UpdateManufacturer.Name != "" {
		currentManufacturer.Name = UpdateManufacturer.Name
	}
	m, _ := currentManufacturer.UpdateManufacturer()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, m, 200)
}

func DeleteManufacturer(c *gin.Context) {
	mIdParam := c.Params.ByName("mId")
	MId, err := strconv.ParseInt(mIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: DeleteManufacturer: Could not parse manufacturer id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	manufacturer, err := models.DeleteManufacturer(MId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, manufacturer, 200)
}
