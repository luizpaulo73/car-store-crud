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
		&cliente.CPF,
		&cliente.Senha)
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

func (cr *ClienteRepository) DeleteCliente(id_cliente int) (string, error) {
	query, err := cr.connection.Prepare("DELETE FROM cliente WHERE id_cliente = $1")
	if err != nil {
		fmt.Println(err)
		return "Erro ao deletar cliente", err
	}
	defer query.Close()

	_, err = query.Exec(id_cliente)
	if err != nil {
		return "Nao foi possivel deletar o cliente", err
	}

	return "Cliente deletado com sucesso", nil
}

func (cr *ClienteRepository) UpdateCliente(id_cliente int, cliente model.Cliente) (*model.Cliente, error) {

	var id int

	query, err := cr.connection.Prepare("UPDATE cliente SET nome = $1, email = $2, telefone" +
							" = $3, cpf = $4, senha = $5 WHERE id_cliente" +
							" = $6 RETURNING id_cliente")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer query.Close()

	err = query.QueryRow(
		cliente.Nome,
		cliente.Email,
		cliente.Telefone,
		cliente.CPF,
		cliente.Senha,
		id_cliente,
		).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cliente com id %d não encontrado", id_cliente)
		}
		fmt.Println("Erro ao atualizar cliente:", err)
		return nil, err
	}
	return &cliente, nil

}