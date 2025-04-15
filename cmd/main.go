package main

import (
	"github.com/luizpaulo73/controller"
	"github.com/luizpaulo73/db"
	"github.com/luizpaulo73/repository"
	"github.com/luizpaulo73/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
	}))

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada repository
	CarroRepository := repository.NewCarroRepository(dbConnection)
	ClienteRepository := repository.NewClienteRepository(dbConnection)
	VendaRepository := repository.NewVendaRepository(dbConnection)
	AuthRepository  := repository.NewAuthRepository(dbConnection)
	//camada de usecase
	CarroUsecase := usecase.NewCarroUseCase(CarroRepository)
	ClienteUsecase := usecase.NewClienteUseCase(ClienteRepository)
	VendaUsecase := usecase.NewVendaUseCase(VendaRepository)
	AuthUsecase := usecase.NewAuthUseCase(AuthRepository)
	//camada de controllers
	CarroController := controller.NewCarroController(CarroUsecase)
	ClienteController := controller.NewClienteController(ClienteUsecase)
	VendaController := controller.NewVendaController(VendaUsecase)
	AuthController := controller.NewAuthController(AuthUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "pong",
		})
	})

	//Auth
	server.POST("/auth", AuthController.Login)

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