package router

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func setupRouters() *gin.Engine{
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.GET("/bubble", controller.GetAllBubbles)
		v1.POST("/bubble", controller.CreateBubble)
		v1.PUT("/bubble", controller.UpdateBubbleById)
		v1.DELETE("/bubble/:id", controller.DeleteBubbleById)
	}
	return r
}

func Run() {
	if err := setupRouters().Run(":9090"); err != nil {
		panic("server start failed")
	}
}
