package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"test.com/config"
	"test.com/model"
)

func ListServidor(ctx *gin.Context) {
	log.Println("Init func Servidor")
	var servidor []model.Servidor
	if err := config.DB.Find(&servidor).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"mensage": "erro na func servidor no banco de dados",
		})
	}

	ctx.JSON(http.StatusOK, servidor)
}
