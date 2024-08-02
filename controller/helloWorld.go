package controller

import (
	"first-app/models/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(r *gin.Engine) {
	r.GET("/hello", func(ctx *gin.Context) {
		responseData := response.Response{
			Result:  "OK",
			Error:   false,
			Message: "Get Data",
			Data: map[string]string{
				"message": "hello world using pattern",
			},
		}
		ctx.JSON(http.StatusOK, responseData)
	})
}
