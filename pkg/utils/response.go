package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, status int, message string) {
    c.JSON(status, gin.H{
        "status":  "error",
        "message": message,
    })
}

func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
    c.JSON(status, gin.H{
        "status":  "success",
        "message": message,
        "data":    data,
    })
}
