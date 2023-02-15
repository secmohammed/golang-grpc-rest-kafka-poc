package user

import (
	"github.com/google/uuid"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
)

type UseCase interface {
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByID(id uuid.UUID) (*entities.User, error)
	Create(payload *entities.RegisterUserInput) (*entities.User, error)
	Login(payload *entities.LoginUserInput) (*entities.UserLoginResponse, error)
}
