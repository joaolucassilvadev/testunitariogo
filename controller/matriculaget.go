package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"test.com/config"
	"test.com/model"
)

func MatriculaGET(ctx *gin.Context) {
	log.Print("Init func MatriculaGet")
	var servidor model.Servidor
	matricula := ctx.Param("Matricula_titular")
	if err := config.DB.First(&servidor, matricula).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Servidor não encontrado",
		})
		return
	}
	ctx.JSON(http.StatusOK, servidor)

	log.Printf("matricula get terminou")
}
