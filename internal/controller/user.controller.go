package controller

import (
	"example.com/m/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// Workflow:  Router ==> Controller ==> Service ==> Repo ==> Model ==> DBs

func (uc *UserController) GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"users":   uc.userService.GetUserInfo(),
	})
}
