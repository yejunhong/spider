package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"spider/service"
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

	if is_test == "1" {
		test = true
	}
	fmt.Println(test)
	var cartoon = controller.Model.GetCartoonByResourceNo(resourceNo)

	go func(){
		var browser service.NodeBrowser = service.NodeBrowser{
			Service: service.Service{controller.Model},
		}
		browser.CreateBrowserClient() // 创建浏览器客户端
		var spider service.Spider = service.Spider{Models: controller.Model, Browser: browser}
		spider.Book(cartoon.Id)
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

	c.JSON(200, gin.H{
		"error": 0,
		"msg": "下载资源漫画章节列表",
	})
}