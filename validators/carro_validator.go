package validators

import (
	"errors"
	"time"

	"github.com/luizpaulo73/model"
)

func ValidarCarro(carro model.Carro) error {

	if carro.Marca == "" {
		return errors.New("marca não pode estar vazia")
	}

	if carro.Modelo == "" {
		return errors.New("modelo não pode estar vazio")
	}

	anoAtual := time.Now().Year()
	if carro.Ano > anoAtual+1 {
		return errors.New("ano inválido")
	}

	if carro.Cor == "" {
		return errors.New("cor não pode estar vazia")
	}

	if carro.Preco <= 0 {
		return errors.New("preço deve ser maior que zero")
	}

	if carro.Quilometragem < 0 {
		return errors.New("quilometragem não pode ser negativa")
	}

	if carro.Transmissao == "" {
		return errors.New("a transmissão não pode estar vazia")
	}

	return nil
}