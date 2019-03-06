package controller

import (
	// "fmt"
	"spider/model"
	"github.com/gin-gonic/gin"
)

func (controller *Controller) CartoonResource(c *gin.Context){
	var page, size, num int64 = controller.Page(c)
	var search string = c.Query("search")
	var list []model.CartoonResource = controller.Model.GetCartoonResources(search, size, num)
	c.JSON(200, gin.H{
		"error": 0,
		"msg": gin.H{
			"page": page,
			"list": list,
			"count": 0,
		},
	})
}

func (controller *Controller) CartoonList(c *gin.Context){
	var page, size, num int64 = controller.Page(c)
	var search string = c.Query("search")
	var list []model.CartoonList = controller.Model.GetCartoons(search, size, num)
	c.JSON(200, gin.H{
		"error": 0,
		"msg": gin.H{
			"page": page,
			"list": list,
			"count": 0,
		},
	})
}

func (controller *Controller) CartoonChapter(c *gin.Context){
	c.JSON(200, gin.H{
		"error": 0,
		"msg": "漫画章节列表",
	})
}

func (controller *Controller) CartoonChapterContent(c *gin.Context){
	c.JSON(200, gin.H{
		"error": 0,
		"msg": "漫画章节内容",
	})
}