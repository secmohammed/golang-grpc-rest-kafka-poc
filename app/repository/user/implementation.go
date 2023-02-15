package user

import (
	"github.com/google/uuid"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	"github.com/secmohammed/golang-kafka-grpc-poc/types"
	"github.com/secmohammed/golang-kafka-grpc-poc/utils"
	"strings"
)

type uc struct {
	c types.Container
}

func (uc *uc) GetUserByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	result := uc.c.Database().Get().First(user, "email = ? ", strings.ToLower(email))
	return user, result.Error
}
func (uc *uc) GetUserByID(id uuid.UUID) (*entities.User, error) {
	user := &entities.User{}
	result := uc.c.Database().Get().First(user, "id = ? ", id)
	return user, result.Error

}
func (uc *uc) Create(payload *entities.RegisterUserInput) (*entities.User, error) {
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	u := &entities.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
	}
	result := uc.c.Database().Get().Create(u)
	return u, result.Error

}

func NewUserRepository(c types.Container) UserRepository {
	return &uc{c}
}
