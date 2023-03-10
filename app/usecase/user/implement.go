package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/secmohammed/golang-kafka-grpc-poc/app/repository/user"
	"github.com/secmohammed/golang-kafka-grpc-poc/config"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	"github.com/secmohammed/golang-kafka-grpc-poc/pkg/queueing"
	"github.com/secmohammed/golang-kafka-grpc-poc/utils"
	"time"
)

type usecase struct {
	uc user.UserRepository
	c  config.Repository
	q  queueing.Messaging
}

func NewUseCase(uc user.UserRepository, c config.Repository, q queueing.Messaging) UseCase {
	return &usecase{uc: uc, c: c, q: q}
}

func (uc *usecase) Login(payload *entities.LoginUserInput) (*entities.UserLoginResponse, error) {

	result, err := uc.GetUserByEmail(payload.Email)
	if err != nil {

		return nil, utils.NewAuthorization("Invalid Credentials")
	}
	if err := utils.VerifyPassword(result.Password, payload.Password); err != nil {

		return nil, utils.NewAuthorization("Invalid Credentials")
	}
	secret, err := uc.c.GetString("app.jwt.secret")
	if err != nil {
		return nil, utils.NewAuthorization("failed to get secret")
	}
	expiration, err := uc.c.GetString("app.jwt.expiration")
	if err != nil {
		return nil, utils.NewAuthorization("failed to get jwt token expiration")
	}
	duration, err := time.ParseDuration(expiration)
	if err != nil {
		return nil, utils.NewAuthorization("failed to get parse jwt token expiration")
	}
	token, err := utils.GenerateToken(duration, result.ID, secret)
	if err != nil {
		return nil, utils.NewAuthorization("Failed to generate token")
	}
	userBytes, _ := json.Marshal(result)
	if err := uc.q.Write("users", []byte("login"), userBytes); err != nil {
		return nil, err
	}
	return &entities.UserLoginResponse{
		BaseModel: entities.BaseModel{
			ID:        result.ID,
			CreatedAt: result.CreatedAt,
		},
		Email: result.Email,
		Name:  result.Name,
		Token: token,
	}, nil
}
func (uc *usecase) GetUserByID(id uuid.UUID) (*entities.User, error) {
	u, err := uc.uc.GetUserByID(id)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return nil, utils.NewNotFound("user", id.String())
		}
		return nil, utils.NewBadRequest(err.Error())
	}
	return u, nil
}
func (uc *usecase) GetUserByEmail(email string) (*entities.User, error) {
	data, err := uc.uc.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return nil, utils.NewNotFound("user", email)
		}
		return nil, utils.NewBadRequest(err.Error())
	}
	return data, nil
}

func (uc *usecase) Create(payload *entities.RegisterUserInput) (*entities.User, error) {
	user, err := uc.uc.Create(payload)
	if err != nil {
		return nil, utils.NewBadRequest(fmt.Sprintf("failed to create user: %s", err))
	}
	userBytes, _ := json.Marshal(user)
	if err := uc.q.Write("users", []byte("register"), userBytes); err != nil {
		return nil, err
	}
	return user, nil
}
