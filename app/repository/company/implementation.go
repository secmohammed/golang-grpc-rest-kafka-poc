package company

import (
    "github.com/google/uuid"
    "github.com/secmohammed/golang-kafka-grpc-poc/entities"
    "github.com/secmohammed/golang-kafka-grpc-poc/types"
)

type cr struct {
    c types.Container
}

func (cr cr) Get(id uuid.UUID) (*entities.Company, error) {
    var c entities.Company
    result := cr.c.Database().Get().Where("id = ?", id).First(&c)

    return &c, result.Error
}

func (cr cr) GetAll(page int) ([]*entities.Company, error) {
    var t []*entities.Company
    offset := 0
    if page > 0 {
        offset = page - 1
    }
    result := cr.c.Database().Get().Order("created_at DESC").Limit(10).Offset(offset).Find(&t)
    return t, result.Error
}

func (cr cr) Create(company *entities.CreateCompanyInput) (*entities.Company, error) {
    c := &entities.Company{
        Name:        company.Name,
        Description: company.Description,
        Registered:  company.Registered,
        CompanyType: company.CompanyType,
        Headcount:   company.Headcount,
    }
    result := cr.c.Database().Get().Create(c)
    return c, result.Error
}

func (cr cr) Update(company *entities.UpdateCompanyInput) (*entities.Company, error) {
    c := &entities.Company{
        Name:        company.Name,
        Description: company.Description,
        Registered:  company.Registered,
        CompanyType: company.CompanyType,
        Headcount:   company.Headcount,
    }
    result := cr.c.Database().Get().Where("id = ?", company.ID).Save(c)
    return c, result.Error

}

func (cr cr) Delete(id uuid.UUID) error {
    result := cr.c.Database().Get().Where("id = ?", id).Delete(&entities.Company{})
    if result.RowsAffected == 0 {
        return ErrCompanyNotFound
    }
    return result.Error
}

func NewCompanyRepository(c types.Container) CompanyRepository {
    return &cr{c}
}
