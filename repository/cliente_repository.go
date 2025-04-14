package repository

import (
	"database/sql"
	"fmt"

	"github.com/luizpaulo73/model"
)

type ClienteRepository struct {
	connection *sql.DB
}

func NewClienteRepository(connection *sql.DB) ClienteRepository {
	return ClienteRepository{
		connection: connection,
	}
}

func (cr *ClienteRepository) GetClienteById(id_cliente int) (*model.Cliente, error) {
	query, err := cr.connection.Prepare("SELECT * FROM cliente WHERE id_cliente = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer query.Close()

	var cliente model.Cliente

	err = query.QueryRow(id_cliente).Scan(
		&cliente.ID,
		&cliente.Nome,
		&cliente.Email,
		&cliente.Telefone,
		&cliente.CPF,)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, nil
	}
	return &cliente, nil
}