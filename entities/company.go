package entities

import (
	"github.com/google/uuid"
)

type CompanyType string

type Company struct {
	BaseModel
	Name        string      `json:"name" gorm:"type:varchar(255);not null"`
	Description string      `json:"description" gorm:"type:varchar(255);not null"`
	Registered  bool        `json:"registered" gorm:"default:false"`
	Headcount   uint        `json:"headcount" gorm:"default:0"`
	CompanyType CompanyType `json:"company_type" sql:"type:ENUM('Corporations', 'NonProfit', 'Cooperative', 'SoleProprietorship')"`
}

type CreateCompanyInput struct {
	Name        string      `json:"name" binding:"required,min=4"`
	Description string      `json:"description" binding:"required,min=30"`
	Registered  bool        `json:"registered" binding:"required,boolean"`
	Headcount   uint        `json:"headcount" binding:"required,number,gte=0"`
	CompanyType CompanyType `json:"company_type" binding:"required,oneof=Corporations NonProfit Cooperative SoleProprietorship"`
}
type UpdateCompanyInput struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name" binding:"required,min=4"`
	Description string      `json:"description" binding:"required,min=30"`
	Registered  bool        `json:"registered" binding:"required,boolean"`
	Headcount   uint        `json:"headcount" binding:"required,number,gte=0"`
	CompanyType CompanyType `json:"company_type" binding:"required,oneof=Corporations NonProfit Cooperative SoleProprietorship"`
}
