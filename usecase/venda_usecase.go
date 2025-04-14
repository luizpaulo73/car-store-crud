package usecase

import (
	"github.com/luizpaulo73/model"
	"github.com/luizpaulo73/repository"
)

type VendaUseCase struct {
	repository repository.VendaRepository
}

func NewVendaUseCase(repo repository.VendaRepository) VendaUseCase {
	return VendaUseCase{
		repository: repo,
	}
}

func (vu *VendaUseCase) GetVendas() ([]model.Venda, error) {
	return vu.repository.GetVendas()
}

func (v *VendaUseCase) GetVendaByIdCliente(idCliente int) ([]model.Venda, error) {
	return v.repository.GetVendaByIdCliente(idCliente)
}

func (v *VendaUseCase) CreateVenda(venda model.Venda) (model.Venda, error) {
	id, err := v.repository.CreateVenda(venda)
	if err != nil {
		return model.Venda{}, err
	}
	venda.ID = id
	return venda, nil
}