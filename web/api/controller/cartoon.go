package controller

import (
	// "fmt"
	"spider/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

func (controller *Controller) CartoonResourceInfo(c *gin.Context){
	var id string = c.Param("id")
	idInt64, _ := strconv.ParseInt(id, 10, 64)
	var res model.CartoonResource = controller.Model.GetCartoonById(idInt64)
	contents, _ := ioutil.ReadFile("./node/config/" + res.ConfigName + ".ts")
	c.JSON(200, gin.H{
		"error": 0,
		"msg": gin.H{
			"info": res,
			"config": string(contents),
		},
	})
}

func (controller *Controller) CartoonResource(c *gin.Context){
	var page, size, num int64 = controller.Page(c)
	var search string = c.Query("search")

	var list []model.CartoonResource
	var count int64
	list, count = controller.Model.GetCartoonResources(search, size, num)
	c.JSON(200, gin.H{
		"error": 0,
		"msg": gin.H{
			"page": page,
			"pagesize": size,
			"list": list,
			"count": count,
		},
	})
}

func (controller *Controller) CartoonList(c *gin.Context){
	var page, size, num int64 = controller.Page(c)
	var resource_no string = c.Query("resource_no")
	var search string = c.Query("search")

	var list []model.CartoonList
	var count int64

	list, count = controller.Model.GetCartoons(resource_no, search, size, num)
	
	c.JSON(200, gin.H{
		"error": 0,
		"msg": gin.H{
			"page": page,
			"pagesize": size,
			"list": list,
			"count": count,
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