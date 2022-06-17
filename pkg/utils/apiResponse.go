package utils

import (
	"log"
	"strconv"

	"de.stuttgart.hft/DBS2-Backend/pkg/models"
	"github.com/gin-gonic/gin"
)

// Template for succesfull API response
func ApiSuccess[T any](c *gin.Context, message [][]string, result T, statuscode int) {
	convertedMessage := []models.Message{}

	for _, message := range message {
		code, err := strconv.ParseInt(message[1], 10, 32)
		if err != nil {
			log.Println("[API-ERROR-RESPONSE]: Could not convert UUID to integer: ", err)
			ApiError(c, [][]string{{"general.error", GetEnvVar("ERROR_CODE_SERVER_ERROR")}}, 500)
			return
		}
		convertedMessage = append(convertedMessage, models.Message{
			Code:    uint32(code),
			Message: message[0],
		})
	}
	apiSuccess := models.Response[T]{
		Success:  true,
		Errors:   []models.Message{},
		Messages: convertedMessage,
		Result:   result,
	}
	c.JSON(statuscode, apiSuccess)
}

// API error response
func ApiError(c *gin.Context, errors [][]string, statusCode int) {
	convertedMessage := []models.Message{}

	for _, message := range errors {
		code, err := strconv.ParseUint(message[1], 10, 32)
		if err != nil {
			log.Println("[API-ERROR-RESPONSE]: Error code conversion failed:", err)
		}
		convertedMessage = append(convertedMessage, models.Message{
			Code:    uint32(code),
			Message: message[0],
		})
	}

	apiError := models.Response[interface{}]{
		Success:  false,
		Errors:   convertedMessage,
		Messages: []models.Message{},
		Result:   nil,
	}

	c.AbortWithStatusJSON(statusCode, apiError)
}
