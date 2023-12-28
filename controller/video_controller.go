package controller

import (
	"intro-gin/entity"
	"intro-gin/service"
	"intro-gin/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
var validate *validator.Validate


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
	err = validate.Struct(video)
	if err !=nil{
		return err
	}
	c.service.Save(video)
	return nil
}

func New(s service.ServiceVideo)VideoController{
	validate = validator.New()
	validate.RegisterValidation("is-cool",validators.ValidateCoolTitle)
	return &controller{
		service: s,
	}
}
