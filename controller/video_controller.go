package controller

import (
	"intro-gin/entity"
	"intro-gin/service"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.ServiceVideo
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err !=nil{
		return err
	}
	c.service.Save(video)
	return nil
}

func New(s service.ServiceVideo)VideoController{
	return &controller{
		service: s,
	}
}
