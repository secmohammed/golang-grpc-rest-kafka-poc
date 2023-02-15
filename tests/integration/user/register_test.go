package user

import (
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	"github.com/secmohammed/golang-kafka-grpc-poc/tests"
	"github.com/secmohammed/golang-kafka-grpc-poc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRegisterUserWithExistingEmail(t *testing.T) {
	password, err := utils.HashPassword("hello")
	assert.NoError(t, err)
	user := &entities.User{
		Email:    "someone@gmail.com",
		Password: password,
		Name:     "dqwdqw",
	}
	app.Database().Get().Create(user)
	registerInput := &entities.RegisterUserInput{
		Email:                user.Email,
		Password:             "hello",
		Name:                 user.Name,
		PasswordConfirmation: "hello",
	}
	writer := tests.MakeRequest("POST", "/api/auth/register", registerInput, "", router(app))
	assert.Contains(t, writer.Body.String(), "users_email_key")
	assert.Equal(t, http.StatusBadRequest, writer.Code)
	app.Database().Get().Delete(user)
}
func TestRegisterWithNonMatchingPassword(t *testing.T) {
	registerInput := &entities.RegisterUserInput{
		Email:                "someone@gmail.com",
		Password:             "hello",
		Name:                 "someone",
		PasswordConfirmation: "123",
	}
	writer := tests.MakeRequest("POST", "/api/auth/register", registerInput, "", router(app))
	assert.Contains(t, writer.Body.String(), "PasswordConfirmation")
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}
func TestCannotRegisterUserIfBearerTokenPersist(t *testing.T) {
	registerInput := &entities.RegisterUserInput{
		Email:                "someone@gmail.com",
		Password:             "hello",
		Name:                 "someone",
		PasswordConfirmation: "hello",
	}
	writer := tests.MakeRequest("POST", "/api/auth/register", registerInput, "12312321", router(app))
	assert.Contains(t, writer.Body.String(), "You are already logged in")
	assert.Equal(t, http.StatusBadRequest, writer.Code)
}
func TestRegisterWithValidPayload(t *testing.T) {
	registerInput := &entities.RegisterUserInput{
		Email:                "someone@gmail.com",
		Password:             "hello",
		Name:                 "someone",
		PasswordConfirmation: "hello",
	}
	writer := tests.MakeRequest("POST", "/api/auth/register", registerInput, "", router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	app.Database().Get().Where("email = ?", registerInput.Email).Delete(&entities.User{})
}
