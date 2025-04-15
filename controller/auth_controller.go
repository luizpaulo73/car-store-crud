package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizpaulo73/model"
	"github.com/luizpaulo73/usecase"
)

type authController struct {
	AuthUseCase usecase.AuthUseCase
}

func NewAuthController(usecase usecase.AuthUseCase) authController {
	return authController{
		AuthUseCase: usecase,
	}
}

func (a *authController) Login(ctx *gin.Context) {
	var user model.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	if user.Email == "" || user.Senha == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Email e senha são obrigatórios"})
		return
	}

	ok, err := a.AuthUseCase.Auth(user.Email, user.Senha)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	if ok {
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"success": false, "erro": "Email ou senha inválidos"})
	}
}