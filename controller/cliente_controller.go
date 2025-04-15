package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luizpaulo73/model"
	"github.com/luizpaulo73/usecase"
	"github.com/luizpaulo73/validators"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Id do cliente nao pode ser nulo"})
		return
	}

	clienteId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Id do cliente invalido"})
		return
	}

	cliente, err := c.ClienteUseCase.GetClienteById(clienteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	if cliente == nil {
		ctx.JSON(http.StatusNotFound, "Cliente não encontrado")
		return
	}

	ctx.JSON(http.StatusOK, cliente)
}

func (c *clienteController) CreateCliente(ctx *gin.Context) {
	var cliente model.Cliente
	err := ctx.BindJSON(&cliente)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err})
		return
	}

	err = validators.ValidarCliente(cliente)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	clienteInserido, err := c.ClienteUseCase.CreateCliente(cliente)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	ctx.JSON(http.StatusCreated, clienteInserido)
}

func (c *clienteController) DeleteCliente(ctx *gin.Context) {
	id := ctx.Param("id_cliente")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Id do cliente nao pode ser nulo"})
		return
}

	clienteId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Id do cliente invalido"})
		return
	}

	_, err = c.ClienteUseCase.DeleteCliente(clienteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err})
		return
	}

	ctx.JSON(http.StatusOK, "Cliente deletado com sucesso")
}

func (c *clienteController) UpdateCliente(ctx *gin.Context) {
	idParam := ctx.Param("id_cliente")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var cliente model.Cliente
	err = ctx.BindJSON(&cliente)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	err = validators.ValidarCliente(cliente)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	clienteAtualizado, err := c.ClienteUseCase.UpdateCliente(id, cliente)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, clienteAtualizado)
}