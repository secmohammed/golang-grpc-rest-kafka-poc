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

func TestCannotUpdateCompanyIfTokenNotPassed(t *testing.T) {
	payload := &entities.CreateCompanyInput{
		Name:        "something",
		Description: "omsoms",
		CompanyType: "Corporates",
		Registered:  false,
		Headcount:   1,
	}
	writer := tests.MakeRequest("PATCH", "/api/companies/123", payload, "", router(app))
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
func TestCannotUpdateCompanyIfPassedTokenIsntValid(t *testing.T) {
	payload := &entities.CreateCompanyInput{
		Name:        "something",
		Description: "omsoms",
		CompanyType: "Corporates",
		Registered:  true,
		Headcount:   1,
	}
	writer := tests.MakeRequest("PATCH", "/api/companies/123", payload, "123123", router(app))
	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestCannotUpdateCompanyIfPassedCompanyTypeIsntPreviouslyDefined(t *testing.T) {
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
	payload := &entities.UpdateCompanyInput{
		ID:          uuid.New(),
		Name:        "something",
		Description: "loremdoqnwdonqwodnqwondqonwdonqwondnoqnoqwd",
		CompanyType: entities.CompanyType("Corporates"),
		Registered:  true,
		Headcount:   1,
	}
	writer := tests.MakeRequest("PATCH", "/api/companies/12312", payload, token, router(app))
	assert.Equal(t, http.StatusBadRequest, writer.Code)
	app.Database().Get().Where("id = ?", user.ID).Delete(&user)

}

func TestCannotUpdateCompanyIfPassedCompanyIDDoesntExist(t *testing.T) {
	password, err := utils.HashPassword("hello")
	assert.NoError(t, err)
	user := entities.User{
		Email:    "company@gmail.com",
		Password: password,
		Name:     "dqwdqw",
	}
	result := app.Database().Get().Create(&user)
	assert.NoError(t, result.Error)
	company := &entities.UpdateCompanyInput{
		Name:        "test-compnay",
		Description: "loremdoqnwdonqwodnqwondqonwdonqwondnoqnoqwd",
		Registered:  true,
		Headcount:   123,
		CompanyType: "SoleProprietorship",
	}
	token, err := getToken(&user)
	assert.NoError(t, err)
	writer := tests.MakeRequest("PATCH", "/api/companies/12312", company, token, router(app))
	assert.Equal(t, http.StatusUnprocessableEntity, writer.Code)
	app.Database().Get().Where("id = ?", user.ID).Delete(&user)
}
func TestUpdateCompanyIfPassedCompanyPayloadIsValid(t *testing.T) {
	password, err := utils.HashPassword("hello")
	assert.NoError(t, err)
	user := entities.User{
		Email:    "company@gmail.com",
		Password: password,
		Name:     "dqwdqw",
	}
	result := app.Database().Get().Create(&user)
	assert.NoError(t, result.Error)
	company := entities.Company{
		Name:        "test-compnay",
		Description: "loremdoqnwdonqwodnqwondqonwdonqwondnoqnoqwd",
		Registered:  true,
		Headcount:   123,
		CompanyType: "SoleProprietorship",
	}
	result = app.Database().Get().Create(&company)
	assert.NoError(t, result.Error)
	token, err := getToken(&user)
	assert.NoError(t, err)
	assert.NotContains(t, token, "00000000")
	writer := tests.MakeRequest("PATCH", "/api/companies/"+company.ID.String(), &entities.UpdateCompanyInput{
		Name:        company.Name,
		Description: company.Description,
		Registered:  true,
		Headcount:   40,
		CompanyType: company.CompanyType,
	}, token, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	assert.Contains(t, writer.Body.String(), "40")
	app.Database().Get().Where("id = ?", user.ID).Delete(&user)
	app.Database().Get().Where("id = ?", company.ID).Delete(&company)
}
