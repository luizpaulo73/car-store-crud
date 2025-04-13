package model

type Venda struct {
	ID           int     `json:"id_venda"`
	DataVenda    string  `json:"data_venda"`
	PrecoVenda   float64 `json:"preco_venda"`
	IDCarro      int     `json:"id_carro"`
	IDCliente    int     `json:"id_cliente"`
}
