package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"test.com/config"
	"test.com/model"
)

type Updateservidor struct {
	Id   string
	Name string
	cpf  int
}

func UpdateServidor(ctx *gin.Context) {
	log.Printf("Init func UpdateServidor")
	var servidor model.Servidor
	var updateServidor Updateservidor
	Cpf := ctx.Param("Cpf_do_titular")

	if err := ctx.ShouldBindJSON(&updateServidor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"mensage": "erro na atualização do servidor",
			"error":   err.Error(),
		})
		return
	}

	if result := config.DB.First(&servidor, Cpf); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"mensage error": "Usuario não encontrado",
			"Error":         result.Error,
		})
		return
	}

	servidor.Cpf_do_titular = updateServidor.cpf
	servidor.Nome_do_titular = updateServidor.Name

	config.DB.Save(&servidor)

	ctx.JSON(http.StatusOK, gin.H{
		"mensage":  "Servidor cadastrado com sucesso",
		"Servidor": &servidor,
	})
	log.Printf("end func UpdateServidor")
}
