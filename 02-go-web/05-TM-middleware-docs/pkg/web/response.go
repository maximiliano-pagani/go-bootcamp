package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Data any `json:"data"`
}

func Success(c *gin.Context, status int, data any) {
	c.JSON(status, SuccessResponse{Data: data})
}

func Failure(c *gin.Context, status int, err error) {
	var msg string

	if err != nil {
		msg = err.Error()
	}

	c.AbortWithStatusJSON(
		status,
		ErrorResponse{
			Status:  status,
			Code:    http.StatusText(status),
			Message: msg,
		},
	)

}
