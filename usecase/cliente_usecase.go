package usecase

import (
	"errors"

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

func (clienteUseCase *ClienteUseCase) CreateCliente(cliente model.Cliente) (model.Cliente, error) {
	clienteId, err := clienteUseCase.repository.CreateCliente(cliente)
	if err != nil {
		return model.Cliente{}, err
	}
	cliente.ID = clienteId

	return cliente, nil
}

func (clienteUseCase *ClienteUseCase) DeleteCliente(id_cliente int) (string, error) {
	_, err := clienteUseCase.repository.DeleteCliente(id_cliente)
	if err != nil {
		return "Cliente nao encontrado", err
	}
	return "Cliente deletado com sucesso", nil
}

func (cliteUseCase *ClienteUseCase) UpdateCliente(id_cliente int, cliente model.Cliente) (model.Cliente, error) {
	updatedCliente, err := cliteUseCase.repository.UpdateCliente(id_cliente, cliente)
	if err != nil {
		return model.Cliente{}, err
	}
	if updatedCliente == nil {
		return model.Cliente{}, errors.New("cliente não encontrado")
	}

	updatedCliente.ID = id_cliente

	return *updatedCliente, nil
}