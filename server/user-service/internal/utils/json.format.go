package utils

import "github.com/gin-gonic/gin"

func JSONResponse(c *gin.Context, statusCode int, status, message string, data any) {
	c.JSON(statusCode, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func JSONError(c *gin.Context, code int, status, message string) {
	c.JSON(code, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: message,
	})
}
