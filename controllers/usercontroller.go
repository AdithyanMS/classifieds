package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AdithyanMS/classifieds/auth"
	"github.com/AdithyanMS/classifieds/models"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	fmt.Println(user.Password)
	if err := user.HashPassword(user.Password); err != nil {
		log.Println("got err")
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	fmt.Println(user.Password)
	models.UserPasswords[user.Email] = user.Password
	context.JSON(http.StatusOK, gin.H{"email": user.Email})
}

func Login(context *gin.Context) {
	var creds, user models.User
	if err := context.ShouldBindJSON(&creds); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if hash, present := models.UserPasswords[creds.Email]; present {
		user.Email = creds.Email
		user.Password = hash
		if err := user.CheckPassword(creds.Password); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			context.Abort()
			return
		}
		tokenString, err := auth.GenerateJWT(user.Email)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
}
