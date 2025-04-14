package usecase

import (
	"github.com/luizpaulo73/model"
	"github.com/luizpaulo73/repository"
)

type ClienteUseCase struct {
	repository repository.ClienteRepository
}

func NewClienteUseCase(repo repository.ClienteRepository) ClienteUseCase {
	return ClienteUseCase{
		repository: repo,
	}
}

func (clienteUseCase *ClienteUseCase) GetClienteById(id_cliente int) (*model.Cliente, error) {
	cliente, err := clienteUseCase.repository.GetClienteById(id_cliente)
	if err != nil {
		return nil, err		
	}
	return cliente, nil
}