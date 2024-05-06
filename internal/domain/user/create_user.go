package user_domain

import (
	"errors"
	"strings"

	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"github.com/LeandroVCastro/applying-manager-api/internal/repository"
)

type createUser struct {
	userRepository repository.UserRepository
}

func (u createUser) Handle(email, password string) (createdUser entity.User, err error) {
	if email == "" {
		err = errors.New("invalid email")
		return
	}
	_, errGetByEmail := u.userRepository.GetByEmail(strings.ToLower(email))
	if errGetByEmail == nil {
		err = errors.New("email already used")
	}
	return
}

func CreateUserFactory() createUser {
	return createUser{
		userRepository: repository.UserRepositoryFactory(),
	}
}
