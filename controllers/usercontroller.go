package controllers

import (
	"ginwith-jwt/database"
	"ginwith-jwt/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		
		return
	}

	// perbandingan password hash
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		
		return
	}

	record := database.Instance.Create(&user)

	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()

		return
	}

	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username, "active_status": user.UserStatus})
}

// Get All User
func GetAllUser(context *gin.Context) {
	var user []models.User
	database.Instance.Find(&user)

	context.JSON(http.StatusOK, user)
}
