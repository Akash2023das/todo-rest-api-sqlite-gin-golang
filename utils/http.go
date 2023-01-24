package utils

import "github.com/gin-gonic/gin"

// Enum for HTTP methods
type HttpMethod string

const ( 
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

// Response - This function writes a success response to the response writer
func Response(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(status, gin.H{
		"status":  status,
		"data":    data,
		"message": message,
	})
}
