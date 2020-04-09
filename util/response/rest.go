package response

import (
	"github.com/gin-gonic/gin"
)

type Response = map[string]interface{}

func Success(ResponseData Response, context *gin.Context) {
	response := make(Response)

	response["code"] = 200
	response["msg"] = "ok"
	response["data"] = ResponseData

	context.JSON(200, response)
}

func Error(code int, error error, context *gin.Context) {
	response := make(Response)

	response["code"] = code
	response["msg"] = error.Error()

	context.JSON(200, response)
}
