package company

import (
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	"github.com/secmohammed/golang-kafka-grpc-poc/tests"
	"github.com/secmohammed/golang-kafka-grpc-poc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCanGetCompaniesIfTokenNotPassed(t *testing.T) {

	writer := tests.MakeRequest("GET", "/api/companies", nil, "", router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestCanGetCompaniesIfTokenPassed(t *testing.T) {
	password, err := utils.HashPassword("hello")
	assert.NoError(t, err)
	user := entities.User{
		Email:    "company@gmail.com",
		Password: password,
		Name:     "dqwdqw",
	}
	result := app.Database().Get().Create(&user)
	assert.NoError(t, result.Error)
	token, err := getToken(&user)
	assert.NoError(t, err)
	writer := tests.MakeRequest("GET", "/api/companies", nil, token, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	app.Database().Get().Delete(&user)
}
