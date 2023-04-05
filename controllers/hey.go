package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hey(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "You're authenticated"})
}
