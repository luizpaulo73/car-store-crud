package main

import (
	"github.com/luizpaulo73/controller"
	"github.com/luizpaulo73/db"
	"github.com/luizpaulo73/repository"
	"github.com/luizpaulo73/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada repository
	CarroRepository := repository.NewCarroRepository(dbConnection)
	ClienteRepository := repository.NewClienteRepository(dbConnection)
	VendaRepository := repository.NewVendaRepository(dbConnection)
	//camada de usecase
	CarroUsecase := usecase.NewCarroUseCase(CarroRepository)
	ClienteUsecase := usecase.NewClienteUseCase(ClienteRepository)
	VendaUsecase := usecase.NewVendaUseCase(VendaRepository)
	//camada de controllers
	CarroController := controller.NewCarroController(CarroUsecase)
	ClienteController := controller.NewClienteController(ClienteUsecase)
	VendaController := controller.NewVendaController(VendaUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "pong",
		})
	})

	//Carro
	server.GET("/carros", CarroController.GetCarros)

	server.GET("/carro/:id_carro", CarroController.GetCarroById)

	server.POST("/carro/cadastrar", CarroController.CreateCarro)

	server.DELETE("/carro/deletar/:id_carro", CarroController.DeleteCarro)

	server.PUT("/carro/atualizar/:id_carro", CarroController.UpdateCarro)

	//Cliente
	server.GET("/cliente/:id_cliente", ClienteController.GetClienteById)

	server.POST("/cliente/cadastro", ClienteController.CreateCliente)

	server.DELETE("/cliente/excluir/:id_cliente", ClienteController.DeleteCliente)

	server.PUT("/cliente/atualizar/:id_cliente", ClienteController.UpdateCliente)

	//Venda
	server.GET("/vendas", VendaController.GetVendas)

	server.GET("/venda/:id_cliente", VendaController.GetVendaById)

	server.POST("/venda/realizar", VendaController.CreateVenda)

	server.Run(":8000")
}