package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"modern-delivery-service/models"
	"modern-delivery-service/utils"
	"net/http"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse body", "error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse body", "error": err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not validate credentials."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		log.Printf("Error while generating token: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not GenerateToken"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfully!", "token": token})
}
