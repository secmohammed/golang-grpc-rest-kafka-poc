package companies

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/company"
	"github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/user"
	"github.com/secmohammed/golang-kafka-grpc-poc/config"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	pb "github.com/secmohammed/golang-kafka-grpc-poc/handlers/grpc/pb/companies"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedCompaniesServer
	cu  company.UseCase
	c   config.Repository
	ucc user.UseCase
}

func (s *server) GetCompanyList(_ context.Context, request *pb.GetCompaniesListRequest) (*pb.GetCompaniesListResponse, error) {
	page := request.Page
	data, err := s.cu.GetAll(int(page))
	if err != nil {
		return &pb.GetCompaniesListResponse{}, nil
	}
	var companies []*pb.Company
	for _, company := range data {
		companies = append(companies, &pb.Company{
			Id:          company.ID.String(),
			Name:        company.Name,
			Description: company.Description,
			Headcount:   int32(company.Headcount),
			Registered:  company.Registered,
			Type:        pb.CompanyType(pb.CompanyType_value[string(company.CompanyType)]),
			Time:        timestamppb.New(company.CreatedAt),
		})
	}
	return &pb.GetCompaniesListResponse{
		Companies: companies,
	}, nil

}

func (s *server) GetCompany(_ context.Context, request *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}
	c, err := s.cu.Get(id)
	if err != nil {
		return nil, err
	}
	return &pb.GetCompanyResponse{
		Company: &pb.Company{
			Id:          id.String(),
			Name:        c.Name,
			Description: c.Description,
			Headcount:   int32(c.Headcount),
			Registered:  c.Registered,
			Type:        pb.CompanyType(pb.CompanyType_value[string(c.CompanyType)]),
			Time:        timestamppb.New(c.CreatedAt),
		},
	}, nil
}

func (s *server) CreateCompany(_ context.Context, request *pb.CreateCompanyRequest) (*pb.GetCompanyResponse, error) {
	param := &entities.CreateCompanyInput{
		Name:        request.Name,
		Description: request.Description,
		Headcount:   uint(request.Headcount),
		Registered:  request.Registered,
		CompanyType: entities.CompanyType(pb.CompanyType_name[int32(request.Type)]),
	}
	result, err := s.cu.Create(param)
	if err != nil {
		return nil, err
	}
	return &pb.GetCompanyResponse{
		Company: &pb.Company{
			Id:          result.ID.String(),
			Name:        result.Name,
			Description: result.Description,
			Headcount:   int32(result.Headcount),
			Registered:  result.Registered,
			Type:        pb.CompanyType(pb.CompanyType_value[string(result.CompanyType)]),
			Time:        timestamppb.New(result.CreatedAt),
		},
	}, nil
}

func (s *server) UpdateCompany(_ context.Context, request *pb.UpdateCompanyRequest) (*pb.GetCompanyResponse, error) {
	param := &entities.UpdateCompanyInput{
		Name:        request.Name,
		Description: request.Description,
		Headcount:   uint(request.Headcount),
		Registered:  request.Registered,
		CompanyType: entities.CompanyType(pb.CompanyType_name[int32(request.Type)]),
	}
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}
	param.ID = id

	result, err := s.cu.Update(param)
	if err != nil {
		return nil, err
	}
	return &pb.GetCompanyResponse{
		Company: &pb.Company{
			Id:          id.String(),
			Name:        result.Name,
			Description: result.Description,
			Headcount:   int32(result.Headcount),
			Registered:  result.Registered,
			Type:        pb.CompanyType(pb.CompanyType_value[string(result.CompanyType)]),
			Time:        timestamppb.New(result.CreatedAt),
		},
	}, nil
}

func (s *server) DeleteCompany(_ context.Context, request *pb.DeleteCompanyRequest) (*pb.DeleteCompanyResponse, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {

		return nil, err

	}
	err = s.cu.Delete(id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCompanyResponse{
		Message: fmt.Sprintf("%s is deleted successfully!", id),
	}, nil
}

func NewCompanyServer(cu company.UseCase, c config.Repository, ucc user.UseCase) pb.CompaniesServer {
	return &server{
		pb.UnimplementedCompaniesServer{},
		cu,
		c,
		ucc,
	}
}
