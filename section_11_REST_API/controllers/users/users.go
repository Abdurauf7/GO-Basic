package users

import (
	"net/http"
	"practice/Rest/models"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed bind data"})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could'n save users"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
