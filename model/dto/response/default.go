package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Default struct {
		Code    int    `json:"code"`
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
	}

	Error struct {
		Default
		Error string `json:"error"`
	}

	Success struct {
		Default
		Data any `json:"data,omitempty"`
	}
)

func NewJSONResponse(gc *gin.Context, statusCode int, message string, data any) {
	if statusCode >= http.StatusBadRequest {
		if message == "" {
			message = "Error"
		}

		gc.JSON(statusCode, Error{
			Default: Default{
				Code:    statusCode,
				Ok:      false,
				Message: message,
			},
			Error: data.(error).Error(),
		})
	}

	if message == "" {
		message = "Success"
	}

	gc.JSON(statusCode, Success{
		Default: Default{
			Code:    statusCode,
			Ok:      false,
			Message: message,
		},
		Data: data,
	})
}
