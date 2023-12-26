package main

import (
	"intro-gin/controller"
	"intro-gin/service"

	"github.com/gin-gonic/gin"
)

var( 
	videoService service.ServiceVideo = service.New()
	videoController controller.VideoController = controller.New(videoService)
)


func main(){
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message" : "Ok",
		})
	})

	server.GET("/videos",func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos",func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}