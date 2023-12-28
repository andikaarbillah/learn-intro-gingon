package main

import (
	"intro-gin/controller"
	"intro-gin/middlewares"
	"intro-gin/service"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.ServiceVideo       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	middlewares.SetupLogFile()
	server := gin.Default()
	server.Use(gin.Recovery(), middlewares.Logger(),middlewares.BasicAuth(),gindump.Dump())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ok",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil{
			ctx.JSON(400,gin.H{"error":err.Error()} )
		}else{
			ctx.JSON(200,gin.H{"message":"Video Input is Valid"})
		}
		
	})

	server.Run(":8080")
}
