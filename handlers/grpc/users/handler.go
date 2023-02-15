package users

import (
	"context"
	"github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/user"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	pb "github.com/secmohammed/golang-kafka-grpc-poc/handlers/grpc/pb/users"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedUsersServer
	ucc user.UseCase
}

func (s *server) Login(_ context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	payload := &entities.LoginUserInput{
		Email:    request.Email,
		Password: request.Password,
	}
	result, err := s.ucc.Login(payload)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Id:        result.ID.String(),
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: timestamppb.New(result.CreatedAt),
		UpdatedAt: timestamppb.New(result.UpdatedAt),
		Token:     result.Token,
	}, nil
}
func (s *server) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	payload := &entities.RegisterUserInput{
		Email:                request.Email,
		Name:                 request.Name,
		Password:             request.Password,
		PasswordConfirmation: request.PasswordConfirmation,
	}
	result, err := s.ucc.Create(payload)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		User: &pb.User{
			Id:        result.ID.String(),
			Email:     result.Email,
			Name:      result.Name,
			CreatedAt: timestamppb.New(result.CreatedAt),
			UpdatedAt: timestamppb.New(result.UpdatedAt),
		},
	}, nil
}
func NewUserServer(ucc user.UseCase) pb.UsersServer {
	return &server{
		pb.UnimplementedUsersServer{},
		ucc,
	}
}
