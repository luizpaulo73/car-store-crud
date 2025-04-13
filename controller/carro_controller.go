package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
