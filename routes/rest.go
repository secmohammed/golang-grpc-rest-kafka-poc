package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/secmohammed/golang-kafka-grpc-poc/app/repository/company"
	"github.com/secmohammed/golang-kafka-grpc-poc/app/repository/user"
	company2 "github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/company"
	user2 "github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/user"
	"github.com/secmohammed/golang-kafka-grpc-poc/handlers/rest/companies"
	"github.com/secmohammed/golang-kafka-grpc-poc/handlers/rest/middleware"
	users "github.com/secmohammed/golang-kafka-grpc-poc/handlers/rest/users"
	"github.com/secmohammed/golang-kafka-grpc-poc/types"
	"github.com/secmohammed/golang-kafka-grpc-poc/utils"
	"github.com/siruspen/logrus"
	"net/http"
	"time"
)

type rest struct {
	r *gin.Engine
	c types.Container
}

func NewRestRepository(c types.Container) *rest {
	env, err := c.Config().GetString("app.env")
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	logEnabled, err := c.Config().GetBool("app.log.debug")
	if err != nil {
		logrus.Error(err)
		return nil
	}
	r := gin.New()
	if logEnabled {
		r.Use(gin.Logger())

	}
	return &rest{c: c, r: r}
}
func setupDefaults(r *gin.Engine) {

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logMessage := fmt.Sprintf("%s |%s %d %s| %s |%s %s %s %s | %s | %s | %s\n",
			param.TimeStamp.Format(time.RFC1123),
			param.StatusCodeColor(),
			param.StatusCode,
			param.ResetColor(),
			param.ClientIP,
			param.MethodColor(),
			param.Method,
			param.ResetColor(),
			param.Path,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
		logrus.Info(fmt.Sprintf("%s | %d | %s | %s | %s | %s | %s | %s", param.TimeStamp.Format(time.RFC1123), param.StatusCode, param.ClientIP, param.Method, param.Path, param.Latency, param.Request.UserAgent(), param.ErrorMessage))
		return logMessage
	}))

	r.ForwardedByClientIP = true
	// recover from error when server fails to start and retry.
	r.Use(gin.Recovery())
	// Health check API
	r.GET("/api/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "OK"})
	})
	// If the user goes to a route that's not defined, we show the user that this route is not found.
	// fallback route.
	r.NoRoute(func(c *gin.Context) {
		c.JSON(utils.CreateApiError(http.StatusNotFound, fmt.Errorf("route %s not found", c.Request.URL.Path)))
	})
}

func (r *rest) registerUserRoutes() {
	uc := user.NewUserRepository(r.c)
	ucc := user2.NewUseCase(uc, r.c.Config())
	uh := users.NewUserHandler(ucc)
	rg := r.r.Group("/api/auth")
	rg.POST("/login", uh.Login)
	rg.POST("/register", uh.Register)

}
func (r *rest) registerCompanyRoutes() {
	ucr := user.NewUserRepository(r.c)
	ucc := user2.NewUseCase(ucr, r.c.Config())

	cr := company.NewCompanyRepository(r.c)
	uc := company2.NewUseCase(cr)
	ch := companies.NewCompanyHandler(uc)
	rg := r.r.Group("/api/companies")
	rg.POST("/", middleware.AuthUser(ucc, r.c.Config()), ch.CreateCompany)
	rg.PATCH("/:name", middleware.AuthUser(ucc, r.c.Config()), ch.UpdateCompany)
	rg.GET("/", ch.GetCompanies)
	rg.GET("/:id", ch.GetCompany)
	rg.DELETE("/:id", middleware.AuthUser(ucc, r.c.Config()), ch.DeleteCompany)
}

func (r *rest) Expose() error {
	port, err := r.c.Config().GetString("app.rest.port")
	if err != nil {
		return err
	}
	setupDefaults(r.r)
	r.registerCompanyRoutes()
	r.registerUserRoutes()

	return r.r.Run(fmt.Sprintf(":%s", port))
}
