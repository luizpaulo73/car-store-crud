package validators

import (
	"errors"

	"github.com/luizpaulo73/model"
)

func ValidarVenda(venda model.Venda) error {
	if venda.DataVenda == "" {
		return errors.New("data da venda não pode ser vazia")
	}

	if venda.PrecoVenda <= 0 {
		return errors.New("preço da venda deve ser maior que zero")
	}

	if venda.IDCarro <= 0 {
		return errors.New("ID do carro deve ser maior que zero")
	}

	if venda.IDCliente <= 0 {
		return errors.New("ID do cliente deve ser maior que zero")
	}

	return nil
}