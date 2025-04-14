package repository

import (
	"database/sql"
	"fmt"

	"github.com/luizpaulo73/model"
)

type VendaRepository struct {
	connection *sql.DB
}

func NewVendaRepository(connection *sql.DB) VendaRepository {
	return VendaRepository {
		connection: connection,
	}
}

func (vr *VendaRepository) GetVendas() ([]model.Venda, error) {
	query := "SELECT * FROM venda"
	row, err := vr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Venda{}, err
	}
	defer row.Close()

	var listaVenda []model.Venda
	var objVenda model.Venda

	for row.Next() {
		err = row.Scan(
			&objVenda.ID,
			&objVenda.DataVenda,
			&objVenda.PrecoVenda,
			&objVenda.IDCarro,
			&objVenda.IDCliente,)
		if err != nil {
			fmt.Println(err)
			return []model.Venda{}, err
		}
		listaVenda = append(listaVenda, objVenda)
	}

	return listaVenda, nil
}

func (vr *VendaRepository) GetVendaByIdCliente(idCliente int) ([]model.Venda, error) {
	query := "SELECT * FROM venda WHERE id_cliente = $1"
	rows, err := vr.connection.Query(query, idCliente)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var vendas []model.Venda

	for rows.Next() {
		var venda model.Venda
		err = rows.Scan(&venda.ID, &venda.DataVenda, &venda.PrecoVenda, &venda.IDCarro, &venda.IDCliente)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		vendas = append(vendas, venda)
	}

	return vendas, nil
}

func (vr *VendaRepository) CreateVenda(v model.Venda) (int, error) {
	query := `INSERT INTO venda (data_venda, preco_venda, id_carro, id_cliente)
			  VALUES ($1, $2, $3, $4) RETURNING id_venda`

	var id int
	err := vr.connection.QueryRow(query, v.DataVenda, v.PrecoVenda, v.IDCarro, v.IDCliente).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	updateQuery := `UPDATE carro SET disponivel = false WHERE id_carro = $1`
	_, err = vr.connection.Exec(updateQuery, v.IDCarro)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}