package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luizpaulo73/model"
	"github.com/luizpaulo73/usecase"
	"github.com/luizpaulo73/validators"
)

type VendaController struct {
	VendaUsecase usecase.VendaUseCase
}

func NewVendaController(usecase usecase.VendaUseCase) VendaController {
	return VendaController{
		VendaUsecase: usecase,
	}
}

func (v *VendaController) GetVendas(ctx *gin.Context) {
	vendas, err := v.VendaUsecase.GetVendas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, vendas)
}

func (vc *VendaController) GetVendaById(ctx *gin.Context) {
	id := ctx.Param("id_cliente")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, "Id do cliente não pode ser nulo")
		return
	}

	idCliente, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Id do cliente inválido")
		return
	}

	vendas, err := vc.VendaUsecase.GetVendaByIdCliente(idCliente)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, vendas)
}

func (vc *VendaController) CreateVenda(ctx *gin.Context) {
	var venda model.Venda
	err := ctx.BindJSON(&venda)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "JSON inválido")
		return
	}

	err = validators.ValidarVenda(venda)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	vendaCriada, err := vc.VendaUsecase.CreateVenda(venda)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, vendaCriada)
}