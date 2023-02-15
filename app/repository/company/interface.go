package company

import (
	"errors"
	"github.com/google/uuid"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
)

var ErrCompanyNotFound = errors.New("company not found")

type CompanyRepository interface {
	Get(id uuid.UUID) (*entities.Company, error)
	GetAll(page int) ([]*entities.Company, error)
	Create(company *entities.CreateCompanyInput) (*entities.Company, error)
	Update(company *entities.UpdateCompanyInput) (*entities.Company, error)
	Delete(id uuid.UUID) error
}
