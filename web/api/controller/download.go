package controller

import (
	"github.com/gin-gonic/gin"
	"spider/model"
)

/**
 *
 * 根据资源id 下载书籍列表, 异步下载 提供ws实时数据
 * @query resource_id 资源id
 * @query is_test 是否下载测试 1测试 默认 非测试
 * @return json{error: 0, msg: interface{}}
 *
 */
func (controller *Controller) DownloadBook(c *gin.Context){

	var resourceNo string = c.Query("resource_no")
	var is_test string = c.Query("is_test") // 是否测试
	var test bool = false

	if is_test == 1{
		test = true
	}

	var cartoon = controller.Model.GetCartoonByResourceNo(resourceNo)

	go func(){
		controller.Service.CrawlBookList(cartoon, "", test)
	}()
	
	c.JSON(200, gin.H{
		"error": 0,
		"msg": "已经提交",
	})
}

/**
 *
 * 根据数据资源id 下载书籍章节列表, 异步下载 提供ws实时数据
 * @query resource_id 资源id
 * @query is_test 是否下载测试 1测试 默认 非测试
 * @return json{error: 0, msg: interface{}}
 *
 */
func (controller *Controller) DownloadChapter(c *gin.Context){

	var listId string = c.Query("id")
	var is_test string = c.Query("is_test") // 是否测试
	var test bool = false

	if is_test == 1{
		test = true
	}

	listIdInt64, _ := strconv.ParseInt(listId, 10, 64)

	var cartoonList = controller.Model.GetCartoonInfoById(listIdInt64)
	var cartoon = controller.Model.GetCartoonByResourceNo(cartoonList.ResourceNo)

	go func(){
		controller.Service.CrawlBookChapter(cartoon, cartoonList, "", test)
	}()
	
	c.JSON(200, gin.H{
		"error": 0,
		"msg": "下载资源漫画章节列表",
	})
}

func (controller *Controller) DownloadChapterContent(c *gin.Context){

	var chapterId string = c.Query("id")
	var is_test string = c.Query("is_test") // 是否测试
	var test bool = false

	if is_test == 1{
		test = true
	}

	chapterIdInt64, _ := strconv.ParseInt(chapterId, 10, 64)

	var cartoonChapter = controller.Model.GetCartoonChapterInfoById(chapterIdInt64)
	var cartoon = controller.Model.GetCartoonByResourceNo(cartoonChapter.ResourceNo)
	
	go func(){
		controller.Service.CrawlBookChapterContent(cartoon, cartoonChapter "", test)
	}()
	
	c.JSON(200, gin.H{
		"error": 0,
		"msg": "下载资源漫画章节内容",
	})
}