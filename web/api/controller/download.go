package controller

import (
	"github.com/gin-gonic/gin"
	"spider/service"
	"strconv"
)

/**
 *
 * 根据资源id 下载书籍列表, 异步下载 提供ws实时数据
 * @query resource_id 资源id
 * @query is_test 是否下载测试 1测试 默认 非测试
 * @return json{error: 0, msg: interface{}}
 *
 */
func (controller *Controller) DownloadBookByResourceId(c *gin.Context){

	var resourceId string = c.Query("resourceId")
	resourceIdInt64, _ := strconv.ParseInt(resourceId, 10, 64) 
	var cartoon = controller.Model.GetCartoonById(resourceIdInt64)

	go func(){
		var browser service.NodeBrowser = service.NodeBrowser{
			Service: service.Service{controller.Model},
		}
		browser.CreateBrowserClient() // 创建浏览器客户端
		var resource service.Spider = service.Spider{Models: controller.Model, Browser: browser}
		resource.SpiderBookByResourceId(cartoon.Id)
	}()
	
	c.JSON(200, gin.H{
		"error": 0,
		"msg": "已经提交",
	})
}
 
/**
 *
 * 根据数据书籍id 下载书籍章节列表, 异步下载 提供ws实时数据
 * @query book_id 书籍id
 * @return json{error: 0, msg: interface{}}
 *
 */
func (controller *Controller) DownloadChapterByBookId(c *gin.Context){
	var bookId string = c.Query("bookId")
	go func(){
		var browser service.NodeBrowser = service.NodeBrowser{
			Service: service.Service{controller.Model},
		}
		browser.CreateBrowserClient() // 创建浏览器客户端
		var resource service.Spider = service.Spider{Models: controller.Model, Browser: browser}

		bookIdInt64, _ := strconv.ParseInt(bookId, 10, 64) 
		resource.SpiderChapterByBookId(bookIdInt64)
		resource.SpiderContentByBookId(bookIdInt64)
	}()
	c.JSON(200, gin.H{
		"error": 0,
		"msg": "下载资源漫画章节列表",
	})
}