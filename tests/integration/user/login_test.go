package user

import (
	"github.com/gin-gonic/gin"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	"github.com/secmohammed/golang-kafka-grpc-poc/tests"
	"github.com/secmohammed/golang-kafka-grpc-poc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	c := setup()
	exitCode := m.Run()
	teardown(c)

	os.Exit(exitCode)
}

func TestLoginWithNotFoundUser(t *testing.T) {
	loginInput := &entities.LoginUserInput{
		Email:    "someone@gmail.com",
		Password: "onqwe",
	}

	writer := tests.MakeRequest("POST", "/api/auth/login", loginInput, "", router(app))
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
func TestLoginWithInvalidUserPassword(t *testing.T) {
	password, err := utils.HashPassword("hello")
	assert.NoError(t, err)
	user := entities.User{
		Email:    "someone@gmail.com",
		Password: password,
		Name:     "dqwdqw",
	}
	app.Database().Get().Create(&user)
	loginInput := &entities.LoginUserInput{
		Email:    "someone@gmail.com",
		Password: "123",
	}

	writer := tests.MakeRequest("POST", "/api/auth/login", loginInput, "", router(app))
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
	app.Database().Get().Delete(&user)

}
func TestCannotLoginIfTokenPersist(t *testing.T) {
	loginInput := &entities.LoginUserInput{
		Email:    "someone@gmail.com",
		Password: "hello",
	}
	writer := tests.MakeRequest("POST", "/api/auth/login", loginInput, "123123", router(app))

	assert.Contains(t, writer.Body.String(), "You are already logged in")
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}
func TestLoginWithValidCredentials(t *testing.T) {
	password, err := utils.HashPassword("hello")
	assert.NoError(t, err)
	user := entities.User{
		Email:    "someone@gmail.com",
		Password: password,
		Name:     "dqwdqw",
	}
	app.Database().Get().Create(&user)
	loginInput := &entities.LoginUserInput{
		Email:    "someone@gmail.com",
		Password: "hello",
	}

	writer := tests.MakeRequest("POST", "/api/auth/login", loginInput, "", router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	app.Database().Get().Delete(&user)
}
