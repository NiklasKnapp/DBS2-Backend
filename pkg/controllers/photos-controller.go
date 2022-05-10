package controllers

import (
	"log"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	newPhoto := &models.Photo{}
	if err := c.ShouldBindJSON(newPhoto); err != nil {
		log.Println("[JSON PARSING]: CreatePhoto: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	p, err := newPhoto.CreatePhoto()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, p, 200)
}

func GetPhoto(c *gin.Context) {
	photo, err := models.GetPhoto()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photo, 200)
}

func GetPhotoById(c *gin.Context) {
	photoIdParams := c.Params.ByName("photo_id")
	photoId, err := strconv.ParseInt(photoIdParams, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetPhotoById: Could not parse photo id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	photo, err := models.GetPhotoById(photoId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photo, 200)

}

func UpdatePhoto(c *gin.Context) {
	updatedPhoto := &models.Photo{}
	if err := c.ShouldBindJSON(updatedPhoto); err != nil {
		log.Println("[JSON PARSING]: UpdatePhoto: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	photoIdParam := c.Params.ByName("photo_id")
	photoId, err := strconv.ParseInt(photoIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: UpdatePhoto: Could not parse Photo id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	currentPhoto, err := models.GetPhotoById(photoId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	if updatedPhoto.Title != "" {
		currentPhoto.Title = updatedPhoto.Title
	}
	if updatedPhoto.UUID != "" {
		currentPhoto.UUID = updatedPhoto.UUID
	}
	if updatedPhoto.Roll_id != 0 {
		roll_id, _ := models.GetFilmRollById(int64(updatedPhoto.Roll_id))
		if roll_id == nil {
			log.Printf("UpdatePhoto: FilmRoll with roll_id %v does not exist", updatedPhoto.Roll_id)
			utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
			return
		}
		currentPhoto.Roll_id = updatedPhoto.Roll_id
	}
	p, _ := currentPhoto.UpdatePhoto()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, p, 200)
}

func DeletePhoto(c *gin.Context) {
	photoIdParam := c.Params.ByName("photo_id")
	photoId, err := strconv.ParseInt(photoIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: DeletePhoto: Could not parse filmRoll id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	photo, err := models.DeletePhoto(photoId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photo, 200)

}
