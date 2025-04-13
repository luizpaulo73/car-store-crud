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

func (cr *CarroRepository) GetCarros() ([]model.Carro, error) {
	query := "SELECT * FROM carro"
	row, err := cr.connection.Query(query)
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

func (cr *CarroRepository) GetCarroById(id_carro int) (*model.Carro, error) {
	query, err := cr.connection.Prepare("SELECT * FROM carro WHERE id_carro = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer query.Close()

	var carro model.Carro

	err = query.QueryRow(id_carro).Scan(
		&carro.ID,
		&carro.Marca,
		&carro.Modelo,
		&carro.Ano,
		&carro.Cor,
		&carro.Preco,
		&carro.Quilometragem,
		&carro.Transmissao,
		&carro.Disponivel,)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}

	return &carro, nil
}