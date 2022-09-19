package controller

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/pathak107/cloudesk/pkg/api/dto"
// 	"github.com/pathak107/cloudesk/pkg/api/entity"
// 	"github.com/pathak107/cloudesk/pkg/api/service"
// )

// type AppsController struct {
// 	userSvc *service.UserService
// 	authSvc *service.AuthService
// 	appsSvc *service.AppsService
// }

// func (u *UserController) Register(ctx *gin.Context) {
// 	var user dto.User
// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	user.UserType = string(entity.Customer)
// 	if err := u.userSvc.Create(&user); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "user registered successfully",
// 	})
// }

// func (u *UserController) Login(ctx *gin.Context) {
// 	var login dto.UserLogin
// 	if err := ctx.ShouldBindJSON(&login); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	jwtToken, err := u.authSvc.Login(&login)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"jwtToken": jwtToken,
// 	})
// }

// func NewAppsController(userSvc *service.UserService, authSvc *service.AuthService) *UserController {
// 	return &UserController{
// 		userSvc: userSvc,
// 		authSvc: authSvc,
// 	}
// }
