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
	//camada de usecase
	CarroUsecase := usecase.NewCarroUseCase(CarroRepository)
	//camada de controllers
	CarroController := controller.NewCarroController(CarroUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "pong",
		})
	})

	server.GET("/carros", CarroController.GetCarros)

	server.GET("/carro/:id_carro", CarroController.GetCarroById)

	server.Run(":8000")
}