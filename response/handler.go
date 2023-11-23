package response

import "github.com/gin-gonic/gin"

func SetResponse(ctx *gin.Context, success bool, statusCode int, body []map[string]interface{}, err error) {
	if success {
		ctx.JSON(statusCode, body)
		return
	}

	ctx.JSON(statusCode, map[string]interface{}{"error": err.Error()})
}
