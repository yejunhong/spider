package main

import (
	"fmt"
	"paimei/app"
	"github.com/gin-gonic/gin"
	"log"
	"io"
	"io/ioutil"
	"os"
	"net/http"
	"strings"
	"github.com/satori/go.uuid"
	// "github.com/unrolled/secure"
)

func main() {
  var Db = InitDb();
  cartoonController := app.Cartoon{Db}
  downloadController := app.Download{Db}

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

	router.POST("/upload", Fileupload)

	// 漫画
  cartoon := router.Group("/cartoon")
  cartoon.GET("/resource", cartoonController.CartoonResource) // 漫画资源
	cartoon.GET("/list", cartoonController.CartoonList) // 漫画列表
	cartoon.GET("/chapter", cartoonController.CartoonChapter)  // 漫画章节列表
  cartoon.GET("/chapter/content", cartoonController.CartoonChapterContent)  // 漫画章节内容
  
  
  // 下载
  download := router.Group("/download")
  download.GET("/book", downloadController.Book)  // 下载指定资源-书籍
  download.GET("/book/content", downloadController.BookContent)  // 下载指定书籍-所有内容

  download.GET("/chapter", downloadController.Chapter)  // 下载指定书籍-章节
  download.GET("/chapter/content", downloadController.GoodsDel)  // 下载指定书籍-章节内容
  
	router.Run(":6010") // listen and serve on 0.0.0.0:8080
}

/**上传方法**/
func Fileupload(c *gin.Context){
  //得到上传的文件
  file, _, err := c.Request.FormFile("file") //image这个是uplaodify参数定义中的   'fileObjName':'image'
  if err != nil {
		c.JSON(200, gin.H{
			"message": "pong",
		})
    return
  }
  //文件的名称
  // filename := header.Filename
	filename := uuid.NewV4().String() + ".jpg"
  fmt.Println(file, err, filename)
  //创建文件
  out, err := os.Create("static/" + filename)
  //注意此处的 static/uploadfile/ 不是/static/uploadfile/
  if err != nil {
    log.Fatal(err)
  }
  defer out.Close()
  _, err = io.Copy(out, file)
  if err != nil {
      log.Fatal(err)
  }
	c.JSON(200, gin.H{
		"error": 0,
		"msg": filename,
	})
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
