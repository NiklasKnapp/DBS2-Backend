package controllers

import (
	"log"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CreateAlbum(c *gin.Context) {
	newAlbum := &models.Album{}
	if err := c.ShouldBindJSON(newAlbum); err != nil {
		log.Println("[JSON PARSING]: CreateAlbum: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	
	utils.ApiSuccess(c, [][]string{}, newAlbum, 200)
}

func GetAlbum(c *gin.Context) {
	album, err := models.GetAlbum()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, album, 200)
}

func GetAlbumById(c *gin.Context) {
	albumIdParam := c.Params.ByName("albumId")
	albumId, err := strconv.ParseInt(albumIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetAlbumById: Could not parse album id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	album, err := models.GetAlbumById(albumId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, album, 200)
}

func UpdateAlbum(c *gin.Context) {
	updatedAlbum := &models.Album{}
	if err := c.ShouldBindJSON(updatedAlbum); err != nil {
		log.Println("[JSON PARSING]: UpdateAlbum: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	albumIdParam := c.Params.ByName("albumId")
	albumId, err := strconv.ParseInt(albumIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: UpdateAlbum: Could not parse album id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	currentAlbum, err := models.GetAlbumById(albumId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	if updatedAlbum.Title != "" {
		currentAlbum.Title = updatedAlbum.Title
	}
	if updatedAlbum.Rating != 0 {
		currentAlbum.Rating = updatedAlbum.Rating
	}
	if updatedAlbum.Description != "" {
		currentAlbum.Description = updatedAlbum.Description
	}
	fr, _ := currentAlbum.UpdateAlbum()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, fr, 200)
}

func DeleteAlbum(c *gin.Context) {
	albumIdParam := c.Params.ByName("albumId")
	albumId, err := strconv.ParseInt(albumIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: DeleteAlbum: Could not parse album id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	album, err := models.DeleteAlbum(albumId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, album, 200)
}

func Photos_Album(c *gin.Context) {
	newpa := &models.PA{}
	if err := c.ShouldBindJSON(newpa); err != nil {
		log.Println("[JSON PARSING]: CreateAlbum: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}
	pa, err := newpa.CreatePA()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, pa, 200)
}