package usecase

import "github.com/luizpaulo73/repository"

type AuthUseCase struct {
	repository repository.AuthRepository
}

func NewAuthUseCase(repo repository.AuthRepository) AuthUseCase {
	return AuthUseCase{
		repository: repo,
	}
}

func (au *AuthUseCase) Auth(email, senha string) (bool, error) {
	senhaBD, err := au.repository.GetAuthUser(email)
	if err != nil {
		return false, err
	}
	if senhaBD == "" {
		return false, nil
	}
	return senhaBD == senha, nil
}