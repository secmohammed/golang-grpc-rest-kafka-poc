package routes

import (
	"context"
	"fmt"
	"github.com/secmohammed/golang-kafka-grpc-poc/app/repository/company"
	"github.com/secmohammed/golang-kafka-grpc-poc/app/repository/user"
	company2 "github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/company"
	user2 "github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/user"
	companies2 "github.com/secmohammed/golang-kafka-grpc-poc/handlers/grpc/companies"
	"github.com/secmohammed/golang-kafka-grpc-poc/handlers/grpc/pb/companies"
	"github.com/secmohammed/golang-kafka-grpc-poc/handlers/grpc/pb/users"
	users2 "github.com/secmohammed/golang-kafka-grpc-poc/handlers/grpc/users"
	"github.com/secmohammed/golang-kafka-grpc-poc/types"
	"github.com/secmohammed/golang-kafka-grpc-poc/utils"
	log "github.com/siruspen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"time"

	"net"
)

type grpcClient struct {
	c types.Container
	s *grpc.Server
}

func (g *grpcClient) serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	if info.FullMethod == "/Companies/CreateCompany" || info.FullMethod == "/Companies/UpdateCompany" || info.FullMethod == "/Companies/DeleteCompany" {
		if err := g.authorize(ctx); err != nil {
			return nil, err
		}
	}
	// Calls the handler
	h, err := handler(ctx, req)

	log.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

// authorize function authorizes the token received from Metadata
func (g *grpcClient) authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}
	authHeader, ok := md["authorization"]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	token := authHeader[0]
	// validateToken function validates the token
	secret, err := g.c.Config().GetString("app.jwt.secret")
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}
	_, err = utils.ValidateToken(token, secret)

	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}
	return nil
}
func (g *grpcClient) withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(g.serverInterceptor)
}
func NewGRPCRepository(c types.Container) *grpcClient {
	client := &grpcClient{
		c: c,
	}
	s := grpc.NewServer(client.withServerUnaryInterceptor())
	reflection.Register(s)
	client.s = s
	return client
}

func (g *grpcClient) Expose() error {
	port, err := g.c.Config().GetString("app.grpc.port")
	if err != nil {
		return err
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}
	ucr := user.NewUserRepository(g.c)
	ucc := user2.NewUseCase(ucr, g.c.Config(), g.c.Queue())
	cr := company.NewCompanyRepository(g.c)
	uc := company2.NewUseCase(cr, g.c.Queue())
	companies.RegisterCompaniesServer(g.s, companies2.NewCompanyServer(uc, g.c.Config(), ucc))

	users.RegisterUsersServer(g.s, users2.NewUserServer(ucc))
	if err := g.s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}
	return nil
}
