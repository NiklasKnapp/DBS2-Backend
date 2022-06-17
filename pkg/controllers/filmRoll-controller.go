package controllers

import (
	"log"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CreateFilmRoll(c *gin.Context) {
	newFilmRoll := &models.FilmRoll{}
	if err := c.ShouldBindJSON(newFilmRoll); err != nil {
		log.Println("[JSON PARSING]: CreateFilmRoll: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	
	fr, err := newFilmRoll.CreateRollType()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, fr, 200)
}

func GetFilmRoll(c *gin.Context) {
	filmRoll, err := models.GetFilmRoll()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, filmRoll, 200)
}

func GetFilmRollById(c *gin.Context) {
	rollIdParam := c.Params.ByName("rollId")
	rollId, err := strconv.ParseInt(rollIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetFilmRollById: Could not parse filmRoll id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	roll, err := models.GetFilmRollById(rollId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, roll, 200)
}

func UpdateFilmRoll(c *gin.Context) {
	updatedFilmRoll := &models.FilmRoll{}
	if err := c.ShouldBindJSON(updatedFilmRoll); err != nil {
		log.Println("[JSON PARSING]: UpdateFilmRoll: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	rollIdParam := c.Params.ByName("rollId")
	rollId, err := strconv.ParseInt(rollIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: UpdateFilmRoll: Could not parse filmRoll id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	currentFilmRoll, err := models.GetFilmRollById(rollId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	if updatedFilmRoll.Title != "" {
		currentFilmRoll.Title = updatedFilmRoll.Title
	}
	if updatedFilmRoll.Rating != 0 {
		currentFilmRoll.Rating = updatedFilmRoll.Rating
	}
	if updatedFilmRoll.Description != "" {
		currentFilmRoll.Description = updatedFilmRoll.Description
	}
	if updatedFilmRoll.Type_id != 0 {
		type_id, _ := models.GetRollTypeById(int64(updatedFilmRoll.Type_id))
		if type_id == nil {
			log.Printf("UpdateFilmRoll: RollType with type_id %v does not exist", updatedFilmRoll.Type_id)
			utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
			return
		}
		currentFilmRoll.Type_id = updatedFilmRoll.Type_id
	}
	fr, _ := currentFilmRoll.UpdateFilmRoll()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, fr, 200)
}

func DeleteFilmRoll(c *gin.Context) {
	rollIdParam := c.Params.ByName("rollId")
	rollId, err := strconv.ParseInt(rollIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: DeleteFilmRoll: Could not parse filmRoll id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	roll, err := models.DeleteFilmRoll(rollId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, roll, 200)
}
