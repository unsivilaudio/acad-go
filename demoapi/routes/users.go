package routes

import (
	"demoapi/models"
	"demoapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not parse data.",
			},
		)
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Could not authenticate user.",
			},
		)
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Trouble signing you in, try again later.",
			},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Welcome back!",
			"token":   token,
		},
	)
}

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not parse data.",
			},
		)
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Could not save user.",
			},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{
			"message": "User created sucessfully.",
		},
	)
}
