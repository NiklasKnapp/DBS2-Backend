package controllers

import (
	"log"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Create rating
func CreateRating(c *gin.Context) {

	// Initialize new rating
	newRating := &models.RatingRaw{}

	// Bind values to rating
	if err := c.ShouldBindJSON(newRating); err != nil {
		log.Println("[JSON PARSING]: CreateRating: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	
	// Create rating in DB
	fr, err := newRating.CreateRating()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, fr, 200)
}

// Get ratings
func GetRating(c *gin.Context) {

	// Get ratings from DB
	rating, err := models.GetRating()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rating, 200)
}

// Get rating by ID
func GetRatingById(c *gin.Context) {

	// Get ID from request
	ratingIdParam := c.Params.ByName("ratingId")

	// Parse ID
	ratingId, err := strconv.ParseInt(ratingIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetRatingById: Could not parse rating id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get rating from DB
	rating, err := models.GetRatingById(ratingId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rating, 200)
}

// Update rating
func UpdateRating(c *gin.Context) {

	// Initialize new rating
	updatedRating := &models.Rating{}

	// Bind values to new rating
	if err := c.ShouldBindJSON(updatedRating); err != nil {
		log.Println("[JSON PARSING]: UpdateRating: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}

	// Get ID from request
	ratingIdParam := c.Params.ByName("ratingId")

	// Parse ID
	ratingId, err := strconv.ParseInt(ratingIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: UpdateRating: Could not parse rating id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get current rating from DB
	currentRating, err := models.GetRatingById(ratingId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Detect updates values
	if updatedRating.Rating != 0 {
		currentRating.Rating = updatedRating.Rating
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

	// Update rating in DB
	fr, _ := currentRating.UpdateRating()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, fr, 200)
}

// Delete rating
func DeleteRating(c *gin.Context) {

	// Get ID from request
	ratingIdParam := c.Params.ByName("ratingId")

	// Parse ID
	ratingId, err := strconv.ParseInt(ratingIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: DeleteRating: Could not parse rating id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Delete rating in DB
	rating, err := models.DeleteRating(ratingId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, rating, 200)
}
