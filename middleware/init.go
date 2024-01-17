package middleware

import "github.com/gin-gonic/gin"

type (
	Middleware interface {
		ImplementCors() gin.HandlerFunc
	}

	middleware struct{}
)

func Init() Middleware {
	return &middleware{}
}
