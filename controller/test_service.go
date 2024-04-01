package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Testservice(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"mensage": "Welcome to the mato"})
}
