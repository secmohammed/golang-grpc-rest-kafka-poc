package companies

import "github.com/gin-gonic/gin"

type CompaniesRestHandler interface {
	GetCompanies(c *gin.Context)
	GetCompany(c *gin.Context)
	CreateCompany(c *gin.Context)
	UpdateCompany(c *gin.Context)
	DeleteCompany(c *gin.Context)
}
