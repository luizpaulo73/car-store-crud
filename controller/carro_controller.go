package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luizpaulo73/model"
	"github.com/luizpaulo73/usecase"
	"github.com/luizpaulo73/validators"
)

type carroController struct {
	CarroUseCase usecase.CarroUseCase
}

func NewCarroController(usecase usecase.CarroUseCase) carroController {
	return carroController{
		CarroUseCase: usecase,
	}
}

func (c *carroController) GetCarros(ctx *gin.Context) {
	carros, err := c.CarroUseCase.GetCarros()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, carros)
}

func (c *carroController) GetCarroById(ctx *gin.Context) {
	id := ctx.Param("id_carro")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Id do carro nao pode ser nulo"})
		return
	}

	carroId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Id do carro invalido"})
		return
	}

	carro, err := c.CarroUseCase.GetCarroById(carroId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	if carro == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": "Carro não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, carro)
}

func (c *carroController) CreateCarro(ctx *gin.Context) {
	var carro model.Carro
	err := ctx.BindJSON(&carro)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	err = validators.ValidarCarro(carro)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	carroInserido, err := c.CarroUseCase.CreateCarro(carro)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, carroInserido)
}

func (c *carroController) DeleteCarro(ctx *gin.Context) {
	id := ctx.Param("id_carro")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Id do carro nao pode ser nulo"})
		return
	}

	carroId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Id do carro invalido"})
		return
	}

	carro, err := c.CarroUseCase.DeleteCarro(carroId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, carro)
}

func (c *carroController) UpdateCarro(ctx *gin.Context) {
	idParam := ctx.Param("id_carro")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var carro model.Carro

	err = ctx.BindJSON(&carro)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	
	err = validators.ValidarCarro(carro)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	carroAtualizado, err := c.CarroUseCase.UpdateCarro(id, carro)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, carroAtualizado)
}
