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

func (cr *ClienteRepository) CreateCliente(cliente model.Cliente) (int, error) {

	var id int
	query, err := cr.connection.Prepare("INSERT INTO cliente" + 
										"(nome, email, telefone, cpf, senha)" +
										"VALUES ($1, $2, $3, $4, $5) RETURNING id_cliente")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	err = query.QueryRow(
		cliente.Nome,
		cliente.Email,
		cliente.Telefone,
		cliente.CPF,
		cliente.Senha, ).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}