package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}
func (p *PongController) GetPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
