package repository

import (
	"errors"

	"github.com/LeandroVCastro/applying-manager-api/internal/configs"
	"github.com/LeandroVCastro/applying-manager-api/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	connection *gorm.DB
}

func (repository UserRepository) GetByEmail(email string) (userFound entity.User, err error) {
	var user = entity.User{}
	result := repository.connection.Where("email = ?", email).First(&user)
	if result.RowsAffected > 0 {
		userFound = user
		return
	}
	err = errors.New("user not found on GetByEmail function")
	return
}

func UserRepositoryFactory() UserRepository {
	repository := UserRepository{
		connection: configs.Connection,
	}
	return repository
}
