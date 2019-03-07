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
	var resource_no string = c.Query("resource_no")
	var search string = c.Query("search")
	var list []model.CartoonList = controller.Model.GetCartoons(resource_no, search, size, num)
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
	var list_unique_id string = c.Query("list_unique_id")
	var list []model.CartoonChapter = controller.Model.GetChaptersFindByListUniqueId(list_unique_id)
	c.JSON(200, gin.H{
		"error": 0,
		"msg": list,
	})
}

func (controller *Controller) CartoonChapterContent(c *gin.Context){
	var chapter_unique_id string = c.Query("chapter_unique_id")
	var list []model.CartoonChapterContent = controller.Model.GetContentsFindByChapterUniqueId(chapter_unique_id)
	c.JSON(200, gin.H{
		"error": 0,
		"msg": list,
	})
}