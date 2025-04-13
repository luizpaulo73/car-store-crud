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

func (carroUsecase *CarroUseCase) GetCarros() ([]model.Carro, error) {
	return carroUsecase.reposytory.GetCarros()
}

func (carroUseCase *CarroUseCase) GetCarroById(id_carro int) (*model.Carro, error) {
	carro, err := carroUseCase.reposytory.GetCarroById(id_carro)
	if err != nil {
		return nil, err
	}
	return carro, nil
}