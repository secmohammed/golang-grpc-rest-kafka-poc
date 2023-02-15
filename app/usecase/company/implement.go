package company

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/secmohammed/golang-kafka-grpc-poc/app/repository/company"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	"gorm.io/gorm"
)

type usecase struct {
	cr company.CompanyRepository
}

func NewUseCase(cr company.CompanyRepository) UseCase {
	return &usecase{cr: cr}
}

func (uc *usecase) Get(id uuid.UUID) (*entities.Company, error) {
	data, err := uc.cr.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCompanyNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

func (uc *usecase) GetAll(page int) ([]*entities.Company, error) {
	data, err := uc.cr.GetAll(page)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}
func (uc *usecase) Create(in *entities.CreateCompanyInput) (*entities.Company, error) {
	data, err := uc.cr.Create(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

func (uc *usecase) Update(in *entities.UpdateCompanyInput) (*entities.Company, error) {
	data, err := uc.cr.Update(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil

}

func (uc *usecase) Delete(id uuid.UUID) error {
	err := uc.cr.Delete(id)
	if err != nil {
		if errors.Is(err, company.ErrCompanyNotFound) {
			return ErrCompanyNotFound
		}
		return fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}

	return err
}
