package controllers

import (
	"api-avaliacao/src/database"
	"api-avaliacao/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null/zero"
)

func Criar(ctx *gin.Context) {
	var pessoa models.Pessoa

	err := ctx.ShouldBindJSON(&pessoa)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "check if all obrigatory fields are filled correctly, " + err.Error(),
		})
		return
	}

	db := database.GetDB()

	err = db.Create(&pessoa).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "could not insert into database, " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, pessoa)
}

func BuscarTodos(ctx *gin.Context) {
	db := database.GetDB()

	var pessoas []models.Pessoa

	err := db.Find(&pessoas).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not select from database, " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, pessoas)
}

func BuscarPorID(ctx *gin.Context) {
	id := ctx.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be integer, " + err.Error(),
		})
		return
	}
	var pessoa models.Pessoa
	db := database.GetDB()
	err = db.Find(&pessoa, newid).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "unable to find, " + err.Error(),
		})
		return
	}
	if zero.Int.IsZero(pessoa.CPF) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "id non existent",
		})
		return
	}

	ctx.JSON(http.StatusOK, pessoa)
}

func BuscarPorCPF(ctx *gin.Context) {
	cpf := ctx.Param("cpf")
	newcpf, err := strconv.Atoi(cpf)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cpf must be an integer, " + err.Error(),
		})
		return
	}
	var pessoa models.Pessoa
	db := database.GetDB()
	err = db.Where("cpf = ?", newcpf).First(&pessoa).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "unable to find, " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, pessoa)
}

func AtualizarDadosID(ctx *gin.Context) {
	id := ctx.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be integer, " + err.Error(),
		})
		return
	}

	var pessoa models.Pessoa
	err = ctx.ShouldBindJSON(&pessoa)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "check if all obrigatory fields are filled correctly, " + err.Error(),
		})
		return
	}

	db := database.GetDB()
	err = db.Model(models.Pessoa{}).Where("id = ?", newid).Updates(pessoa).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not update database, " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func AtualizarDadosCPF(ctx *gin.Context) {
	cpf := ctx.Param("cpf")
	newcpf, err := strconv.Atoi(cpf)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cpf must be integer, " + err.Error(),
		})
		return
	}

	db := database.GetDB()

	var pessoa models.Pessoa
	err = ctx.ShouldBindJSON(&pessoa)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "check if all obrigatory fields are filled correctly, " + err.Error(),
		})
		return
	}

	err = db.Model(models.Pessoa{}).Where("cpf = ?", newcpf).Updates(pessoa).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "could not update database, " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cpf must be integer, " + err.Error(),
		})
		return
	}

	db := database.GetDB()
	err = db.Where("id = ?", newid).Delete(&models.Pessoa{}).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "could not update database, " + err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}
