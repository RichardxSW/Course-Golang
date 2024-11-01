package routes

import (
	"net/http"

	"example.com/webserver/models"
	"example.com/webserver/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users"})
		return
	}
	context.JSON(http.StatusOK, users)
}

func login(context *gin.Context) {
	var user models.User

    err := context.ShouldBindJSON(&user)

    if err!= nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
        return
    }

    err = user.ValidateCredentials()
	
    if err!= nil {
        context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not validate credentials"})
        return
    }

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}

    context.JSON(http.StatusOK, gin.H{"message": "Login success", "token": token})
}
 