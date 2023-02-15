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

func TestCannotDeleteCompanyIfTokenNotPassed(t *testing.T) {

    writer := tests.MakeRequest("DELETE", "/api/companies/123123", nil, "", router(app))
    assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
func TestCannotDeleteCompanyIfPassedTokenIsntValid(t *testing.T) {
    writer := tests.MakeRequest("DELETE", "/api/companies/12312312", nil, "123123", router(app))
    assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
func TestCannotDeleteCompanyIfIDIsInvalid(t *testing.T) {
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
    writer := tests.MakeRequest("DELETE", "/api/companies/12312312", nil, token, router(app))
    assert.Equal(t, http.StatusUnprocessableEntity, writer.Code)
    app.Database().Get().Delete(&user)

}
func TestCannotDeleteCompanyIfIDDoesntExist(t *testing.T) {
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

    writer := tests.MakeRequest("DELETE", "/api/companies/"+uuid.New().String(), nil, token, router(app))
    assert.Equal(t, http.StatusNotFound, writer.Code)
    app.Database().Get().Delete(&user)

}
func TestCanDeleteCompanyIfExists(t *testing.T) {
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
        Name:        "something",
        Description: "loremdoqnwdonqwodnqwondqonwdonqwondnoqnoqwd",
        CompanyType: entities.CompanyType("Corporates"),
        Registered:  true,
        Headcount:   1,
    }
    result = app.Database().Get().Create(&company)
    assert.NoError(t, result.Error)

    token, err := getToken(&user)
    assert.NoError(t, err)

    writer := tests.MakeRequest("DELETE", "/api/companies/"+company.ID.String(), nil, token, router(app))
    assert.Equal(t, http.StatusOK, writer.Code)
    app.Database().Get().Delete(&user)
    app.Database().Get().Delete(&company)

}
