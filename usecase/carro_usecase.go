package usecase

import (
	"errors"

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

func (carroUseCase *CarroUseCase) CreateCarro(carro model.Carro) (model.Carro, error) {
	carroId, err := carroUseCase.reposytory.CreateCarro(carro)
	if err != nil {
		return model.Carro{}, err
	}
	carro.ID = carroId

	return carro, nil
}

func (carroUseCase *CarroUseCase) DeleteCarro(id_carro int) (string, error) {
	carro, err := carroUseCase.reposytory.DeleteCarro(id_carro)
	if err != nil {
		return "Carro não encontrado", err
	}
	return carro, nil
}

func (carroUseCase *CarroUseCase) UpdateCarro(id int, carro model.Carro) (model.Carro, error) {
	updatedCarro, err := carroUseCase.reposytory.UpdateCarro(id, carro)
	if err != nil {
		return model.Carro{}, err
	}
	if updatedCarro == nil {
		return model.Carro{}, errors.New("carro não encontrado")
	}

	updatedCarro.ID = id

	return *updatedCarro, nil
}