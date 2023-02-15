package company

import (
	"fmt"
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

func TestCannotCreateCompanyIfTokenNotPassed(t *testing.T) {
	payload := &entities.CreateCompanyInput{
		Name:        "something",
		Description: "omsoms",
		CompanyType: "Corporates",
		Registered:  false,
		Headcount:   1,
	}
	writer := tests.MakeRequest("POST", "/api/companies", payload, "", router(app))
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
func TestCannotCreateCompanyIfPassedTokenIsntValid(t *testing.T) {
	payload := &entities.CreateCompanyInput{
		Name:        "something",
		Description: "omsoms",
		CompanyType: "Corporates",
		Registered:  true,
		Headcount:   1,
	}
	writer := tests.MakeRequest("POST", "/api/companies", payload, "123123", router(app))
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestCannotCreateCompanyIfPassedCompanyTypeIsntPreviouslyDefined(t *testing.T) {
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
	payload := &entities.CreateCompanyInput{
		Name:        "something",
		Description: "loremdoqnwdonqwodnqwondqonwdonqwondnoqnoqwd",
		CompanyType: entities.CompanyType("Corporates"),
		Registered:  true,
		Headcount:   1,
	}
	writer := tests.MakeRequest("POST", "/api/companies", payload, token, router(app))
	assert.Equal(t, http.StatusBadRequest, writer.Code)
	app.Database().Get().Delete(&user)

}

func TestCreateCompanyIfPassedCompanyPayloadIsValid(t *testing.T) {
	password, err := utils.HashPassword("hello")
	assert.NoError(t, err)
	user := entities.User{
		Email:    "company@gmail.com",
		Password: password,
		Name:     "dqwdqw",
	}
	result := app.Database().Get().Create(&user)
	fmt.Println(user.ID)
	assert.NoError(t, result.Error)
	token, err := getToken(&user)
	assert.NoError(t, err)
	payload := &entities.CreateCompanyInput{
		Name:        "something",
		Description: "loremdoqnwdonqwodnqwondqonwdonqwondnoqnoqwd",
		CompanyType: entities.CompanyType("SoleProprietorship"),
		Registered:  true,
		Headcount:   1,
	}
	writer := tests.MakeRequest("POST", "/api/companies", payload, token, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	app.Database().Get().Delete(&user)

}
