package controllers

import (
	"log"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Create album
func CreateAlbum(c *gin.Context) {

	//Initialize new album
	newAlbum := &models.Album{}

	//Bind values from request to album
	if err := c.ShouldBindJSON(newAlbum); err != nil {
		log.Println("[JSON PARSING]: CreateAlbum: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}

	// Create album in DB
	a, err := newAlbum.CreateAlbum()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	
	utils.ApiSuccess(c, [][]string{}, a, 200)
}

// Get albums
func GetAlbum(c *gin.Context) {

	// Get albums from DB
	album, err := models.GetAlbum()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, album, 200)
}

// Get album by ID
func GetAlbumById(c *gin.Context) {

	// Read ID from request
	albumIdParam := c.Params.ByName("albumId")

	// Parse ID
	albumId, err := strconv.ParseInt(albumIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetAlbumById: Could not parse album id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get album from DB
	album, err := models.GetAlbumById2(albumId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, album, 200)
}

// Update album
func UpdateAlbum(c *gin.Context) {

	// Initialize new album
	updatedAlbum := &models.Album{}

	// Bind new values from request to new album
	if err := c.ShouldBindJSON(updatedAlbum); err != nil {
		log.Println("[JSON PARSING]: UpdateAlbum: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}

	// Get ID from request
	albumIdParam := c.Params.ByName("albumId")

	// Parse ID
	albumId, err := strconv.ParseInt(albumIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: UpdateAlbum: Could not parse album id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get current album from DB
	currentAlbum, err := models.GetAlbumById2(albumId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Detect new values
	if updatedAlbum.Title != "" {
		currentAlbum.Title = updatedAlbum.Title
	}
	if updatedAlbum.Rating != 0 {
		currentAlbum.Rating = updatedAlbum.Rating
	}
	if updatedAlbum.Description != "" {
		currentAlbum.Description = updatedAlbum.Description
	}

	// Update album in DB
	fr, _ := currentAlbum.UpdateAlbum()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, fr, 200)
}

// Delete album
func DeleteAlbum(c *gin.Context) {

	// Get ID from request
	albumIdParam := c.Params.ByName("albumId")

	// Parse ID
	albumId, err := strconv.ParseInt(albumIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: DeleteAlbum: Could not parse album id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Delete album in DB
	album, err := models.DeleteAlbum(albumId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, album, 200)
}

// Add photo to album
func Photos_Album(c *gin.Context) {

	// Initialize new link
	newpa := &models.PA{}

	// Bind values to link
	if err := c.ShouldBindJSON(newpa); err != nil {
		log.Println("[JSON PARSING]: Photos_Album: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}

	// Create link in DB
	pa, err := newpa.CreatePA()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, pa, 200)
}