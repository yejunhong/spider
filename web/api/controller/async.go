package controller

import (
	"spider/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
)

/**
 *
 * 书籍同步到生产库
 *
 */
 func (controller *Controller) AsyncBookProduce(c *gin.Context){

	var bookId string = c.Query("bookId")
	bookIdInt64, _ := strconv.ParseInt(bookId, 10, 64) 
	var bookInfo model.CartoonList = controller.Model.GetCartoonInfoById(bookIdInt64)
	var chapterList []model.CartoonChapter = controller.Model.GetCartoonChapterListByNo(bookInfo.ResourceNo)

	var resource = controller.Model.GetCartoonByResourceNo(bookInfo.ResourceNo)

	if resource.BookType == 1 { // 漫画
		var chapterContents []model.CartoonChapterContent =  controller.Model.GetContentsFindByChapterListUniqueId(bookInfo.UniqueId)
		fmt.Println(chapterContents)
	} else {
		for _, v := range chapterList {
			fmt.Println(v)
		}
	}

	c.JSON(200, gin.H{
		"error": 0,
		"msg": "上传同步成功",
	})
}