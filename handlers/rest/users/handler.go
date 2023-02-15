package users

import (
    "github.com/gin-gonic/gin"
    "github.com/secmohammed/golang-kafka-grpc-poc/app/usecase/user"
    "github.com/secmohammed/golang-kafka-grpc-poc/entities"
    "github.com/secmohammed/golang-kafka-grpc-poc/utils"
    "net/http"
)

type restHandler struct {
    uuc user.UseCase
}

func NewUserHandler(uuc user.UseCase) UserRestHandler {
    return &restHandler{uuc: uuc}
}

func (h *restHandler) Login(c *gin.Context) {
    payload := &entities.LoginUserInput{}
    if ok := utils.BindData(c, payload); !ok {
        return
    }
    result, err := h.uuc.Login(payload)
    if err != nil {
        c.JSON(utils.Status(err), gin.H{
            "error": err,
        })
        return
    }
    c.JSON(http.StatusOK, SuccessResponse{Data: result})
}

func (h *restHandler) Register(c *gin.Context) {
    param := entities.RegisterUserInput{}
    if ok := utils.BindData(c, &param); !ok {
        return
    }
    result, err := h.uuc.Create(&param)
    if err != nil {
        c.JSON(utils.Status(err), gin.H{
            "error": err,
        })
        return
    }
    c.JSON(http.StatusOK, SuccessResponse{Data: result})
}
