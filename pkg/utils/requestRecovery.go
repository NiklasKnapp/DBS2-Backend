package utils

import "github.com/gin-gonic/gin"

func RequestRecovery(c *gin.Context, recovered interface{}) {
	ApiError(c, [][]string{{"general.error", GetEnvVar("ERROR_CODE_NO_ROUTE")}}, 404)
}
