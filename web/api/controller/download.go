package controller

import (
	"github.com/gin-gonic/gin"
)

func (controller *Controller) DownloadBook(c *gin.Context){
	c.JSON(200, gin.H{
		"error": 0,
		"msg": "下载资源-书籍",
	})
}

func (controller *Controller) DownloadChapter(c *gin.Context){
	c.JSON(200, gin.H{
		"error": 0,
		"msg": "下载资源漫画章节列表",
	})
}

func (controller *Controller) DownloadChapterContent(c *gin.Context){
	c.JSON(200, gin.H{
		"error": 0,
		"msg": "下载资源漫画章节内容",
	})
}