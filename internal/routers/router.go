package routers

import (
	"example.com/m/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		//v1.GET("/ping", controller.NewPongController().GetPong)
		v1.GET("/user/1", controller.NewUserController().GetUserInfo)
	}
	return r
}
