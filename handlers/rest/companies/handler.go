package companies

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	company2 "github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/company"
	"github.com/secmohammed/golang-kafka-grpc-poc/entities"
	"github.com/secmohammed/golang-kafka-grpc-poc/utils"
	"net/http"
	"strconv"
)

type restHandler struct {
	cu company2.UseCase
}

func NewCompanyHandler(cu company2.UseCase) CompaniesRestHandler {
	return &restHandler{cu: cu}
}

func (h *restHandler) GetCompanies(c *gin.Context) {
	page := 0
	if pageString, exists := c.GetQuery("page"); exists {
		p, err := strconv.Atoi(pageString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.NewBadRequest(err.Error()))
			return
		}
		page = p

	}
	data, err := h.cu.GetAll(page)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		c.JSON(http.StatusInternalServerError, utils.NewInternal())
	}
}

func (h *restHandler) GetCompany(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewBadRequest(err.Error()))
		return

	}
	data, err := h.cu.Get(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: data})
	} else {
		if errors.Is(err, company2.ErrCompanyNotFound) {
			c.JSON(http.StatusNotFound, utils.NewNotFound("company", fmt.Sprintf("%s", id)))
		} else {
			c.JSON(http.StatusInternalServerError, utils.NewInternal())
		}
	}
}

func (h *restHandler) CreateCompany(c *gin.Context) {
	param := entities.CreateCompanyInput{}
	if ok := utils.BindData(c, &param); !ok {
		return
	}
	result, err := h.cu.Create(&param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, utils.NewInternal())
	}
}

func (h *restHandler) UpdateCompany(c *gin.Context) {
	param := &entities.UpdateCompanyInput{}
	if ok := utils.BindData(c, &param); !ok {
		return
	}
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewBadRequest(err.Error()))
		return
	}
	param.ID = id
	result, err := h.cu.Update(param)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: result})
	} else {
		c.JSON(http.StatusInternalServerError, utils.NewInternal())
	}
}

func (h *restHandler) DeleteCompany(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewBadRequest(err.Error()))
		return

	}
	err = h.cu.Delete(id)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: fmt.Sprintf("id: %s. successfully deleted", id)})
	} else {
		c.JSON(http.StatusInternalServerError, utils.NewInternal())
	}
}
