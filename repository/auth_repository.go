package repository

import (
	"database/sql"
	"fmt"
)

type AuthRepository struct {
	connection *sql.DB
}

func NewAuthRepository(connection *sql.DB) AuthRepository {
	return AuthRepository{
		connection: connection,
	}
}

func (ar *AuthRepository) GetAuthUser(email string) (string, error) {
	var senha string
	query, err := ar.connection.Prepare("SELECT senha FROM cliente WHERE email = $1")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer query.Close()

	err = query.QueryRow(email).Scan(&senha)
	if err != nil {
		if err == sql.ErrNoRows {
			return "Login ou senha inválidos", err
		}
		return "", err
	}
	return senha, nil
}