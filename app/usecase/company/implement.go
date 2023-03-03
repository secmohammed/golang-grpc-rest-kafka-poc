package company

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/secmohammed/golang-kafka-grpc-poc/app/repository/company"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	"github.com/secmohammed/golang-kafka-grpc-poc/pkg/queueing"
	"gorm.io/gorm"
)

type usecase struct {
	cr company.CompanyRepository
	q  queueing.Messaging
}

func NewUseCase(cr company.CompanyRepository, q queueing.Messaging) UseCase {
	return &usecase{cr: cr, q: q}
}

func (uc *usecase) Get(id uuid.UUID) (*entities.Company, error) {
	data, err := uc.cr.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCompanyNotFound
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	dataBytes, _ := json.Marshal(data)
	if err := uc.q.Write("companies", []byte("getOne"), dataBytes); err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *usecase) GetAll(page int) ([]*entities.Company, error) {
	data, err := uc.cr.GetAll(page)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	dataBytes, _ := json.Marshal(data)
	if err := uc.q.Write("companies", []byte("getAll"), dataBytes); err != nil {
		return nil, err
	}
	return data, nil
}
func (uc *usecase) Create(in *entities.CreateCompanyInput) (*entities.Company, error) {
	data, err := uc.cr.Create(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	dataBytes, _ := json.Marshal(data)
	if err := uc.q.Write("companies", []byte("create"), dataBytes); err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *usecase) Update(in *entities.UpdateCompanyInput) (*entities.Company, error) {
	data, err := uc.cr.Update(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	dataBytes, _ := json.Marshal(data)
	if err := uc.q.Write("companies", []byte("update"), dataBytes); err != nil {
		return nil, err
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
	if err := uc.q.Write("companies", []byte("delete"), []byte(fmt.Sprintf("%s has been deleted", id.String()))); err != nil {
		return err
	}
	return err
}
