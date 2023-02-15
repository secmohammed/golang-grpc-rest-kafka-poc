package user

import (
	"errors"
	"github.com/google/uuid"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
)

var ErrUserNotFound = errors.New("company not found")

type UserRepository interface {
	GetUserByEmail(email string) (*entities.User, error)
	GetUserByID(id uuid.UUID) (*entities.User, error)
	Create(payload *entities.RegisterUserInput) (*entities.User, error)
}
