package utils

import (
	"net/http"

	"github.com/Dev-Awaab/go_job_portal_api/pkg/logger"
	"github.com/gin-gonic/gin"
)


func ErrorResponse(c *gin.Context, statusCode int, message string, err error, data interface{}) {
	

	
	response := gin.H{
		"message": message,
		"data":    data,
	}

	
	if err != nil {
		logger.Error("Request [%s %s]: %s, Error: %v", c.Request.Method, c.FullPath(), message, err)
	}

	
	logger.Info("Response [%s %s]: Status %d, Message: %s", c.Request.Method, c.FullPath(), statusCode, message)

	
	c.JSON(statusCode, response)
}


func ServerErrorResponse(c *gin.Context, context string, err error) {
	

	// Standard error message for internal server errors
	message := "Internal Server Error"
	response := gin.H{
		"message": message,
	}

	
	logger.Error("Server Error [%s] [%s %s]: %v", context, c.Request.Method, c.FullPath(), err)

	
	logger.Info("Response [%s %s]: Status %d, Message: %s", c.Request.Method, c.FullPath(), http.StatusInternalServerError, message)


	c.JSON(http.StatusInternalServerError, response)
}


func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	


	response := gin.H{
		"message": message,
		"data":    data,
	}

	
	logger.Info("Response [%s %s]: Status %d, Message: %s", c.Request.Method, c.FullPath(), statusCode, message)


	c.JSON(statusCode, response)
}