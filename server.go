package main

import (
	"github.com/gin-gonic/gin"
	"github.com/paradise98425/go-gin/controller"
	"github.com/paradise98425/go-gin/service"
)

var (
	videoService service.VideoService = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()
	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.Run(":8081")
}
