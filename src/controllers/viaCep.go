package controllers

import (
	"api-avaliacao/src/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConsultarCEP(ctx *gin.Context) {
	cep := ctx.Param("cep")
	if len(cep) != 8 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cep must be an integer of 8 digits",
		})
		return
	}

	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid cep",
		})
		return
	}
	defer response.Body.Close()

	var result models.CEP

	err = json.NewDecoder(response.Body).Decode(&result)
	if result.Erro {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid cep",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to decode results, check for invalid cep",
		})
		return
	}

	ctx.JSON(200, result)
}
