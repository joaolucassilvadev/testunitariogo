package router

import (
	"github.com/gin-gonic/gin"
	"test.com/controller"
)

func Routes(ctx *gin.Engine) {
	ctx.GET("test_service", controller.Testservice)
	ctx.GET("ListServidor", controller.ListServidor)
	ctx.PUT("EditServidor/:Cpf_do_titular", controller.UpdateServidor)
	ctx.GET("Getmatricula/:Matricula_titular", controller.MatriculaGET)
}
