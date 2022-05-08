package controllers

import (
	"log"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CreateRollType(c *gin.Context) {
	newRollType := &models.RollType{}
	if err := c.ShouldBindJSON(newRollType); err != nil {
		log.Println("[JSON PARSING]: CreateRollType: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	rt, err := newRollType.CreateRollType()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rt, 200)
}

func GetRollType(c *gin.Context) {
	rollTypes, err := models.GetRollType()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rollTypes, 200)
}

func GetRollTypeById(c *gin.Context) {
	typeIdParam := c.Params.ByName("typeId")
	typeId, err := strconv.ParseInt(typeIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetRollTypeById: Could not parse rollType id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	rollType, err := models.GetRollTypeById(typeId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rollType, 200)
}

func UpdateRollType(c *gin.Context) {
	updatedRollType := &models.RollType{}
	if err := c.ShouldBindJSON(updatedRollType); err != nil {
		log.Println("[JSON PARSING]: UpdateRollType: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	typeIdParam := c.Params.ByName("typeId")
	typeId, err := strconv.ParseInt(typeIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: UpdateRollType: Could not parse rollType id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	currentRollType, err := models.GetRollTypeById(typeId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	if updatedRollType.StockName != "" {
		currentRollType.StockName = updatedRollType.StockName
	}
	rt, _ := currentRollType.UpdateRollType()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rt, 200)

}

func DeleteRollType(c *gin.Context) {
	typeIdParam := c.Params.ByName("typeId")
	typeId, err := strconv.ParseInt(typeIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: DeleteRollType: Could not parse rollType id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	rollType, err := models.DeleteRollType(typeId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rollType, 200)
}
