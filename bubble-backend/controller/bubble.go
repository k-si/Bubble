package controller

import (
	"bubble/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllBubbles(ctx *gin.Context) {
	bubbleList, err := model.RetrieveAllBubbles()

	// 更复杂的逻辑写在logic层

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		ctx.JSON(http.StatusOK, bubbleList)
	}
}

func CreateBubble(ctx *gin.Context) {
	// 接收bubble
	var bubble model.Bubble
	_ = ctx.BindJSON(&bubble)

	// 更复杂的逻辑写在logic层

	// 入库
	if err := model.CreateBubble(&bubble); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		ctx.JSON(http.StatusOK, bubble)
	}
}

func UpdateBubbleById(ctx *gin.Context) {
	// 接收bubble
	var bubble model.Bubble
	_ = ctx.BindJSON(&bubble)

	// 更复杂的逻辑写在logic层

	// 入库
	if err := model.UpdateBubbleById(&bubble); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		ctx.JSON(http.StatusOK, bubble)
	}
}

func DeleteBubbleById(ctx *gin.Context) {
	// 接收bubble
	id, _ := ctx.Params.Get("id")

	// 更复杂的逻辑写在logic层

	// 入库
	if err := model.DeleteBubbleById(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		ctx.JSON(http.StatusOK, id)
	}
}
