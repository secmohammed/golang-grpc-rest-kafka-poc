package user

import (
	"github.com/gin-gonic/gin"
	"github.com/secmohammed/golang-kafka-grpc-poc/app/repository/user"
	user2 "github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/user"
	"github.com/secmohammed/golang-kafka-grpc-poc/config"
	"github.com/secmohammed/golang-kafka-grpc-poc/container"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	"github.com/secmohammed/golang-kafka-grpc-poc/handlers/rest/middleware"
	"github.com/secmohammed/golang-kafka-grpc-poc/handlers/rest/users"
	"github.com/secmohammed/golang-kafka-grpc-poc/types"
	"log"
)

var c = config.Factory("local")
var app = container.NewApplication(c)

func setup() types.Container {
	if err := app.Database().Get().Migrator().AutoMigrate(&entities.User{}); err != nil {
		log.Fatal(err)
	}
	return app

}
func router(c types.Container) *gin.Engine {
	r := gin.New()
	uc := user.NewUserRepository(c)
	ucc := user2.NewUseCase(uc, c.Config())
	uh := users.NewUserHandler(ucc)
	rg := r.Group("/api/auth")
	rg.POST("/login", middleware.GuestUser(), uh.Login)
	rg.POST("/register", middleware.GuestUser(), uh.Register)
	return r

}
func teardown(c types.Container) {
	if err := c.Database().Get().Migrator().DropTable(&entities.User{}); err != nil {
		log.Fatal(err)
	}
}
