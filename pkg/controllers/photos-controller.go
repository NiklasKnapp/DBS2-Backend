package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"de.stuttgart.hft/DBS2-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
 // Create photo
func CreatePhoto(c *gin.Context) {

	// Upload new photo
	newPhoto := &models.PhotoUpload{}

	// Bind values to photo
	if err := c.ShouldBind(newPhoto); err != nil {
		log.Println("[FORM PARSING]: CreatePhoto: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}

	//Save Files to server
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("CreatePhoto: MultipartForm could not be instantiated")
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	var collection = models.PhotoUploadResponse{}
	formFiles, _ := form.File["files"]
	for _, file := range formFiles {
		newFileUpload := &models.PhotoUpload{}
		extention := filepath.Ext(file.Filename)
		newFileName := uuid.New().String() + extention
		filename := "../pkg/tmp/" + newFileName
		newFileUpload.UUID = newFileName
		newFileUpload.Roll_id = newPhoto.Roll_id
		newFileUpload.Files = append(newFileUpload.Files, file)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			fmt.Println("CreatePhoto: Unable to upload Photo(s)")
			utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
			return
		}
		log.Println("CreatePhoto: File Uploaded: ", newFileName)

		// Create photo in DB
		p, err := newFileUpload.CreatePhoto()
		if err != nil {
			log.Println("[SQL]: ", err)
			utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
			return
		}
		// collection = append(collection, p...)
		collection.PhotoUpload = append(collection.PhotoUpload, *p)
	}
	utils.ApiSuccess(c, [][]string{}, collection, 200)
}

// Get photos
func GetPhoto(c *gin.Context) {

	// Get photos from DB
	photo, err := models.GetPhoto()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photo, 200)
}

// Get photos by type ID
func GetPhotoByTypeId(c *gin.Context) {

	// Get type ID from request
	typeIdParams := c.Params.ByName("type_id")

	// Parse ID
	typeId, err := strconv.ParseInt(typeIdParams, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetPhotoByTypeId: Could not parse type id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get photos from DB
	photo, err := models.GetPhotoByTypeId(typeId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photo, 200)
}

//Get photo by ID
func GetPhotoById(c *gin.Context) {

	// Get photo ID from request
	photoIdParams := c.Params.ByName("photo_id")

	// Parse ID
	photoId, err := strconv.ParseInt(photoIdParams, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetPhotoById: Could not parse photo id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get photo from DB
	photo, err := models.GetPhotoById(photoId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photo, 200)
}

// Get data from photo
func GetPhotoData(c *gin.Context) {

	// Get UUID from requests
	uuid := c.Params.ByName("uuid")

	// Get file data
	photoData, _ := ioutil.ReadFile("../pkg/tmp/" + uuid)
	mimeType := http.DetectContentType(photoData)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+uuid)
	switch mimeType {
	case "image/jpeg":
		c.Header("Content-Type", "image/jpeg")
	case "image/png":
		c.Header("Content-Type", "image/png")
	}

	c.Writer.Write(photoData)

	// photoData, _ := ioutil.ReadFile("../pkg/tmp/" + uuid)
	// var base64Encoding string
	// mimeType := http.DetectContentType(photoData)
	// switch mimeType {
	// case "image/jpeg":
	// 	base64Encoding += "data:image/jpeg;base64,"
	// case "image/png":
	// 	base64Encoding += "data:image/png;base64,"
	// }
	// base64Encoding += base64.StdEncoding.EncodeToString(photoData)
	// c.Writer.Write(photoData)

}

// Get photos by roll ID
func GetPhotosByRollId(c *gin.Context) {

	// Get roll ID from request
	rollIdParams := c.Params.ByName("roll_id")

	// Parse ID
	rollId, err := strconv.ParseInt(rollIdParams, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetPhotosByRollId: Could not parse roll id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get photos from DB
	photos, err := models.GetPhotosByRollId(rollId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photos, 200)
}

// Get photos by album ID
func GetPhotosByAlbumId(c *gin.Context) {

	// Get album ID from request
	albumIdParams := c.Params.ByName("album_id")

	// Parse ID
	albumId, err := strconv.ParseInt(albumIdParams, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetPhotosByAlbumId: Could not parse album id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get photos from DB
	photos, err := models.GetPhotosByAlbumId(albumId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photos, 200)
}

// Update photo
func UpdatePhoto(c *gin.Context) {

	// Initialize new photo
	updatedPhoto := &models.Photo{}

	// Bind values to new photo
	if err := c.ShouldBindJSON(updatedPhoto); err != nil {
		log.Println("[JSON PARSING]: UpdatePhoto: Could not map required fields")
		utils.ApiError(c, [][]string{{"bad.request", utils.GetEnvVar("ERROR_CODE_BODY_INVALID")}}, 400)
		return
	}

	// Get ID from request
	photoIdParam := c.Params.ByName("photo_id")

	// Parse ID
	photoId, err := strconv.ParseInt(photoIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: UpdatePhoto: Could not parse Photo id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get current photo from DB
	currentPhoto, err := models.GetPhotoById(photoId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Detect updated values
	if updatedPhoto.Title != "" {
		currentPhoto.Title = updatedPhoto.Title
	}
	if updatedPhoto.UUID != "" {
		currentPhoto.UUID = updatedPhoto.UUID
	}
	if updatedPhoto.Rating != 0 {
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

	// Update photo in DB
	p, _ := currentPhoto.UpdatePhoto()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, p, 200)
}

// Delete photo
func DeletePhoto(c *gin.Context) {

	// Get photo ID from request
	photoIdParam := c.Params.ByName("photo_id")

	// Parse ID
	photoId, err := strconv.ParseInt(photoIdParam, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: DeletePhoto: Could not parse filmRoll id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Get photo from DB
	photo, err := models.GetPhotoById(photoId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}

	// Remove file
	errr := os.Remove("../pkg/tmp/" + photo.UUID)
	if errr != nil {
		fmt.Println("DeletePhoto: Could not delete Photo from Server")
		return
	}

	// Delete photo from DB
	photo, er := models.DeletePhoto(photoId)
	if er != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photo, 200)

}
