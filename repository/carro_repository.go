package repository

import (
	"database/sql"
	"fmt"

	"github.com/luizpaulo73/model"
)

type CarroRepository struct {
	connection *sql.DB
}

func NewCarroRepository (connection *sql.DB) CarroRepository {
	return CarroRepository {
		connection: connection,
	}
}

func (pr *CarroRepository) GetCarros() ([]model.Carro, error) {
	query := "SELECT * FROM carro"
	row, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Carro{}, err
	}
	defer row.Close()

	var listaCarro []model.Carro
	var objCarro model.Carro

	for row.Next() {
		err = row.Scan(
			&objCarro.ID,
			&objCarro.Marca,
			&objCarro.Modelo,
			&objCarro.Ano,
			&objCarro.Cor,
			&objCarro.Preco,
			&objCarro.Quilometragem,
			&objCarro.Transmissao,
			&objCarro.Disponivel,)
		if err != nil {
			fmt.Println(err)
			return []model.Carro{}, err
		}
		listaCarro = append(listaCarro, objCarro)
	}

	return listaCarro, nil
}