package router

import (
	"api-avaliacao/src/controllers"

	"github.com/gin-gonic/gin"
)

func Rotas(r *gin.Engine) {
	r.POST("/", controllers.Criar)
	r.GET("/", controllers.BuscarTodos)
	r.GET("/id/:id", controllers.BuscarPorID)
	r.GET("/cpf/:cpf", controllers.BuscarPorCPF)
	r.GET("/cep/:cep", controllers.ConsultarCEP)
	r.PUT("/id/:id", controllers.AtualizarDadosID)
	r.PUT("/cpf/:cpf", controllers.AtualizarDadosCPF)
	r.DELETE("/:id", controllers.Delete)
}
