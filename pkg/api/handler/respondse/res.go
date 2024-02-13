package res

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int         `json:"stastuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"error"`
}

func ErrrorResponse(ctx *gin.Context, Code int, message string, Errors interface{}) {
	response := Response{
		StatusCode: Code,
		Message:    message,
		Data:       nil,
		Errors:     Errors,
	}
	ctx.JSON(response.StatusCode, response)
}

func SuccessResponse(ctx *gin.Context, Code int, message string, data interface{}) {
	response := Response{
		StatusCode: Code,
		Message:    message,
		Data:       data,
		Errors:     nil,
	}
	ctx.JSON(response.StatusCode, response)
}