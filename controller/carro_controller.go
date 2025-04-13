package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luizpaulo73/model"
	"github.com/luizpaulo73/usecase"
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
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, carros)
}

func (c *carroController) GetCarroById(ctx *gin.Context) {
	id := ctx.Param("id_carro")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, "Id do carro nao pode ser nulo")
		return
	}

	carroId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Id do carro invalido")
		return
	}

	carro, err := c.CarroUseCase.GetCarroById(carroId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if carro == nil {
		ctx.JSON(http.StatusNotFound, "Carro não encontrado")
		return
	}

	ctx.JSON(http.StatusOK, carro)
}

func (c *carroController) CreateCarro(ctx *gin.Context) {
	var carro model.Carro
	err := ctx.BindJSON(&carro)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	carroInserido, err := c.CarroUseCase.CreateCarro(carro)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, carroInserido)
}