package model

type Carro struct {
	ID             int     `json:"id_carro"`
	Marca          string  `json:"marca"`
	Modelo         string  `json:"modelo"`
	Ano            int     `json:"ano"`
	Cor            string  `json:"cor"`
	Preco          float64 `json:"preco"`
	Quilometragem  int     `json:"quilometragem"`
	Transmissao    string  `json:"transmissao"`
	Disponivel     bool    `json:"disponivel"`
}