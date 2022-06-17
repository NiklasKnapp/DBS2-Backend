package controllers

import (
	"log"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Create manufacturer
func CreateManufacturer(c *gin.Context) {

	// Initialize new manufacturer
	CreateManufacturer := &models.Manufacturer{}

	// Bind values to manufacturer
	if err := c.ShouldBindJSON(CreateManufacturer); err != nil {
		log.Println("[JSON PARSING]: CreateManufacturer: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}

	//Create manufacturer in DB
	m, err := CreateManufacturer.CreateManufacturer()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, m, 200)
}

// Get manufacturers
func GetManufacturer(c *gin.Context) {

	// Get manufacturers from DB
	newManufacturers, err := models.GetManufacturer()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, newManufacturers, 200)
}

// Get manufacturer by ID
func GetManufacturerById(c *gin.Context) {

	// Get ID from request
	mIdParam := c.Params.ByName("mId")

	// Parse ID
	MId, err := strconv.ParseInt(mIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetManufacturerById: Could not parse manufacturer id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get manufacturer from DB
	manufacturer, err := models.GetManufacturerById(MId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, manufacturer, 200)
}

// Update manufacturer
func UpdateManufacturer(c *gin.Context) {

	// Initialize new manufacturer
	UpdateManufacturer := &models.Manufacturer{}

	// Bind values to nw manufacturer
	if err := c.ShouldBindJSON(UpdateManufacturer); err != nil {
		log.Println("[JSON PARSING]: UpdateManufacturer: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}

	// Get ID from request
	mIdParam := c.Params.ByName("mId")

	// Parse ID
	MId, err := strconv.ParseInt(mIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: UpdateManufacturer: Could not parse manufacturer id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get current manufacturer from DB
	currentManufacturer, err := models.GetManufacturerById(MId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Detect updated values
	if UpdateManufacturer.Name != "" {
		currentManufacturer.Name = UpdateManufacturer.Name
	}

	// Update manufacturer in DB
	m, _ := currentManufacturer.UpdateManufacturer()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, m, 200)
}

// Delete manufacturer
func DeleteManufacturer(c *gin.Context) {

	// Get ID from request
	mIdParam := c.Params.ByName("mId")

	// Parse ID
	MId, err := strconv.ParseInt(mIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: DeleteManufacturer: Could not parse manufacturer id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Delete manufacturer in DB
	manufacturer, err := models.DeleteManufacturer(MId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, manufacturer, 200)
}
