package usecase

import (
	"github.com/luizpaulo73/model"
	"github.com/luizpaulo73/repository"
)

type CarroUseCase struct {
	reposytory repository.CarroRepository
}

func NewCarroUseCase(repo repository.CarroRepository) CarroUseCase {
	return CarroUseCase{
		reposytory: repo,
	}
}

func (pu *CarroUseCase) GetCarros() ([]model.Carro, error) {
	return pu.reposytory.GetCarros()
}