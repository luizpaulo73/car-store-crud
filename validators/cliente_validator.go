package validators

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	"github.com/luizpaulo73/model"
)

func apenasNumeros(s string) bool {
	for _,r := range s{
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func EmailValido(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(strings.ToLower(email))
}

func ValidarCliente(cliente model.Cliente) error {
	if len(cliente.Nome) < 3 {
		return errors.New("o nome deve ter no mínimo 3 caracteres")
	}

	if cliente.Telefone == "" || !apenasNumeros(cliente.Telefone) {
		return errors.New("o telefone não é válido")
	}

	if len(cliente.CPF) != 11 || !apenasNumeros(cliente.CPF) {
		return errors.New("o CPF não é válido")
	}

	if len(cliente.Senha) < 6 {
		return errors.New("a senha deve ter no mínimo 6 caracteres")
	}

	if !EmailValido(cliente.Email) {
		return errors.New("o e-mail não é válido")
	}
	
	return nil
}