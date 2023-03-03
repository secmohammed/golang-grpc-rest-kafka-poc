package company

import (
    "github.com/gin-gonic/gin"
    "github.com/secmohammed/golang-kafka-grpc-poc/app/repository/company"
    "github.com/secmohammed/golang-kafka-grpc-poc/app/repository/user"
    company2 "github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/company"
    user2 "github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/user"
    "github.com/secmohammed/golang-kafka-grpc-poc/config"
    "github.com/secmohammed/golang-kafka-grpc-poc/container"
    "github.com/secmohammed/golang-kafka-grpc-poc/entities"
    "github.com/secmohammed/golang-kafka-grpc-poc/handlers/rest/companies"
    "github.com/secmohammed/golang-kafka-grpc-poc/handlers/rest/middleware"
    "github.com/secmohammed/golang-kafka-grpc-poc/types"
    "github.com/secmohammed/golang-kafka-grpc-poc/utils"
    "log"
    "time"
)

var c = config.Factory("local")
var app = container.NewApplication(c)

func setup() types.Container {
    if err := app.Database().Get().Migrator().AutoMigrate(&entities.Company{}, &entities.User{}); err != nil {
        log.Fatal(err)
    }
    return app

}
func router(c types.Container) *gin.Engine {
    r := gin.New()
    ucr := user.NewUserRepository(c)
    ucc := user2.NewUseCase(ucr, c.Config(), c.Queue())

    cr := company.NewCompanyRepository(c)
    uc := company2.NewUseCase(cr, c.Queue())
    ch := companies.NewCompanyHandler(uc)
    rg := r.Group("/api/companies")
    rg.POST("", middleware.AuthUser(ucc, c.Config()), ch.CreateCompany)
    rg.PATCH("/:id", middleware.AuthUser(ucc, c.Config()), ch.UpdateCompany)
    rg.GET("", ch.GetCompanies)
    rg.GET("/:id", ch.GetCompany)
    rg.DELETE("/:id", middleware.AuthUser(ucc, c.Config()), ch.DeleteCompany)
    return r

}
func teardown(c types.Container) {
    if err := c.Database().Get().Migrator().DropTable(&entities.Company{}, &entities.User{}); err != nil {
        log.Fatal(err)
    }
}
func getToken(u *entities.User) (string, error) {
    c := app.Config()
    expiration, err := c.GetString("app.jwt.expiration")
    if err != nil {
        return "", err
    }
    secret, err := c.GetString("app.jwt.secret")
    if err != nil {
        return "", err
    }
    parsedDuration, err := time.ParseDuration(expiration)
    if err != nil {
        return "", err
    }
    return utils.GenerateToken(parsedDuration, u.ID, secret)

}
