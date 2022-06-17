package controllers

import (
	"log"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Create roll type
func CreateRollType(c *gin.Context) {

	// Initialize new roll type
	newRollType := &models.RollType{}

	// Bind values to roll type
	if err := c.ShouldBindJSON(newRollType); err != nil {
		log.Println("[JSON PARSING]: CreateRollType: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}

	// Create roll type in DB
	rt, err := newRollType.CreateRollType()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rt, 200)
}

// Get roll types
func GetRollType(c *gin.Context) {

	// Get roll types from DB
	rollTypes, err := models.GetRollType()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rollTypes, 200)
}

// Get Roll type by ID
func GetRollTypeById(c *gin.Context) {

	// Get ID from request
	typeIdParam := c.Params.ByName("typeId")

	// Parse ID
	typeId, err := strconv.ParseInt(typeIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetRollTypeById: Could not parse rollType id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get roll type from DB
	rollType, err := models.GetRollTypeById(typeId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rollType, 200)
}

// Update roll type
func UpdateRollType(c *gin.Context) {

	// Initialize new roll type
	updatedRollType := &models.RollType{}

	// Bind values to new roll type
	if err := c.ShouldBindJSON(updatedRollType); err != nil {
		log.Println("[JSON PARSING]: UpdateRollType: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}

	// Get ID from request
	typeIdParam := c.Params.ByName("typeId")

	// Parse ID
	typeId, err := strconv.ParseInt(typeIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: UpdateRollType: Could not parse rollType id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	//Get current roll type from DB
	currentRollType, err := models.GetRollTypeById(typeId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Detect updated values
	if updatedRollType.StockName != "" {
		currentRollType.StockName = updatedRollType.StockName
	}
	if updatedRollType.Format != "" {
		currentRollType.Format = updatedRollType.Format
	}
	if updatedRollType.M_id != 0 {
		m_id, _ := models.GetManufacturerById(int64(updatedRollType.M_id))
		if m_id == nil {
			log.Printf("UpdateRollType: Manufacturer with m_id %v does not exist", updatedRollType.M_id)
			utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
			return
		}
		currentRollType.M_id = updatedRollType.M_id
	}

	// Update roll type in DB
	rt, _ := currentRollType.UpdateRollType()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rt, 200)

}

// Delete roll type
func DeleteRollType(c *gin.Context) {

	//Get ID from request
	typeIdParam := c.Params.ByName("typeId")

	// Parse ID
	typeId, err := strconv.ParseInt(typeIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: DeleteRollType: Could not parse rollType id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Delete roll type from DB
	rollType, err := models.DeleteRollType(typeId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rollType, 200)
}
