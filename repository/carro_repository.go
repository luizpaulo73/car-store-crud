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

func (cr *CarroRepository) CreateCarro(carro model.Carro) (int, error) {

	var id int
	query, err := cr.connection.Prepare("INSERT INTO carro" + 
										"(marca, modelo, ano, cor, preco, quilometragem, transmissao, disponivel)" +
										"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id_carro")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(
		carro.Marca,
		carro.Modelo,
		carro.Ano,
		carro.Cor,
		carro.Preco,
		carro.Quilometragem,
		carro.Transmissao,
		carro.Disponivel ).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}

func (cr *CarroRepository) DeleteCarro(id_carro int) (string, error) {
	query, err := cr.connection.Prepare("DELETE FROM carro WHERE id_carro = $1")
	if err != nil {
		fmt.Println(err)
		return "Nao foi possivel deletar o carro", err
	}
	defer query.Close()

	_, err = query.Exec(id_carro)
	if err != nil {
		return "Nao foi possivel deletar o carro", err
	}
	return "Carro deletado com sucesso", nil
}

func (cr *CarroRepository) UpdateCarro(id_carro int, carro model.Carro) (*model.Carro, error) {
	var id int

	query, err := cr.connection.Prepare("UPDATE carro SET marca = $1, modelo = $2, ano" +
							" = $3, cor = $4, preco = $5, quilometragem = $6, transmissao = $7, disponivel = $8 WHERE id_carro" +
							" = $9 RETURNING id_carro")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer query.Close()

	err = query.QueryRow(
		carro.Marca,
		carro.Modelo,
		carro.Ano,
		carro.Cor,
		carro.Preco,
		carro.Quilometragem,
		carro.Transmissao,
		carro.Disponivel,
		id_carro,
	).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}

	return &carro, nil
}