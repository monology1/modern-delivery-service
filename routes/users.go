package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"modern-delivery-service/models"
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
