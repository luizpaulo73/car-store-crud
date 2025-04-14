package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luizpaulo73/usecase"
)

type clienteController struct {
	ClienteUseCase usecase.ClienteUseCase
}

func NewClienteController(usecase usecase.ClienteUseCase) clienteController {
	return clienteController{
		ClienteUseCase: usecase,
	}
}

func (c *clienteController) GetClienteById(ctx *gin.Context){
	id := ctx.Param("id_cliente")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, "Id do cliente nao pode ser nulo")
		return
	}

	clienteId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Id do cliente invalido")
		return
	}

	cliente, err := c.ClienteUseCase.GetClienteById(clienteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if cliente == nil {
		ctx.JSON(http.StatusNotFound, "Cliente não encontrado")
		return
	}

	ctx.JSON(http.StatusOK, cliente)
}