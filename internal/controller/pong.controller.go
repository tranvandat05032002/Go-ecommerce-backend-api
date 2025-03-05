package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Pong struct{}

func PongController() *Pong {
	return &Pong{}
}
func (p *Pong) GetPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
