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

func CreatePhoto(c *gin.Context) {

	fmt.Printf("%#v\n", c.Request)

	newPhoto := &models.PhotoUpload{}
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

func GetPhoto(c *gin.Context) {
	photo, err := models.GetPhoto()
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"general.error", utils.GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photo, 200)
}

func GetPhotoByTypeId(c *gin.Context) {
	typeIdParams := c.Params.ByName("type_id")
	typeId, err := strconv.ParseInt(typeIdParams, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetPhotoByTypeId: Could not parse type id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	photo, err := models.GetPhotoByTypeId(typeId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
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

func GetPhotoData(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
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

func GetPhotosByRollId(c *gin.Context) {
	rollIdParams := c.Params.ByName("roll_id")
	rollId, err := strconv.ParseInt(rollIdParams, 0, 0)
	if err != nil {
		log.Println("[STRCONV]: GetPhotosByRollId: Could not parse roll id: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	photos, err := models.GetPhotosByRollId(rollId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photos, 200)
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
	photo, err := models.GetPhotoById(photoId)
	if err != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	errr := os.Remove("../pkg/tmp/" + photo.UUID)
	if errr != nil {
		fmt.Println("DeletePhoto: Could not delete Photo from Server")
		return
	}
	photo, er := models.DeletePhoto(photoId)
	if er != nil {
		log.Println("[SQL]: ", err)
		utils.ApiError(c, [][]string{{"resource.notFound", utils.GetEnvVar("ERROR_RESOURCE_NOT_FOUND")}}, 404)
		return
	}
	utils.ApiSuccess(c, [][]string{}, photo, 200)

}
