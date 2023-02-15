package company

import (
	"github.com/google/uuid"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	"github.com/secmohammed/golang-kafka-grpc-poc/tests"
	"github.com/secmohammed/golang-kafka-grpc-poc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCanGetCompanyIfTokenNotPassed(t *testing.T) {

	writer := tests.MakeRequest("GET", "/api/companies/"+uuid.New().String(), nil, "", router(app))
	assert.Equal(t, http.StatusNotFound, writer.Code)
}

func TestCannotGetCompanyIfIDIsInvalid(t *testing.T) {

	writer := tests.MakeRequest("GET", "/api/companies/12312312", nil, "", router(app))
	assert.Equal(t, http.StatusUnprocessableEntity, writer.Code)

}
func TestCanGetCompanyIfTokenPassed(t *testing.T) {
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
	writer := tests.MakeRequest("GET", "/api/companies/"+uuid.New().String(), nil, token, router(app))
	assert.Equal(t, http.StatusNotFound, writer.Code)
	app.Database().Get().Where("id = ? ", user.ID).Delete(&user)
}
func TestCannotGetCompanyIfIDDoesntExist(t *testing.T) {

	writer := tests.MakeRequest("GET", "/api/companies/"+uuid.New().String(), nil, "", router(app))
	assert.Equal(t, http.StatusNotFound, writer.Code)

}
func TestCanGetCompanyIfExists(t *testing.T) {
	company := entities.Company{
		Name:        "something",
		Description: "loremdoqnwdonqwodnqwondqonwdonqwondnoqnoqwd",
		CompanyType: entities.CompanyType("Corporates"),
		Registered:  true,
		Headcount:   1,
	}
	result := app.Database().Get().Create(&company)
	assert.NoError(t, result.Error)

	writer := tests.MakeRequest("GET", "/api/companies/"+company.ID.String(), nil, "", router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	app.Database().Get().Delete(&company)

}
