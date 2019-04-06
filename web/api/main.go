package api

import (
	"fmt"
  "spider/model"
  "spider/service"
  "spider/web/api/controller"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
  "strings"
)

func HttpRun(Model *model.Model, listen string) {
  
  var services *service.Service = &service.Service{Models: Model}
  
  var controllers *controller.Controller = &controller.Controller{Model, services}

	router := gin.Default()
	router.Use(Cors())
	router.Static("/static", "./static")
	router.Static("/css", "./admin/dist/css")
	router.Static("/js", "./admin/dist/js")

	router.GET("/admin", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		html, _ := ioutil.ReadFile("./admin/dist/index.html")
		c.String(200, string(html))
  })

  //监听claw websocket连接
	// 漫画
  cartoon := router.Group("/cartoon")
  cartoon.GET("/resource/:id", controllers.CartoonResourceInfo) // 漫画资源
  cartoon.GET("/resource", controllers.CartoonResource) // 漫画资源
  cartoon.POST("/resource", controllers.SetCartoonResource) // 漫画资源
	cartoon.GET("/list", controllers.CartoonList) // 漫画列表
	cartoon.GET("/chapter", controllers.CartoonChapter)  // 漫画章节列表
  cartoon.GET("/chapter/content", controllers.CartoonChapterContent)  // 漫画章节内容
  
  // 下载
  download := router.Group("/download")
  download.GET("/book", controllers.DownloadBookByResourceId)  // 下载指定资源-书籍
  download.GET("/chapter", controllers.DownloadChapterByBookId)  // 下载指定书籍-章节
  
  async := router.Group("/async")
  async.GET("/book", controllers.AsyncBookProduce)  // 同步书籍到生产库
  
  fmt.Println("开启0.0.0.0:", listen)
	router.Run(":" + listen) // listen and serve on 0.0.0.0:8080
}

////// 跨域
func Cors() gin.HandlerFunc {
  return func(c *gin.Context) {
    method := c.Request.Method      //请求方法
    origin := c.Request.Header.Get("Origin")        //请求头部
    var headerKeys []string                             // 声明请求头keys
    for k, _ := range c.Request.Header {
      headerKeys = append(headerKeys, k)
    }
    headerStr := strings.Join(headerKeys, ", ")
    if headerStr != "" {
      headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
    } else {
      headerStr = "access-control-allow-origin, access-control-allow-headers"
    }
    if origin != "" {
      c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
      c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
      c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
      //  header的类型
      c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
      //              允许跨域设置                                                                                                      可以返回其他子段
      c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
      c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
      c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
      c.Set("content-type", "application/json")       // 设置返回格式是json
    }

    //放行所有OPTIONS方法
    if method == "OPTIONS" {
      c.JSON(http.StatusOK, "Options Request!")
    }
    // 处理请求
    c.Next()        //  处理请求
  }
}
