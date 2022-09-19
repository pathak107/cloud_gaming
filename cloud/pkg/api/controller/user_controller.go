package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pathak107/cloudesk/pkg/api/apierrors"
	"github.com/pathak107/cloudesk/pkg/api/dto"
	"github.com/pathak107/cloudesk/pkg/api/entity"
	"github.com/pathak107/cloudesk/pkg/api/service"
)

type UserController struct {
	userSvc *service.UserService
	authSvc *service.AuthService
}

func (c *UserController) Register(ctx *gin.Context) {
	var user dto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user.UserType = string(entity.Customer)
	if err := c.userSvc.Create(&user); err != nil {
		apiError, _ := err.(*apierrors.ApiError)
		ctx.AbortWithError(apiError.Code, apiError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user registered successfully",
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	var login dto.UserLogin
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	jwtToken, err := c.authSvc.Login(&login)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": jwtToken,
	})
}

func NewUserController(userSvc *service.UserService, authSvc *service.AuthService) *UserController {
	return &UserController{
		userSvc: userSvc,
		authSvc: authSvc,
	}
}
