package controllers

import (
	"log"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CreateRating(c *gin.Context) {
	newRating := &models.Rating{}
	if err := c.ShouldBindJSON(newRating); err != nil {
		log.Println("[JSON PARSING]: CreateRating: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	utils.ApiSuccess(c, [][]string{}, newRating, 200)
}

func GetRating(c *gin.Context) {
	rating, err := models.GetRating()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rating, 200)
}

func GetRatingById(c *gin.Context) {
	ratingIdParam := c.Params.ByName("ratingId")
	ratingId, err := strconv.ParseInt(ratingIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetRatingById: Could not parse rating id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	rating, err := models.GetRatingById(ratingId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rating, 200)
}

func UpdateRating(c *gin.Context) {
	updatedRating := &models.Rating{}
	if err := c.ShouldBindJSON(updatedRating); err != nil {
		log.Println("[JSON PARSING]: UpdateRating: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	ratingIdParam := c.Params.ByName("ratingId")
	ratingId, err := strconv.ParseInt(ratingIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: UpdateRating: Could not parse rating id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	currentRating, err := models.GetRatingById(ratingId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	if updatedRating.rating != 0 {
		currentRating.rating = updatedRating.rating
	}
	if updatedRating.Photo_id != 0 {
		photo_id, _ := models.GetPhotoById(int64(updatedRating.Photo_id))
		if photo_id == nil {
			log.Printf("UpdateRating: Photo with photo_id %v does not exist", updatedRating.Photo_id)
			utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
			return
		}
		currentRating.Photo_id = updatedRating.Photo_id
	}
	fr, _ := currentRating.UpdateRating()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, fr, 200)
}

func DeleteRating(c *gin.Context) {
	ratingIdParam := c.Params.ByName("ratingId")
	ratingId, err := strconv.ParseInt(ratingIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: DeleteRating: Could not parse rating id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	rating, err := models.DeleteRating(ratingId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rating, 200)
}
